package swift

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// AppCredentials - credentials to make OAuth2 request
type AppCredentials struct {
	BasicAuthUser string
	BasicAuthPass string
	UserName      string
	Password      string
}

// NewAppCredentials - factory function
//
// baUser - {{consumer_key}}
// baPass - {{consumer_secret}}
// userName - {{license_id}}
// password - {{license_key}}
func NewAppCredentials(baUser, baPass, userName, password string) *AppCredentials {
	return &AppCredentials{
		BasicAuthUser: baUser,
		BasicAuthPass: baPass,
		UserName:      userName,
		Password:      password,
	}
}

// AuthenticationToken - returned after successful authentication
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
	IdmService            string `json:"idm_service"`
}

// InvokeAuthTokens - fetches from SWIFT api new `AuthenticationToken`
// https://developer.swift.com/oauth-reference#section/Authentication/clientAuth
func InvokeAuthTokens(ctx context.Context, creds *AppCredentials, e env) (*AuthenticationToken, error) {
	urlValues := url.Values{}
	urlValues.Set("username", creds.UserName)
	urlValues.Set("password", creds.Password)
	urlValues.Set("grant_type", "password")

	return invokeAccessToken(ctx, creds, urlValues, e)
}

// RefreshAuthTokens - refreshes given `AuthenticationToken`
func RefreshAuthTokens(ctx context.Context, creds *AppCredentials, authToken *AuthenticationToken, e env) (*AuthenticationToken, error) {
	urlValues := url.Values{}
	urlValues.Set("refresh_token", authToken.RefreshToken)
	urlValues.Set("grant_type", "refresh_token")

	return invokeAccessToken(ctx, creds, urlValues, e)
}

// invokeAccessToken - sends requests to SWIFT api and creates new or refreshes existing token
func invokeAccessToken(ctx context.Context, creds *AppCredentials, urlValues url.Values, e env) (*AuthenticationToken, error) {
	req, err := http.NewRequest(http.MethodPost, accessTokenUrl(e), strings.NewReader(urlValues.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// set any context given to request
	req = req.WithContext(ctx)

	// set basic auth
	req.SetBasicAuth(creds.BasicAuthUser, creds.BasicAuthPass)

	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// load response body
	var b bytes.Buffer
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	// handle non 200 code
	if resp.StatusCode != http.StatusOK {
		return nil, NewHttpError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}

	// parse response to appropriate struct
	token := AuthenticationToken{}
	if err := json.Unmarshal(b.Bytes(), &token); err != nil {
		return nil, err
	}
	return &token, nil
}

// RevokeAccessToken - removes/disposes authentication tokens
// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
func RevokeAccessToken(ctx context.Context, creds *AppCredentials, authToken *AuthenticationToken, e env) error {
	urlValues := url.Values{}
	urlValues.Set("token", authToken.AccessToken)
	urlValues.Set("token_type_hint", "access_token")

	return revokeToken(ctx, creds, urlValues, e)
}

// RevokeRefreshToken - removes/disposes authentication tokens
// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
func RevokeRefreshToken(ctx context.Context, creds *AppCredentials, authToken *AuthenticationToken, e env) error {
	urlValues := url.Values{}
	urlValues.Set("token", authToken.RefreshToken)
	urlValues.Set("token_type_hint", "refresh_token")

	return revokeToken(ctx, creds, urlValues, e)
}

// revokeToken - actual implementation of revoking access tokens
func revokeToken(ctx context.Context, creds *AppCredentials, urlValues url.Values, e env) error {
	// create request with form values url encoded
	req, err := http.NewRequest(http.MethodPost, revokeTokenUrl(e), strings.NewReader(urlValues.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// set any context given to request
	req = req.WithContext(ctx)

	// set basic auth
	req.SetBasicAuth(creds.BasicAuthUser, creds.BasicAuthPass)

	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// handle non 200 code
	if resp.StatusCode != http.StatusOK {
		// load response body
		var b bytes.Buffer
		_, err = b.ReadFrom(resp.Body)
		if err != nil {
			return err
		}

		return NewHttpError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}
	return nil
}
