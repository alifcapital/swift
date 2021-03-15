# Ssi

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerBic** | **string** | The owner of the SSI Nostro Account | [default to null]
**InstitutionName** | **string** | Name by which an institution is known and which is usually used to identify that institution | [optional] [default to null]
**BranchInformation** | **string** | A free text description of the branch as provided by the financial institution to which it belongs; for the time being this will be provided only for entries with a BIC and only when the financial institution concerned wants to provide this extra information; the information is sourced from the BIC Directory | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**CurrencyCode** | **string** | The ISO 4217 currency code of the requested SSIs | [default to null]
**Direct** | **bool** | Indicates whether there is a direct account relationship between the owner of the SSI and the correspondent; if false, then this means at least the first intermediary must be present | [default to null]
**Correspondent** | [***CorrespondentBic**](CorrespondentBic.md) |  | [default to null]
**FirstIntermediary** | [***IntermediaryBic**](IntermediaryBic.md) |  | [optional] [default to null]
**SecondIntermediary** | [***IntermediaryBic**](IntermediaryBic.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

