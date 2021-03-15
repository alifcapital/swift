# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBicFromLeiV2**](LeisApi.md#GetBicFromLeiV2) | **Get** /v2/leis/{lei}/bic | Get the BIC for an LEI.
[**GetLeiDetailsV2**](LeisApi.md#GetLeiDetailsV2) | **Get** /v2/leis/{lei} | Get details of a LEI.
[**GetLeiValidityV2**](LeisApi.md#GetLeiValidityV2) | **Get** /v2/leis/{lei}/validity | Check validity of a LEI.

# **GetBicFromLeiV2**
> InlineResponse2008 GetBicFromLeiV2(ctx, lei, optional)
Get the BIC for an LEI.

For a given LEI, you can retrieve the BIC of that LEI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lei** | **string**| LEI (20-characters) for which the corresponding BIC is requested | 
 **optional** | ***LeisApiGetBicFromLeiV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeisApiGetBicFromLeiV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeiDetailsV2**
> InlineResponse20020 GetLeiDetailsV2(ctx, lei, optional)
Get details of a LEI.

For a given LEI, you can retrieve the details of that LEI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lei** | **string**| LEI for which details are requested | 
 **optional** | ***LeisApiGetLeiDetailsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeisApiGetLeiDetailsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeiValidityV2**
> InlineResponse20019 GetLeiValidityV2(ctx, lei, optional)
Check validity of a LEI.

You can check whether a LEI is correct and valid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lei** | **string**| LEI to validate | 
 **optional** | ***LeisApiGetLeiValidityV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeisApiGetLeiValidityV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

