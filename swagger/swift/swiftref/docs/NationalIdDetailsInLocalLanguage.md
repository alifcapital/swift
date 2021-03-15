# NationalIdDetailsInLocalLanguage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | National ID for which details were requested | [default to null]
**Scheme** | **string** | The scheme under which the National ID is defined | [default to null]
**InstitutionName** | **string** | Name by which a party is known and which is usually used to identify that party | [default to null]
**BranchInformation** | **string** | A free text description of the branch as provided by the financial institution to which it belongs; this is currently provided for entries with a BIC when the financial institution concerned wants to provide this extra information | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**ContactDetails** | [***ContactDetails**](ContactDetails.md) |  | [optional] [default to null]
**OfficeType** | **string** | Status of the entity in the office hierarchy | [optional] [default to null]
**LanguageIsoCode3** | **string** | ISO 639-2 code (three letters) of language | [optional] [default to null]
**LanguageScriptCode4** | **string** | ISO 15924 code (four letters) of language script | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

