# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIbanFromBbanV2**](BbansApi.md#GetIbanFromBbanV2) | **Get** /v2/bbans/{bban}/iban | Get the IBAN from a BBAN.

# **GetIbanFromBbanV2**
> InlineResponse2004 GetIbanFromBbanV2(ctx, bban, countryCode, optional)
Get the IBAN from a BBAN.

For a given Basic Bank Account Number (BBAN), you can retrieve the corresponding IBAN; you must also provide a country code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bban** | **string**| BBAN for which the corresponding IBAN is requested | 
  **countryCode** | **string**| The 2-character ISO 3166-1 country code of the country that issued the National ID | 
 **optional** | ***BbansApiGetIbanFromBbanV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BbansApiGetIbanFromBbanV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

