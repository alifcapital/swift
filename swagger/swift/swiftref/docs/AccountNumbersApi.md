# {{classname}}

All URIs are relative to *https://api.swift.com/swiftrefdata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountNumberValidityV2**](AccountNumbersApi.md#GetAccountNumberValidityV2) | **Get** /v2/account_numbers/{account_number}/validity | Check validity of account number formats, including IBANs.

# **GetAccountNumberValidityV2**
> InlineResponse20021 GetAccountNumberValidityV2(ctx, accountNumber, optional)
Check validity of account number formats, including IBANs.

Check validity of account number formats, including IBANs, issued in almost any country in the world

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountNumber** | **string**| Account number to validate | 
 **optional** | ***AccountNumbersApiGetAccountNumberValidityV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountNumbersApiGetAccountNumberValidityV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryCode** | **optional.String**| ISO 2-letter country code of the country where the account is held | 
 **isIban** | **optional.String**| Whether the account number must be handled as an IBAN | 
 **usage** | **optional.String**| How the account number is used (only for a few countries) | 
 **nationalId** | **optional.String**| National ID (bank ID, clearing code, sort code, routing number) of the account holding institution | 
 **bic** | **optional.String**| BIC-8 or BIC-11 of the account holding institution | 
 **aPIStatus** | **optional.String**| Whether status item must always be returned within responses (by default, positive responses do not report status) | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[oAuth2Password](../README.md#oAuth2Password)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

