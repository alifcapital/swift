package swift

import (
	"encoding/json"
)

// HttpError - since all SWIFT apis are RESTFul
// this Error implementation should come in handy for processing different network errors
type HttpError struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	ContentType string `json:"content_type"`
}

// NewHttpError - factory
func NewHttpError(code int, message, contentType string) *HttpError {
	return &HttpError{
		Code:        code,
		Message:     message,
		ContentType: contentType,
	}
}

func (err *HttpError) Error() string {
	errB, _ := json.Marshal(err)
	return string(errB)
}
