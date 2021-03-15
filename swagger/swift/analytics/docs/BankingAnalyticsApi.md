# {{classname}}

All URIs are relative to *https://api.swift.com/bi/banking-analytics/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BankingAnalytics**](BankingAnalyticsApi.md#BankingAnalytics) | **Get** /value-and-currency | Banking Data.

# **BankingAnalytics**
> InlineResponse200 BankingAnalytics(ctx, market, reportingPeriod, optional)
Banking Data.

Returns an array of banking data filtered on market and reporting period. Each item in the array is defined by its unique combination of the following properties - market, reporting_period, message_definition_identifier, my_institution, counterparty and currency.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **market** | [**Market1Code**](.md)| Specifies the market to which the banking data belongs. | 
  **reportingPeriod** | [**string**](.md)| specifies the month the reported data covers. | 
 **optional** | ***BankingAnalyticsApiBankingAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BankingAnalyticsApiBankingAnalyticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | [**optional.Interface of int32**](.md)| Limits the maximum number of results to be displayed on a page with a minimum of 25 (putting less defaults to 25) and a maximum of 1000 (putting more defaults to 1000). | 
 **offset** | [**optional.Interface of int32**](.md)| Specifies the number of items from which to start returning results. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oauthPassword](../README.md#oauthPassword)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

