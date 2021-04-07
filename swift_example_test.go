package swift_test

import (
	"context"
	"github.com/alifcapital/swift"
	"github.com/alifcapital/swift/swagger/swift/analytics"
	"github.com/alifcapital/swift/swagger/swift/swiftref"
	"github.com/antihax/optional"
	"log"
	"net/http"
)

func ExampleNewAPI() {
	swiftAPI := swift.NewAPI(swift.Sandbox, swift.AppCredentials{
		BasicAuthUser: "some_user",
		BasicAuthPass: "some_pass",
		Username:      "some_username",
		Password:      "some_password",
	}, swift.HTTPClient(http.DefaultClient))

	bicDetails, err := swiftAPI.GetBicDetails(context.Background(), "DEUTDEFF")
	if err != nil {
		switch {
		case swift.NotFound(err):
			log.Println("no such bic")
		case swift.UnAuthorized(err):
			log.Println("unauthorized")
		default:
			log.Println("getting bic details:", err)
		}
		return
	}
	_ = bicDetails
	// Output:
	//
}

func ExampleAPI_WithContext() {
	// swiftref:
	swiftAPI := swift.NewAPI(swift.Sandbox, swift.AppCredentials{
		BasicAuthUser: "some_user",
		BasicAuthPass: "some_pass",
		Username:      "some_username",
		Password:      "some_password",
	})
	ctx := context.Background()
	bicDetails, httpResponse, err := swiftAPI.Reference.BicsApi.
		GetBicDetailsV2(swiftAPI.WithContext(ctx, swiftref.ContextOAuth2), "DEUTDEFF", nil)
	if err != nil {
		// ...
	}
	_, _ = bicDetails, httpResponse

	// bank analytics:
	bankingAnalyticsOpts := analytics.BankingAnalyticsApiBankingAnalyticsOpts{
		Limit:  optional.NewInterface(10),
		Offset: optional.NewInterface(0),
	}

	bankingAnalytics, httpResponse, err := swiftAPI.Analytics.BankingAnalyticsApi.
		BankingAnalytics(swiftAPI.WithContext(ctx, analytics.ContextOAuth2), "PMTS_payments", "2018-08", &bankingAnalyticsOpts)
	if err != nil {
		// ...
	}
	_, _ = bankingAnalytics, httpResponse
	// Output:
	//
}
