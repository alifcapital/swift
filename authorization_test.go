package swift_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

// testsConfig - credentials for testing apis
// provide ONLY `sandbox environment` credentials
type testsConfig struct {
	BasicAuthUser string `json:"basic_auth_user"`
	BasicAuthPass string `json:"basic_auth_pass"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	LauAppId      string `json:"lau_app_id"`
	LauVersion    string `json:"lau_version"`
	LauHmacKey    string `json:"lau_hmac_key"`
	AppApiKey     string `json:"app_api_key"`
	Xbic          string `json:"xbic"`
	XApiKey       string `json:"x_api_key"`
}

// newTestsConfig - read file and make tests configs
func newTestsConfig() *testsConfig {
	testFilePath := path.Join(os.Getenv("GOPATH"), "src/github.com/alifcapital/swift_sdk/tests.json")
	fileContent, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}

	tc := testsConfig{}
	if err := json.Unmarshal(fileContent, &tc); err != nil {
		panic(err)
	}
	return &tc
}

// validateAccessTokenFields - check lifetimes and type of token
func validateAccessTokenFields(t *testing.T, token *AccessToken) {
	expectedTokenTime := fmt.Sprintf("%d", (60*30)-1) // 30 min
	if token.ExpiresIn != expectedTokenTime {
		t.Errorf("unexpected `ExpiresIn`, expected: %s, got: %s", expectedTokenTime, token.RefreshTokenExpiresIn)
	}

	expectedExpRefreshTime := fmt.Sprintf("%d", (60*60*24)-1) // 1 day
	if token.RefreshTokenExpiresIn != expectedExpRefreshTime {
		t.Errorf("unexpected `RefreshTokenExpiresIn`, expected: %s, got: %s", expectedExpRefreshTime, token.RefreshTokenExpiresIn)
	}

	if token.TokenType != "Bearer" {
		t.Errorf("accessToken must be a `Bearer` type in switf OAuth APIs")
	}
}

// ------------------------------------------------------------------

// TestGetAccessToken - testing https://developer.swift.com/oauth-reference#operation/getAccessToken
func TestGetAccessToken(t *testing.T) {
	// grab configs
	conf := newTestsConfig()

	// make credentials instance
	creds := NewAuthorizationCredentials(conf.BasicAuthUser, conf.BasicAuthPass, conf.UserName, conf.Password)
	ctx := context.Background()
	e := SandBoxEnv

	// sends request to SWIFT api, returns new access token instance
	accessToken, err := GetAccessToken(creds, ctx, e)
	if err != nil {
		t.Fatal(err)
	}

	// assert expectations
	validateAccessTokenFields(t, accessToken)

	// after performing tests, it would be a good idea to `remove` token
	// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
	if err := RevokeAccessToken(accessToken, creds, ctx, e); err != nil {
		t.Fatal(err)
	}
}

// TestRefreshAccessToken - test refreshing access token
func TestRefreshAccessToken(t *testing.T) {
	// grab configs
	conf := newTestsConfig()

	// make credentials instance
	creds := NewAuthorizationCredentials(conf.BasicAuthUser, conf.BasicAuthPass, conf.UserName, conf.Password)
	ctx := context.Background()
	e := SandBoxEnv

	// sends request to SWIFT api, returns new access token instance
	accessToken, err := GetAccessToken(creds, ctx, e)
	if err != nil {
		t.Fatal(err)
	}

	// try to refresh accessToken
	refreshedToken, err := RefreshAccessToken(accessToken, creds, ctx, e)
	if err != nil {
		t.Fatal(err)
	}

	// assert expectations
	validateAccessTokenFields(t, refreshedToken)

	// after validating fields clean-up token
	// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
	if err := RevokeRefreshToken(refreshedToken, creds, ctx, e); err != nil {
		t.Fatal(err)
	}

	// NOTE: original access token is not valid after refreshing it
	// and server error will be returned
	if err := RevokeAccessToken(accessToken, creds, ctx, e); err == nil {
		t.Fatal("expected error but got nil")
	}
}
