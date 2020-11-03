package swift

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewHttpError_Error(t *testing.T) {
	testData := []struct {
		msg  string
		code int
	}{
		{
			msg:  "failed, no such service",
			code: http.StatusNotFound,
		},
		{
			msg:  `{"msg":"no found","code":500,"reportable":true}`,
			code: http.StatusInternalServerError,
		},
	}

	for _, v := range testData {
		httpErr := NewHttpError(v.code, v.msg, "application/json")
		expected, _ := json.Marshal(httpErr)

		assert.Equal(t, string(expected), httpErr.Error())
	}
}
