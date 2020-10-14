package swift_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// AuthorizationCredentials - required to make OAuth 2.0 request
// https://developer.swift.com/oauth-reference#section/Authentication/clientAuth
type AuthorizationCredentials struct {
	BasicAuthUser string
	BasicAuthPass string
	UserName      string
	Password      string
}

// NewAuthorizationCredentials - factory function
// baUser - {{consumer_key}}
// baPass - {{consumer_secret}}
// userName - {{license_id}}
// password - {{license_key}}
func NewAuthorizationCredentials(baUser, baPass, userName, password string) *AuthorizationCredentials {
	return &AuthorizationCredentials{
		BasicAuthUser: baUser,
		BasicAuthPass: baPass,
		UserName:      userName,
		Password:      password,
	}
}

// ------------------------------------------------------------------

// AccessToken - returned after successful access token initialization
type AccessToken struct {
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	IdmService            string `json:"idm_service"`
	ExpiresIn             string `json:"expires_in"`
}

// GetAccessToken - fetches from remote SWIFT api new `AccessToken` with given credentials by specified `env`
func GetAccessToken(creds *AuthorizationCredentials, ctx context.Context, e env) (*AccessToken, error) {
	// create request with form values url encoded
	data := url.Values{}
	data.Set("username", creds.UserName)
	data.Set("password", creds.Password)
	data.Set("grant_type", "password")

	return invokeAccessToken(data, creds, ctx, e)
}

// RefreshAccessToken - refresh given access token
func RefreshAccessToken(accessToken *AccessToken, creds *AuthorizationCredentials, ctx context.Context, e env) (*AccessToken, error) {
	// create request with form values url encoded
	data := url.Values{}
	data.Set("refresh_token", accessToken.RefreshToken)
	data.Set("grant_type", "refresh_token")

	return invokeAccessToken(data, creds, ctx, e)
}

// invokeAccessToken - sends requests to get new or refresh token
func invokeAccessToken(data url.Values, creds *AuthorizationCredentials, ctx context.Context, e env) (*AccessToken, error) {
	req, err := http.NewRequest(http.MethodPost, getAccessTokenUrl(e), strings.NewReader(data.Encode()))
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
		msg := fmt.Sprintf("request failed with status: %d, with message: %s", resp.StatusCode, b.String())
		return nil, errors.New(msg)
	}

	// parse response to appropriate struct
	token := AccessToken{}
	if err := json.Unmarshal(b.Bytes(), &token); err != nil {
		return nil, err
	}
	return &token, nil
}

// ------------------------------------------------------------------

// RevokeAccessToken - removes access token
// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
func RevokeAccessToken(token *AccessToken, creds *AuthorizationCredentials, ctx context.Context, e env) error {
	data := url.Values{}
	data.Set("token", token.AccessToken)
	data.Set("token_type_hint", "access_token")

	return revokeToken(data, creds, ctx, e)
}

// RevokeAccessToken - removes access token
// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
func RevokeRefreshToken(token *AccessToken, creds *AuthorizationCredentials, ctx context.Context, e env) error {
	data := url.Values{}
	data.Set("token", token.RefreshToken)
	data.Set("token_type_hint", "refresh_token")

	return revokeToken(data, creds, ctx, e)
}

// revokeToken - actual implementation of revoking tokens
func revokeToken(data url.Values, creds *AuthorizationCredentials, ctx context.Context, e env) error {
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

		msg := fmt.Sprintf("request failed with status: %d, with message: %s", resp.StatusCode, b.String())
		return errors.New(msg)
	}

	return nil
}
