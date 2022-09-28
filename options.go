package swift

import (
	"net/http"
)

// Option is a function that sets some option on API.
type Option func(api *API)

// HTTPClient creates new option for using custom http.Client
func HTTPClient(client *http.Client) Option {
	return func(api *API) {
		api.httpClient = client
	}
}
