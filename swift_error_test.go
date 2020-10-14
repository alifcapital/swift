package swift_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewSWIFTErrorWithString(t *testing.T) {
	message := "Bad gateWay given"
	expectedErr := fmt.Sprintf(`{"code":%d,"message":"%s"}`, http.StatusBadGateway, message)

	swiftError := NewSWIFTError(http.StatusBadGateway, message, "")
	assert.Equal(t, expectedErr, swiftError.Error())
}

func TestNewSWIFTErrorWithJsonString(t *testing.T) {
	dummy := struct {
		HttpMessage string `json:"http_message"`
		HttpCode    int    `json:"http_code"`
		Processed   bool   `json:"processed"`
	}{
		HttpMessage: "failed, no such service",
		HttpCode:    http.StatusNotFound,
		Processed:   false,
	}
	expectedErr := fmt.Sprintf(`{"code":%d,"message":%s}`, http.StatusBadGateway, `{"http_message":"failed, no such service","http_code":404,"processed":false}`)

	message, _ := json.Marshal(&dummy)
	swiftError := NewSWIFTError(http.StatusBadGateway, string(message), "application/json")
	assert.Equal(t, expectedErr, swiftError.Error())
}

func TestNewSWIFTErrorWithArrayOfJson(t *testing.T) {
	dummies := []struct {
		HttpMessage string `json:"http_message"`
		HttpCode    int    `json:"http_code"`
		Processed   bool   `json:"processed"`
	}{
		{
			HttpMessage: "failed, no such service",
			HttpCode:    http.StatusNotFound,
			Processed:   false,
		},
		{
			HttpMessage: "failed, no such service",
			HttpCode:    http.StatusNotFound,
			Processed:   false,
		},
	}
	jsonStr := fmt.Sprintf(`[{"http_message":"failed, no such service","http_code":404,"processed":false},{"http_message":"failed, no such service","http_code":404,"processed":false}]`)
	expectedErr := fmt.Sprintf(`{"code":%d,"message":%s}`, http.StatusBadGateway, jsonStr)

	message, _ := json.Marshal(&dummies)
	swiftError := NewSWIFTError(http.StatusBadGateway, string(message), "text/json")
	assert.Equal(t, expectedErr, swiftError.Error())
}
