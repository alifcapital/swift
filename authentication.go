package swift

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// AppCredentials is a credentials to make OAuth2 request
type AppCredentials struct {
	BasicAuthUser string // consumer_key
	BasicAuthPass string // consumer_secret
	Username      string // license_id
	Password      string // license_key
}

// AuthenticationToken returned after successful authentication
// `AccessToken` used in http Authorization headers
// `ExpiresIn` by default is 1800-1 seconds (30min)
// `RefreshToken` used for requiring new AuthenticationToken and is alive for 86400-1 seconds (1 day)
// `TokenType` is a part of Authorization header: ex. req.Header.Set("Authorization", "Bearer "+t.AuthenticationToken
//
// It is strongly recommended that your application dispose tokens that are no longer needed.
// SWIFT will invalidate the tokens from further use if you do.
// Once invalidated, they can no longer be used to access SWIFT APIs.
type AuthenticationToken struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             string `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
}

// GetAuthToken fetches from SWIFT api new `AuthenticationToken`.
// https://developer.swift.com/oauth-reference#section/Authentication/clientAuth
//
// CAUTION: getting new access token invalidates previous one, so DO NOT use it
// when you are using API methods for SWIFT-API. Getting new access tokens by need is
// managed automatically by API.
func (api *API) GetAuthToken(ctx context.Context) (*AuthenticationToken, error) {
	urlValues := make(url.Values)
	urlValues.Set("username", api.credentials.Username)
	urlValues.Set("password", api.credentials.Password)
	urlValues.Set("grant_type", "password")
	return api.invokeToken(ctx, "/oauth2/v1/token", urlValues)
}

// RefreshAuthToken fetches from SWIFT api new `AuthenticationToken`.
// https://developer.swift.com/oauth-reference#section/Authentication/clientAuth
func (api *API) RefreshAuthToken(ctx context.Context) (*AuthenticationToken, error) {
	token, err := api.reuseTokenSource.Token()
	if err != nil {
		return nil, err
	}
	urlValues := make(url.Values)
	urlValues.Set("refresh_token", token.RefreshToken)
	urlValues.Set("grant_type", "refresh_token")
	return api.invokeToken(ctx, "/oauth2/v1/token", urlValues)
}

// RevokeAuthToken removes/disposes authentication tokens.
// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
func (api *API) RevokeAuthToken(ctx context.Context) error {
	token, err := api.reuseTokenSource.Token()
	if err != nil {
		return err
	}
	urlValues := make(url.Values)
	urlValues.Set("token", token.AccessToken)
	return api.revokeToken(ctx, "/oauth2/v1/revoke", urlValues)
}

func (api *API) invokeToken(ctx context.Context, path string, urlValues url.Values) (*AuthenticationToken, error) {
	response, err := api.sendRequest(ctx, path, urlValues)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var b bytes.Buffer
	_, err = b.ReadFrom(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, NewErrWithHTTPResponse(response, errors.New("not ok"))
	}
	var token AuthenticationToken
	if err := json.Unmarshal(b.Bytes(), &token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (api *API) revokeToken(ctx context.Context, path string, urlValues url.Values) error {
	response, err := api.sendRequest(ctx, path, urlValues)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return NewErrWithHTTPResponse(response, err)
	}
	return nil
}

func (api *API) sendRequest(ctx context.Context, path string, urlValues url.Values) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, api.url(path), strings.NewReader(urlValues.Encode()))
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(api.credentials.BasicAuthUser, api.credentials.BasicAuthPass)
	return api.httpClient.Do(request)
}

// Token gets new access token. You SHOULD NOT call it.
func (api *API) Token() (*oauth2.Token, error) {
	token, err := api.GetAuthToken(context.Background())
	if err != nil {
		return nil, err
	}
	oauth2Token, err := token.toOAUTH2Token()
	if err != nil {
		return nil, err
	}
	return oauth2Token, nil
}

func (token *AuthenticationToken) toOAUTH2Token() (*oauth2.Token, error) {
	expiresIn, err := strconv.Atoi(token.ExpiresIn)
	if err != nil {
		return nil, err
	}
	expiresInDate := time.Now().Add(time.Duration(expiresIn) * time.Second)
	oauth2Token := oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       expiresInDate,
	}
	return &oauth2Token, nil
}

func (api *API) url(path string) string {
	return api.basePath + path
}
