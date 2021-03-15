# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBicsFromNationalIdV2**](NationalIdsApi.md#GetBicsFromNationalIdV2) | **Get** /v2/national_ids/{national_id}/bics | Get BICs of a National ID.
[**GetNationalIdDetailsV2**](NationalIdsApi.md#GetNationalIdDetailsV2) | **Get** /v2/national_ids/{national_id} | Get details of a National ID.
[**GetNationalIdValidityV2**](NationalIdsApi.md#GetNationalIdValidityV2) | **Get** /v2/national_ids/{national_id}/validity | Check the Validity of a National ID.

# **GetBicsFromNationalIdV2**
> InlineResponse20014 GetBicsFromNationalIdV2(ctx, nationalId, optional)
Get BICs of a National ID.

For a given National ID, you can retrieve the BIC or the BICs of that National ID; you need to provide the National ID, and either a country code, or an indication to which scheme this National ID belongs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nationalId** | **string**| National ID (a synonym for clearing code, routing code, sort code, bank and branch ID) for which the corresponding BICs are requested | 
 **optional** | ***NationalIdsApiGetBicsFromNationalIdV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NationalIdsApiGetBicsFromNationalIdV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryCode** | **optional.String**| The 2-character ISO 3166-1 country code of the country that issued the National ID (mandatory, when scheme parameter is not provided) | 
 **scheme** | **optional.String**| The scheme (a synonym for the National ID type) under which the National ID is defined (mandatory, when country_code parameter is not provided) | 
 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNationalIdDetailsV2**
> InlineResponse20012 GetNationalIdDetailsV2(ctx, nationalId, optional)
Get details of a National ID.

For a given National ID, you can retrieve the details of that National ID, that is the bank name or business name, and the address. You must provide either a country code, or an indication to which scheme the National ID belongs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nationalId** | **string**| National ID (a synonym for clearing code, routing code, sort code, bank and branch ID) for which details are requested | 
 **optional** | ***NationalIdsApiGetNationalIdDetailsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NationalIdsApiGetNationalIdDetailsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryCode** | **optional.String**| The 2-character ISO 3166-1 country code of the country that issued the National ID (mandatory, when scheme parameter is not provided) | 
 **scheme** | **optional.String**| The scheme (a synonym for the National ID type) under which the National ID is defined (mandatory, when country_code parameter is not provided) | 
 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 
 **onlyLocalLanguage** | **optional.Bool**| Whether fields must be returned in local language | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNationalIdValidityV2**
> InlineResponse20013 GetNationalIdValidityV2(ctx, nationalId, optional)
Check the Validity of a National ID.

With this API call, you can check whether a National ID format is correct and the value exists in SWIFT's list of National IDs; you need to provide the National ID, and either a country code, or an indication to which scheme this National ID belongs; in addition to v1, v2 of the API provides a country-specific validation of the National ID's format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nationalId** | **string**| National ID (a synonym for clearing code, routing code, sort code, bank and branch ID) to validate | 
 **optional** | ***NationalIdsApiGetNationalIdValidityV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NationalIdsApiGetNationalIdValidityV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryCode** | **optional.String**| The 2-character ISO 3166-1 country code of the country that issued the National ID (mandatory, when scheme parameter is not provided) | 
 **scheme** | **optional.String**| The scheme (a synonym for the National ID type) under which the National ID is defined (mandatory, when country_code parameter is not provided) | 
 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

