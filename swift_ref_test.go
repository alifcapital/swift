package swift

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetDetailsOfBic_WithoutCredentials(t *testing.T) {
	conf := newTestsConfig()
	ctx := context.Background()

	// send request
	details, err := GetDetailsOfBic(ctx, conf.RealBic, "no-key", SandBoxEnv)
	assert.Error(t, err)
	assert.Nil(t, details)

	switch v := err.(type) {
	case *HttpError:
		assert.Equal(t, http.StatusUnauthorized, v.Code)
	default:
		// exactly `HttpError` expected
		assert.True(t, false)
	}
}

func TestGetDetailsOfBic(t *testing.T) {
	conf := newTestsConfig()
	ctx := context.Background()

	// send request
	details, err := GetDetailsOfBic(ctx, conf.RealBic, conf.XApiKey, SandBoxEnv)
	assert.Nil(t, err)
	assert.NotNil(t, details)
}
