package swift_sdk

import (
	"net/url"
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
		if err != nil {
			t.Fatal(err)
		}
		if prodUrl.Host != "api.swift.com" {
			t.Fatal("not a production host", prodUrl.Host)
		}

		sandBoxUr, err := url.Parse(v(SandBoxEnv))
		if err != nil {
			t.Fatal(err)
		}
		if sandBoxUr.Host != "sandbox.swift.com" {
			t.Fatal("not a sandbox host")
		}
	}
}
