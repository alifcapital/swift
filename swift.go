// Package swift implements SWIFT API

/*
	Usage

	Create API by NewAPI and call it's methods. See more about usage
	in NewAPI's docs.

	You have two ways of using API.
	First is by calling API's method:

			swiftAPI := swift.NewAPI(...)
			swiftAPI.GetBicDetails(ctx, bic)

	This way is more convenient and less verbose. In most cases that's
	enough, but by using this approach you can provide only required
	arguments in SWIFT API. If you need to provide optional arguments
	you can use second way:

			swiftAPI := swift.NewAPI(...)
			optionalArgs := swiftref.BicsApiGetBicDetailsV2Opts{APIStatus: optional.NewString("APIStatus")}
			swiftAPI.Reference.BicsApi.GetBicDetailsV2(swiftAPI.WithContext(ctx, swiftref.ContextOAuth2), bic, &optionalArgs)

	More verbose, but now you can specify optional arguments to SWIFT API.
	See more about this approach in API.WithContext docs.

*/
package swift

import (
	"context"
	"github.com/alifcapital/swift/swagger/swift/analytics"
	"github.com/alifcapital/swift/swagger/swift/prevalidation"
	"github.com/alifcapital/swift/swagger/swift/swiftref"
	"golang.org/x/oauth2"
	"net/http"
)

type env int

const (
	// Sandbox is a testing environment, but can be used in production for some services
	// from docs:  ... KYC Registry API, Banking Analytics API, Banking Premium API and Compliance Analytics API, Sandbox CAN be used for UAT.
	Sandbox = env(iota)
	// Production is a production environment
	Production
)

// API aggregates swiftref, bank-analytics APIs and deals with
// authentication (gets access token as needed)
type API struct {
	basePath         string
	credentials      AppCredentials
	httpClient       *http.Client
	oauth2Token      *oauth2.Token
	reuseTokenSource oauth2.TokenSource
	Reference        *swiftref.APIClient
	Analytics        *analytics.APIClient
	PreValidation    *prevalidation.APIClient
}

// NewAPI return new API based on given args and options
func NewAPI(environment env, credentials AppCredentials, options ...Option) *API {
	api := API{credentials: credentials}
	for _, o := range options {
		o(&api)
	}
	if api.httpClient == nil {
		api.httpClient = http.DefaultClient
	}
	reuseTokenSource := oauth2.ReuseTokenSource(nil, &api)
	api.reuseTokenSource = reuseTokenSource
	api.setEnvironment(environment)
	conf := swiftref.Configuration{
		BasePath:      api.basePath + "/swiftrefdata",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		HTTPClient:    api.httpClient,
	}
	analyticsConf := analytics.Configuration{
		BasePath:      api.basePath + "/bi/banking-analytics/v1", //"https://api-pilot.swift.com"
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		HTTPClient:    api.httpClient,
	}
	api.Reference = swiftref.NewAPIClient(&conf)
	api.Analytics = analytics.NewAPIClient(&analyticsConf)

	return &api
}

// WithContext wraps ctx with authentication information.
func (api *API) WithContext(ctx context.Context, key interface{}) context.Context {
	return context.WithValue(ctx, key, api.reuseTokenSource)
}

func (api *API) setEnvironment(environment env) {
	switch environment {
	case Sandbox:
		api.basePath = "https://sandbox.swift.com"
	case Production:
		api.basePath = "https://api.swift.com"
	}
}
