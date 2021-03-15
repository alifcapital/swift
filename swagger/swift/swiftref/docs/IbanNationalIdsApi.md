# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBicFromIbanNationalIdV2**](IbanNationalIdsApi.md#GetBicFromIbanNationalIdV2) | **Get** /v2/iban_national_ids/{iban_national_id}/bic | Get the BIC for an IBAN national ID.

# **GetBicFromIbanNationalIdV2**
> InlineResponse2003 GetBicFromIbanNationalIdV2(ctx, ibanNationalId, countryCode, optional)
Get the BIC for an IBAN national ID.

For a given IBAN national ID, you can retrieve the BIC of the institution that services the IBAN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ibanNationalId** | **string**| IBAN national ID for which the corresponding BIC is requested | 
  **countryCode** | **string**| The 2-character ISO 3166-1 country code of the country that issued the IBAN National ID | 
 **optional** | ***IbanNationalIdsApiGetBicFromIbanNationalIdV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IbanNationalIdsApiGetBicFromIbanNationalIdV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

