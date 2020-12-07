package swift

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	testFilePath := path.Join(os.Getenv("GOPATH"), "src/github.com/alifcapital/swift/tests.json")
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
func validateAccessTokenFields(t *testing.T, token *AuthenticationToken) {
	expectedTokenTime := fmt.Sprintf("%d", (60*30)-1) // 30 min
	assert.Equal(t, token.ExpiresIn, expectedTokenTime)

	expectedExpRefreshTime := fmt.Sprintf("%d", (60*60*24)-1) // 1 day
	assert.Equal(t, token.RefreshTokenExpiresIn, expectedExpRefreshTime)

	assert.Equal(t, token.TokenType, "Bearer")
}

// ------------------------------------------------------------------

// TestInvokeAuthTokens - testing https://developer.swift.com/oauth-reference#operation/getAccessToken
func TestInvokeAuthTokens(t *testing.T) {
	// grab configs
	conf := newTestsConfig()

	// make credentials instance
	creds := NewAppCredentials(conf.BasicAuthUser, conf.BasicAuthPass, conf.UserName, conf.Password)
	ctx := context.Background()
	e := SandBoxEnv

	// sends request to SWIFT api, returns new access token instance
	authToken, err := InvokeAuthTokens(ctx, creds, e)
	assert.Nil(t, err)

	// assert expectations
	validateAccessTokenFields(t, authToken)

	// after performing tests, it would be a good idea to `remove` token
	// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
	err = RevokeAccessToken(ctx, creds, authToken, e)
	assert.Nil(t, err)
}

func TestRefreshAuthTokens(t *testing.T) {
	// grab configs
	conf := newTestsConfig()

	// make credentials instance
	creds := NewAppCredentials(conf.BasicAuthUser, conf.BasicAuthPass, conf.UserName, conf.Password)
	ctx := context.Background()
	e := SandBoxEnv

	// sends request to SWIFT api, returns new access token instance
	authTokens, err := InvokeAuthTokens(ctx, creds, e)
	assert.Nil(t, err)

	// try to refresh authTokens
	refreshedToken, err := RefreshAuthTokens(ctx, creds, authTokens, e)
	assert.Nil(t, err)

	// assert expectations
	validateAccessTokenFields(t, refreshedToken)

	// after validating fields clean-up token
	// https://developer.swift.com/oauth-reference#operation/revokeAccessToken
	err = RevokeRefreshToken(ctx, creds, refreshedToken, e)
	assert.Nil(t, err)

	// NOTE: original access token is not valid after refreshing it
	// and server error will be returned
	err = RevokeAccessToken(ctx, creds, authTokens, e)
	assert.Error(t, err)
}