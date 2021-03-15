# InlineResponse2005

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***Status**](Status.md) |  | [optional] [default to null]
**Bic** | **string** | The BIC of the institution on which more information is requested | [default to null]
**InstitutionName** | **string** | Name by which a party is known and which is usually used to identify that party | [default to null]
**BranchInformation** | **string** | A free text description of the branch as provided by the financial institution to which it belongs; this is currently provided for entries with a BIC when the financial institution concerned wants to provide this extra information | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [default to null]
**ContactDetails** | [***ContactDetails**](ContactDetails.md) |  | [optional] [default to null]
**OfficeType** | **string** | Status of the entity in the office hierarchy | [optional] [default to null]
**SwiftServices** | [**[]SwiftService**](SwiftService.md) | A list of 3-character codes and names of SWIFT FIN service codes (also called value added service codes) | [default to null]
**SwiftConnectivity** | [***SwiftConnectivity**](SwiftConnectivity.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

