# Alert

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the alert | [optional] [default to null]
**AccountId** | **string** | ID of the account this alert belongs to | [optional] [default to null]
**Name** | **string** | Name of the alert | [optional] [default to null]
**Enabled** | **bool** | Whether the alert is enabled | [optional] [default to null]
**EntityType** | **string** | Type of entity the alert is associated with | [optional] [default to null]
**Recipients** | [**[]AlertRecipient**](AlertRecipient.md) | List of notification recipients | [optional] [default to null]
**Events** | [**[]AlertEvent**](AlertEvent.md) | List of events that trigger the alert | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when the alert was created | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when the alert was last updated | [optional] [default to null]
**AssociatedEntities** | [**[]AlertEntity**](AlertEntity.md) | List of entities associated with this alert | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

