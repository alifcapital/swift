package swift

import (
	"net/http"
)

func NotFound(err error) bool {
	httpError, ok := err.(*ErrWithHTTPResponse)
	if !ok {
		return false
	}
	if !httpError.NotFound() {
		return false
	}
	return httpError.NotFound()
}

type ErrWithHTTPResponse struct {
	HTTPResponse *http.Response
	Err          error
}

func NewErrWithHTTPResponse(HTTPResponse *http.Response, err error) error {
	if err == nil {
		return nil
	}
	return &ErrWithHTTPResponse{HTTPResponse: HTTPResponse, Err: err}
}

func (e *ErrWithHTTPResponse) NotFound() bool {
	return e.HTTPResponse.StatusCode == http.StatusNotFound
}

func (e *ErrWithHTTPResponse) Unwrap() error {
	return e.Err
}

func (e *ErrWithHTTPResponse) Error() string {
	return "http status code " + e.HTTPResponse.Status + ": " + e.Err.Error()
}
