# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBicFromIbanV2**](IbansApi.md#GetBicFromIbanV2) | **Get** /v2/ibans/{iban}/bic | Get the BIC for an IBAN.
[**GetIbanDetailsV2**](IbansApi.md#GetIbanDetailsV2) | **Get** /v2/ibans/{iban} | Get details for an IBAN.
[**GetIbanValidityV2**](IbansApi.md#GetIbanValidityV2) | **Get** /v2/ibans/{iban}/validity | Check validity of an IBAN.

# **GetBicFromIbanV2**
> InlineResponse2001 GetBicFromIbanV2(ctx, iban, optional)
Get the BIC for an IBAN.

For a given IBAN, you can retrieve the BIC of the institution that services the IBAN. This API is compliant with the requirement for BIC from IBAN derivation defined in the Regulation (EU) No 260/2012.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iban** | **string**| IBAN for which the corresponding BIC is requested | 
 **optional** | ***IbansApiGetBicFromIbanV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IbansApiGetBicFromIbanV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIbanDetailsV2**
> InlineResponse2002 GetIbanDetailsV2(ctx, iban, optional)
Get details for an IBAN.

For a given IBAN, you can obtain the components of the IBAN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iban** | **string**| IBAN for which details are requested | 
 **optional** | ***IbansApiGetIbanDetailsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IbansApiGetIbanDetailsV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIbanValidityV2**
> InlineResponse200 GetIbanValidityV2(ctx, iban, optional)
Check validity of an IBAN.

You can check whether an IBAN is valid, that is its country code, structure, length, and checksum are valid. It also checks that the bank ID exists and that it is allowed for use in the IBANs. If the bank ID is not allowed, then it is listed in the Exclusion List. This list contains bank IDs that must not be used in IBANs. This validity call allows for very fast, low overhead checking of the validity of an IBAN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iban** | **string**| IBAN to validate | 
 **optional** | ***IbansApiGetIbanValidityV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IbansApiGetIbanValidityV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

