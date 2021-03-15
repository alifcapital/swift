# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountryCodeDetailsV2**](CountryCodesApi.md#GetCountryCodeDetailsV2) | **Get** /v2/country_codes/{country_code} | Get details of a country code.
[**GetCountryCodeValidityV2**](CountryCodesApi.md#GetCountryCodeValidityV2) | **Get** /v2/country_codes/{country_code}/validity | Check validity of a country code.

# **GetCountryCodeDetailsV2**
> InlineResponse20016 GetCountryCodeDetailsV2(ctx, countryCode, optional)
Get details of a country code.

For a given country code, you can retrieve the details of that country code, that is the name, other codes and official currencies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| Country code (2-letters, 3-letters or 3-digits) for which details are requested | 
 **optional** | ***CountryCodesApiGetCountryCodeDetailsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CountryCodesApiGetCountryCodeDetailsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryCodeValidityV2**
> InlineResponse20015 GetCountryCodeValidityV2(ctx, countryCode, optional)
Check validity of a country code.

You can check whether a country code is correct and valid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| Country code (2-letters, 3-letters or 3-digits) to validate | 
 **optional** | ***CountryCodesApiGetCountryCodeValidityV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CountryCodesApiGetCountryCodeValidityV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

