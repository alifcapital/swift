package swift

import (
	"context"
	"github.com/alifcapital/swift/swagger/swift/swiftref"
	"github.com/antihax/optional"
)

func (api *API) GetBicDetailsV5(ctx context.Context, BIC string) (*swiftref.GetBicDetailsV5Response, error) {

	r, httpResponse, err := api.Reference.BicsApi.
		GetBicDetailsV5(api.WithContext(ctx, swiftref.ContextOAuth2), BIC, &swiftref.BicsApiGetBicDetailsV2Opts{APIStatus: optional.NewString("APIStatus")})
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetBicDetails(ctx context.Context, BIC string) (*swiftref.InlineResponse2005, error) {

	r, httpResponse, err := api.Reference.BicsApi.
		GetBicDetailsV2(api.WithContext(ctx, swiftref.ContextOAuth2), BIC, &swiftref.BicsApiGetBicDetailsV2Opts{APIStatus: optional.NewString("APIStatus")})
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetBicValidity(ctx context.Context, BIC string) (*swiftref.InlineResponse2006, error) {
	r, httpResponse, err := api.Reference.BicsApi.
		GetBicValidityV2(api.WithContext(ctx, swiftref.ContextOAuth2), BIC, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetLeiFromBic(ctx context.Context, BIC string) (*swiftref.InlineResponse2007, error) {
	r, httpResponse, err := api.Reference.BicsApi.
		GetLeiFromBicV2(api.WithContext(ctx, swiftref.ContextOAuth2), BIC, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetNationalIdsFromBic(ctx context.Context, BIC string) (*swiftref.InlineResponse2009, error) {
	r, httpResponse, err := api.Reference.BicsApi.
		GetNationalIdsFromBicV2(api.WithContext(ctx, swiftref.ContextOAuth2), BIC, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetSepaReachabilityFromBic(ctx context.Context, BIC, sepaScheme string) (*swiftref.InlineResponse20010, error) {
	r, httpResponse, err := api.Reference.BicsApi.
		GetSepaReachabilityFromBicV2(api.WithContext(ctx, swiftref.ContextOAuth2), BIC, sepaScheme, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetSsisFromBic(ctx context.Context, BIC, currencyCode string) (*swiftref.InlineResponse20011, error) {
	r, httpResponse, err := api.Reference.BicsApi.
		GetSsisFromBicV2(api.WithContext(ctx, swiftref.ContextOAuth2), BIC, currencyCode, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetAccountNumberValidity(ctx context.Context, accountNumber string) (*swiftref.InlineResponse20021, error) {
	r, httpResponse, err := api.Reference.AccountNumbersApi.
		GetAccountNumberValidityV2(api.WithContext(ctx, swiftref.ContextOAuth2), accountNumber, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetIbanFromBban(ctx context.Context, BBAN, countryCode string) (*swiftref.InlineResponse2004, error) {
	r, httpResponse, err := api.Reference.BbansApi.
		GetIbanFromBbanV2(api.WithContext(ctx, swiftref.ContextOAuth2), BBAN, countryCode, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetCountryCodeDetails(ctx context.Context, countryCode string) (*swiftref.InlineResponse20016, error) {
	r, httpResponse, err := api.Reference.CountryCodesApi.
		GetCountryCodeDetailsV2(api.WithContext(ctx, swiftref.ContextOAuth2), countryCode, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetCountryCodeValidity(ctx context.Context, countryCode string) (*swiftref.InlineResponse20015, error) {
	r, httpResponse, err := api.Reference.CountryCodesApi.
		GetCountryCodeValidityV2(api.WithContext(ctx, swiftref.ContextOAuth2), countryCode, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetCurrencyCodeDetails(ctx context.Context, countryCode string) (*swiftref.InlineResponse20018, error) {
	r, httpResponse, err := api.Reference.CurrencyCodesApi.
		GetCurrencyCodeDetailsV2(api.WithContext(ctx, swiftref.ContextOAuth2), countryCode, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetCurrencyCodeValidity(ctx context.Context, currencyCode string) (*swiftref.InlineResponse20017, error) {
	r, httpResponse, err := api.Reference.CurrencyCodesApi.
		GetCurrencyCodeValidityV2(api.WithContext(ctx, swiftref.ContextOAuth2), currencyCode, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetBicFromIbanNationalId(ctx context.Context, ibanNationalID, countryCode string) (*swiftref.InlineResponse2003, error) {
	r, httpResponse, err := api.Reference.IbanNationalIdsApi.
		GetBicFromIbanNationalIdV2(api.WithContext(ctx, swiftref.ContextOAuth2), ibanNationalID, countryCode, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetBicFromIban(ctx context.Context, iban string) (*swiftref.InlineResponse2001, error) {
	r, httpResponse, err := api.Reference.IbansApi.
		GetBicFromIbanV2(api.WithContext(ctx, swiftref.ContextOAuth2), iban, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetIbanDetails(ctx context.Context, iban string) (*swiftref.InlineResponse2002, error) {
	r, httpResponse, err := api.Reference.IbansApi.
		GetIbanDetailsV2(api.WithContext(ctx, swiftref.ContextOAuth2), iban, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetIbanValidity(ctx context.Context, iban string) (*swiftref.InlineResponse200, error) {
	r, httpResponse, err := api.Reference.IbansApi.
		GetIbanValidityV2(api.WithContext(ctx, swiftref.ContextOAuth2), iban, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetBicFromLei(ctx context.Context, lei string) (*swiftref.InlineResponse2008, error) {
	r, httpResponse, err := api.Reference.LeisApi.
		GetBicFromLeiV2(api.WithContext(ctx, swiftref.ContextOAuth2), lei, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetLeiDetails(ctx context.Context, lei string) (*swiftref.InlineResponse20020, error) {
	r, httpResponse, err := api.Reference.LeisApi.
		GetLeiDetailsV2(api.WithContext(ctx, swiftref.ContextOAuth2), lei, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetLeiValidity(ctx context.Context, lei string) (*swiftref.InlineResponse20019, error) {
	r, httpResponse, err := api.Reference.LeisApi.
		GetLeiValidityV2(api.WithContext(ctx, swiftref.ContextOAuth2), lei, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetBicsFromNationalId(ctx context.Context, nationalID string) (*swiftref.InlineResponse20014, error) {
	r, httpResponse, err := api.Reference.NationalIdsApi.
		GetBicsFromNationalIdV2(api.WithContext(ctx, swiftref.ContextOAuth2), nationalID, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetNationalIdDetails(ctx context.Context, nationalID string) (*swiftref.InlineResponse20012, error) {
	r, httpResponse, err := api.Reference.NationalIdsApi.
		GetNationalIdDetailsV2(api.WithContext(ctx, swiftref.ContextOAuth2), nationalID, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}

func (api *API) GetNationalIdValidity(ctx context.Context, nationalID string) (*swiftref.InlineResponse20013, error) {
	r, httpResponse, err := api.Reference.NationalIdsApi.
		GetNationalIdValidityV2(api.WithContext(ctx, swiftref.ContextOAuth2), nationalID, nil)
	return &r, NewErrWithHTTPResponse(httpResponse, err)
}
