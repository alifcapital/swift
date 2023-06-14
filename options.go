package swift

import (
	"crypto/tls"
	"net/http"
)

// Option is a function that sets some option on API.
type Option func(api *API)

// HTTPClient creates new option for using custom http.Client
func HTTPClient(client *http.Client) Option {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Transport = tr
	return func(api *API) {
		api.httpClient = client
	}
}
