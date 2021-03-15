# {{classname}}

All URIs are relative to *https://api-test.swiftnet.sipn.swift.com/swift-preval-pilot/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountFormatValidationStatus**](PaymentPreValidationApi.md#GetAccountFormatValidationStatus) | **Post** /payment/account-format | Validates the account format.
[**GetCategoryPurposeValidationStatus**](PaymentPreValidationApi.md#GetCategoryPurposeValidationStatus) | **Post** /payment/category-purpose | Validates the category purpose.
[**GetFinancialInstitutionIdentityValidationStatus**](PaymentPreValidationApi.md#GetFinancialInstitutionIdentityValidationStatus) | **Post** /payment/financial-institution-identity | Validates a financial institution identifier.
[**GetPaymentPurposeCodeValidationStatus**](PaymentPreValidationApi.md#GetPaymentPurposeCodeValidationStatus) | **Post** /payment/purpose-code | Validates the purpose code.
[**GetPaymentPurposeValidationStatus**](PaymentPreValidationApi.md#GetPaymentPurposeValidationStatus) | **Post** /payment/payment-purpose | Validates the payment purpose.
[**VerifyAccount**](PaymentPreValidationApi.md#VerifyAccount) | **Post** /accounts/verification | Verify that a beneficiary account could be able to receive incoming funds.

# **GetAccountFormatValidationStatus**
> AccountFormatValidationStatus1 GetAccountFormatValidationStatus(ctx, optional)
Validates the account format.

The service performs the format validation of the instructed account identifier against all the formats that apply to the instructed context (country or/ and the  financial institution that holds the account).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentPreValidationApiGetAccountFormatValidationStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentPreValidationApiGetAccountFormatValidationStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AccountFormatValidation1**](AccountFormatValidation1.md)|  | 

### Return type

[**AccountFormatValidationStatus1**](AccountFormatValidationStatus1.md)

### Authorization

[oauthBearerToken](../README.md#oauthBearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategoryPurposeValidationStatus**
> CategoryPurposeValidationStatus1 GetCategoryPurposeValidationStatus(ctx, optional)
Validates the category purpose.

This request validates the instructed category purpose code against the list of category purpose codes available in the SWIFT respository for the instructed payment context  (creditor country).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentPreValidationApiGetCategoryPurposeValidationStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentPreValidationApiGetCategoryPurposeValidationStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CategoryPurposeValidation1**](CategoryPurposeValidation1.md)|  | 

### Return type

[**CategoryPurposeValidationStatus1**](CategoryPurposeValidationStatus1.md)

### Authorization

[oauthBearerToken](../README.md#oauthBearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFinancialInstitutionIdentityValidationStatus**
> FinancialInstitutionIdentityValidationStatus1 GetFinancialInstitutionIdentityValidationStatus(ctx, optional)
Validates a financial institution identifier.

 Validates a financial institution identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentPreValidationApiGetFinancialInstitutionIdentityValidationStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentPreValidationApiGetFinancialInstitutionIdentityValidationStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FinancialInstitutionIdentityValidation1**](FinancialInstitutionIdentityValidation1.md)|  | 

### Return type

[**FinancialInstitutionIdentityValidationStatus1**](FinancialInstitutionIdentityValidationStatus1.md)

### Authorization

[oauthBearerToken](../README.md#oauthBearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentPurposeCodeValidationStatus**
> PaymentPurposeCodeValidationStatus1 GetPaymentPurposeCodeValidationStatus(ctx, optional)
Validates the purpose code.

This request validates the instructed information regarding the purpose of the payment ( purpose code, purpose description) based on  the specific requirements available in the SWFIT repository, for the instructed payment context  (creditor country,creditor account currency and  local payment instrument).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentPreValidationApiGetPaymentPurposeCodeValidationStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentPreValidationApiGetPaymentPurposeCodeValidationStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PaymentPurposeCodeValidation1**](PaymentPurposeCodeValidation1.md)|  | 

### Return type

[**PaymentPurposeCodeValidationStatus1**](PaymentPurposeCodeValidationStatus1.md)

### Authorization

[oauthBearerToken](../README.md#oauthBearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentPurposeValidationStatus**
> PaymentPurposeValidationStatus1 GetPaymentPurposeValidationStatus(ctx, optional)
Validates the payment purpose.

This request validates the instructed information regarding the purpose of the payment ( purpose code, purpose description) based on  the specific requirements available in the SWFIT repository, for the instructed payment context  (creditor country,creditor account currency and  local payment instrument).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentPreValidationApiGetPaymentPurposeValidationStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentPreValidationApiGetPaymentPurposeValidationStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PaymentPurposeValidation1**](PaymentPurposeValidation1.md)|  | 

### Return type

[**PaymentPurposeValidationStatus1**](PaymentPurposeValidationStatus1.md)

### Authorization

[oauthBearerToken](../README.md#oauthBearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyAccount**
> AccountVerificationResponse1 VerifyAccount(ctx, xBic, optional)
Verify that a beneficiary account could be able to receive incoming funds.

The service verifies that an account exists at the beneficiary          bank and is capable of receiving incoming funds. This usually implies that          the account is open, properly identified by the given number and, depending          on the jurisdiction and market practices in use where the account is held,          that the creditor name matches the name of the account holder. The service          provider does not take liability for the response and does not provide any          guarantee on the outcome of an actual transaction being sent to this account.          The information provided is meant to be as accurate as possible at the time          that the request was processed. The requester must pass the creditor name          and the service provider can use this information as part of the verification          or not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xBic** | **string**| Describe the BIC for SWIFT to route the request to. Providers get the value from the Gateway and consumers are not required to fill it in. | 
 **optional** | ***PaymentPreValidationApiVerifyAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentPreValidationApiVerifyAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AccountVerificationRequest**](AccountVerificationRequest.md)| Verify account details request. | 
 **xRequestID** | [**optional.Interface of string**](.md)| Specify an unique end to end tracking request ID header element. In each API request, the consumer must supply the ID. The value of the tracking ID MUST be a uuid as described in IETF RFC 4122 Universally Unique Identifier (UUID) URN Namespace. Pattern - ^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$. | 
 **xRequestDateTime** | **optional.**| The timestamp of the request, based on the client&#x27;s clock. Request date time must be created and supplied at the time of request. The date-time must be represented as specified in RFC 7231. For example - Sun, 06 Nov 1994 08:49:37 GMT. | 

### Return type

[**AccountVerificationResponse1**](AccountVerificationResponse1.md)

### Authorization

[oauthBearerToken](../README.md#oauthBearerToken), [oauthClientCredentials](../README.md#oauthClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

