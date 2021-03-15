package swift

import (
	"context"
	"github.com/alifcapital/swift/swagger/swift/analytics"
	"github.com/alifcapital/swift/swagger/swift/swiftref"
)

func (api *API) BankingAnalytics(ctx context.Context, code analytics.Market1Code, reportingPeriod string) (*analytics.InlineResponse200, error) {
	r, httpResponse, err := api.Analytics.BankingAnalyticsApi.
		BankingAnalytics(api.WithContext(ctx, swiftref.ContextOAuth2), code, reportingPeriod, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}
