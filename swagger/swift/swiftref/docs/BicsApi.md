# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBicDetailsV2**](BicsApi.md#GetBicDetailsV2) | **Get** /v2/bics/{bic} | Get details of a BIC.
[**GetBicValidityV2**](BicsApi.md#GetBicValidityV2) | **Get** /v2/bics/{bic}/validity | Check validity of a BIC.
[**GetLeiFromBicV2**](BicsApi.md#GetLeiFromBicV2) | **Get** /v2/bics/{bic}/lei | Get the LEI for a BIC.
[**GetNationalIdsFromBicV2**](BicsApi.md#GetNationalIdsFromBicV2) | **Get** /v2/bics/{bic}/national_ids | Get National IDs for a BIC.
[**GetSepaReachabilityFromBicV2**](BicsApi.md#GetSepaReachabilityFromBicV2) | **Get** /v2/bics/{bic}/reachability | Validate SEPA reachability of a BIC.
[**GetSsisFromBicV2**](BicsApi.md#GetSsisFromBicV2) | **Get** /v2/bics/{bic}/ssis | Get SSIs for a BIC.

# **GetBicDetailsV2**
> InlineResponse2005 GetBicDetailsV2(ctx, bic, optional)
Get details of a BIC.

For a given BIC, you can retrieve the details of that BIC, that is the bank name or business name, and the address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bic** | **string**| BIC (8-characters or 11-characters) for which details are requested | 
 **optional** | ***BicsApiGetBicDetailsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BicsApiGetBicDetailsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBicValidityV2**
> InlineResponse2006 GetBicValidityV2(ctx, bic, optional)
Check validity of a BIC.

You can check whether a BIC is correct and valid, that is whether it is published in the BIC Directory; optionally, you can check whether a BIC was valid at a particular date in the past

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bic** | **string**| BIC (8-characters or 11-characters) to validate | 
 **optional** | ***BicsApiGetBicValidityV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BicsApiGetBicValidityV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **effectiveDate** | **optional.String**| A date in YYYY-MM-DD format | 
 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeiFromBicV2**
> InlineResponse2007 GetLeiFromBicV2(ctx, bic, optional)
Get the LEI for a BIC.

For a given BIC, you can retrieve the LEI of that BIC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bic** | **string**| BIC (8-characters or 11-characters) for which the corresponding LEI is requested | 
 **optional** | ***BicsApiGetLeiFromBicV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BicsApiGetLeiFromBicV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNationalIdsFromBicV2**
> InlineResponse2009 GetNationalIdsFromBicV2(ctx, bic, optional)
Get National IDs for a BIC.

For a given BIC, you can retrieve the National IDs of that BIC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bic** | **string**| BIC (8-characters or 11-characters) for which the corresponding National IDs are requested | 
 **optional** | ***BicsApiGetNationalIdsFromBicV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BicsApiGetNationalIdsFromBicV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSepaReachabilityFromBicV2**
> InlineResponse20010 GetSepaReachabilityFromBicV2(ctx, bic, sepaScheme, optional)
Validate SEPA reachability of a BIC.

For a given BIC and SEPA payment scheme, you can validate that it can be reached for SEPA payments and return the SEPA channel CSM (Clearing and Settlement System) through which the bank owning the BIC can be reached; the response also returns the adherence BIC that signed the adherence agreement with the EPC and the intermediary institution's BIC (if applicable and available)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bic** | **string**| BIC (8-characters or 11-characters) for which SEPA reachability is requested | 
  **sepaScheme** | **string**| SEPA service code | 
 **optional** | ***BicsApiGetSepaReachabilityFromBicV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BicsApiGetSepaReachabilityFromBicV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSsisFromBicV2**
> InlineResponse20011 GetSsisFromBicV2(ctx, bic, currencyCode, optional)
Get SSIs for a BIC.

For a given BIC, you can retrieve the Standing Settlement Instructions (SSIs) for that BIC, for a given currency, and (optionally) an SSI category (wholesale or retail).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bic** | **string**| BIC (8-characters or 11-characters) for which the corresponding SSIs are requested | 
  **currencyCode** | **string**| A 3-character currency code for which SSI data needs to be retrieved | 
 **optional** | ***BicsApiGetSsisFromBicV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BicsApiGetSsisFromBicV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ssiCategory** | **optional.String**| Code that indicates the SSI category for which SSI data needs to be retrieved, that is COPA or WHLS | 
 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

