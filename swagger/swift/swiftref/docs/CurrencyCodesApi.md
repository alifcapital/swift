# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrencyCodeDetailsV2**](CurrencyCodesApi.md#GetCurrencyCodeDetailsV2) | **Get** /v2/currency_codes/{currency_code} | Get details of a currency code.
[**GetCurrencyCodeValidityV2**](CurrencyCodesApi.md#GetCurrencyCodeValidityV2) | **Get** /v2/currency_codes/{currency_code}/validity | Check validity of a currency code.

# **GetCurrencyCodeDetailsV2**
> InlineResponse20018 GetCurrencyCodeDetailsV2(ctx, currencyCode, optional)
Get details of a currency code.

For a given currency code, you can retrieve the details of that currency code, that is the name, other codes, the fractional digits and the list of countries where the currency is used

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyCode** | **string**| Currency code (3-letters or 3-digits) for which details are requested | 
 **optional** | ***CurrencyCodesApiGetCurrencyCodeDetailsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrencyCodesApiGetCurrencyCodeDetailsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrencyCodeValidityV2**
> InlineResponse20017 GetCurrencyCodeValidityV2(ctx, currencyCode, optional)
Check validity of a currency code.

You can check whether a currency code is correct and valid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currencyCode** | **string**| Currency code (3-letters or 3-digits) to validate | 
 **optional** | ***CurrencyCodesApiGetCurrencyCodeValidityV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CurrencyCodesApiGetCurrencyCodeValidityV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

