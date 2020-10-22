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
// https://developer.swift.com/oauth-reference#section/Authentication/clientAuth
type AppCredentials struct {
	BasicAuthUser string
	BasicAuthPass string
	UserName      string
	Password      string
}

// NewAppCredentials - factory function
// `variables` in Postman collection
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

// AuthenticationTokens - returned after successful authentication
// `AccessToken` used in http Authorization headers
// `ExpiresIn` by default is 1800-1 seconds (30min)
// `RefreshToken` used for requiring new AuthenticationTokens and is alive for 86400-1 seconds (1 day)
// `TokenType` is a part of Authorization header: ex. req.Header.Set("Authorization", "Bearer "+t.AuthenticationTokens
//
// It is strongly recommended that your application dispose tokens that are no longer needed.
// SWIFT will invalidate the tokens from further use if you do.
// Once invalidated, they can no longer be used to access SWIFT APIs.
type AuthenticationTokens struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             string `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
	IdmService            string `json:"idm_service"`
}

// InvokeAuthTokens - fetches from SWIFT api new `AuthenticationTokens`
func InvokeAuthTokens(creds *AppCredentials, ctx context.Context, e env) (*AuthenticationTokens, error) {
	urlValues := url.Values{}
	urlValues.Set("username", creds.UserName)
	urlValues.Set("password", creds.Password)
	urlValues.Set("grant_type", "password")

	return invokeAccessToken(urlValues, creds, ctx, e)
}

// RefreshAuthTokens - refreshes given `AuthenticationTokens`
func RefreshAuthTokens(authTokens *AuthenticationTokens, creds *AppCredentials, ctx context.Context, e env) (*AuthenticationTokens, error) {
	data := url.Values{}
	data.Set("refresh_token", authTokens.RefreshToken)
	data.Set("grant_type", "refresh_token")

	return invokeAccessToken(data, creds, ctx, e)
}

// invokeAccessToken - sends requests to SWIFT api and creates new or refreshes existing token
func invokeAccessToken(reqValues url.Values, creds *AppCredentials, ctx context.Context, e env) (*AuthenticationTokens, error) {
	req, err := http.NewRequest(http.MethodPost, getAccessTokenUrl(e), strings.NewReader(reqValues.Encode()))
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
		return nil, NewSWIFTError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}

	// parse response to appropriate struct
	token := AuthenticationTokens{}
	if err := json.Unmarshal(b.Bytes(), &token); err != nil {
		return nil, err
	}
	return &token, nil
}

// RevokeAccessToken - removes/disposes authentication tokens
// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
func RevokeAccessToken(authToken *AuthenticationTokens, creds *AppCredentials, ctx context.Context, e env) error {
	data := url.Values{}
	data.Set("token", authToken.AccessToken)
	data.Set("token_type_hint", "access_token")

	return revokeToken(data, creds, ctx, e)
}

// RevokeRefreshToken - removes/disposes authentication tokens
// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
func RevokeRefreshToken(authToken *AuthenticationTokens, creds *AppCredentials, ctx context.Context, e env) error {
	data := url.Values{}
	data.Set("token", authToken.RefreshToken)
	data.Set("token_type_hint", "refresh_token")

	return revokeToken(data, creds, ctx, e)
}

// revokeToken - actual implementation of revoking access tokens
func revokeToken(data url.Values, creds *AppCredentials, ctx context.Context, e env) error {
	// create request with form values url encoded
	req, err := http.NewRequest(http.MethodPost, getRevokeTokenUrl(e), strings.NewReader(data.Encode()))
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

		return NewSWIFTError(resp.StatusCode, b.String(), resp.Header.Get("Content-Type"))
	}
	return nil
}
