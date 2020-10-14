package swift_sdk

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

func TestUrls(t *testing.T) {
	testFns := []func(e env) string{
		getAccessTokenUrl,
		getRevokeTokenUrl,
		getVerificationUrl,
	}

	for _, v := range testFns {
		prodUrl, err := url.Parse(v(ProductionEnv))
		assert.Nil(t, err)
		assert.Equal(t, prodUrl.Host, "api.swift.com")

		sandBoxUr, err := url.Parse(v(SandBoxEnv))
		assert.Nil(t, err)
		assert.Equal(t, sandBoxUr.Host, "sandbox.swift.com")
	}
}

func TestUrls_WithParams(t *testing.T) {
	bic := "FAKEBIC123"
	prodUrl, err := url.Parse(getBikDetailsUrl(ProductionEnv, bic))
	assert.Nil(t, err)
	assert.Equal(t, prodUrl.Host, "api.swift.com")

	// expecting bic appended at the end of Path
	prodParams := strings.Split(prodUrl.Path, "/")
	assert.Equal(t, prodParams[len(prodParams)-1], bic)

	// ------------------------------------------------------------------

	sandBoxUr, err := url.Parse(getBikDetailsUrl(SandBoxEnv, bic))
	assert.Nil(t, err)
	assert.Equal(t, sandBoxUr.Host, "sandbox.swift.com")

	// expecting bic appended at the end of Path
	sandBoxParams := strings.Split(sandBoxUr.Path, "/")
	assert.Equal(t, sandBoxParams[len(sandBoxParams)-1], bic)
}
