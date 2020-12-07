package swift

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	realBic  = "BUKBGB22XXX"
	realIBAN = "GB33BUKB20201555555555"
)

func TestGetDetailsOfBic(t *testing.T) {
	conf := newTestsConfig()
	ctx := context.Background()

	// send request
	details, err := GetDetailsOfBic(ctx, realBic, conf.XApiKey, SandBoxEnv)
	assert.Nil(t, err)
	assert.NotNil(t, details)
}

func TestGetDetailsOfBic_WithoutCredentials(t *testing.T) {
	ctx := context.Background()

	// send request
	details, err := GetDetailsOfBic(ctx, realBic, "no-key", SandBoxEnv)
	assert.Error(t, err)
	assert.Nil(t, details)

	switch v := err.(type) {
	case *HttpError:
		assert.Equal(t, http.StatusUnauthorized, v.Code)
	default:
		// `HttpError` expected
		assert.True(t, false)
	}
}

func TestCheckValidityOfIBAN(t *testing.T) {
	conf := newTestsConfig()
	ctx := context.Background()

	respSchema, err := CheckValidityOfIBAN(ctx, realIBAN, conf.AppApiKey, SandBoxEnv)
	assert.Nil(t, err)
	assert.True(t, respSchema.Valid())
}

func TestGetDetailsForIBAN(t *testing.T) {
	conf := newTestsConfig()
	ctx := context.Background()

	components, err := GetDetailsForIBAN(ctx, realIBAN, conf.AppApiKey, SandBoxEnv)
	assert.Nil(t, err)
	assert.Equal(t, realIBAN, components.IBAN)
}

func TestGetBicByIBAN(t *testing.T) {
	conf := newTestsConfig()
	ctx := context.Background()

	gotBic, err := GetBicByIBAN(ctx, realIBAN, conf.AppApiKey, SandBoxEnv)
	assert.Nil(t, err)
	assert.Equal(t, realBic, gotBic)
}
