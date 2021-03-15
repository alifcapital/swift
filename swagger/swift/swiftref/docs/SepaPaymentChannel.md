# SepaPaymentChannel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the payment channel (SEPA-compliant CSM) | [default to null]
**MemberType** | **string** | A payment channel has a member_type, the member_type has one of the codes DRCT, IDRC or UKWN | [default to null]
**DirectParticipant** | **string** | The BIC of the direct participant through which the indirectly reachable BIC can be reached; it may only be present if the payment channel&#x27;s member is an indirect participant | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

