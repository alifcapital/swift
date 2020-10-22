package swift

import (
	"fmt"
	"strings"
)

// SWIFTError - since all SWIFT apis are RESTFul
// this Error implementation should come in handy for processing different network errors
type SWIFTError struct {
	// http.StatusCode
	Code int
	// response body
	Message string
	// contentType
	ContentType string
}

// NewSWIFTResponseError - factory function
func NewSWIFTError(code int, message, contentType string) *SWIFTError {
	return &SWIFTError{
		Code:        code,
		Message:     message,
		ContentType: contentType,
	}
}

func (swiftErr *SWIFTError) Error() string {
	// when content-type is a json, format the message as json string
	// this is done because json.Marshall escapes double quotes
	if strings.Contains(swiftErr.ContentType, "/json") {
		return fmt.Sprintf(`{"code":%d,"message":%s}`, swiftErr.Code, swiftErr.Message)
	}
	// otherwise return the message as enquoted string
	return fmt.Sprintf(`{"code":%d,"message":"%s"}`, swiftErr.Code, swiftErr.Message)
}
