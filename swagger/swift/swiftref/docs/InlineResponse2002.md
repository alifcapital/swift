# InlineResponse2002

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***Status**](Status.md) |  | [optional] [default to null]
**Iban** | **string** | The IBAN that was validated | [default to null]
**CountryCode** | **string** | The ISO 3166-1 alpha-2 code of the country of the institution/branch. | [default to null]
**Checksum** | **string** | The checksum extracted from the IBAN | [default to null]
**BankId** | **string** | The BANK ID extracted from the IBAN. Its structure (bank ID only or bank ID + branch ID) is as defined by the IBAN BIC LENGTH in the IBANSTRUCTURE file. | [default to null]
**BranchId** | **string** | Unique and unambiguous identification of a branch of a financial institution. | [optional] [default to null]
**AccountNumber** | **string** | The remainder of the IBAN including the account number. | [default to null]
**Length** | **int32** | The length of the IBAN | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

