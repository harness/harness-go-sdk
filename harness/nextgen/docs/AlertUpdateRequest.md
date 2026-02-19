# AlertUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the alert to update. This field is required for update operations. | [default to null]
**Name** | **string** | Name of the alert | [default to null]
**Enabled** | **bool** | Whether the alert is enabled | [optional] [default to true]
**EntityType** | **string** | Type of entity the alert is associated with | [default to null]
**Recipients** | [**[]AlertRecipient**](AlertRecipient.md) | List of notification recipients. At least one recipient is required. | [default to null]
**Events** | [**[]AlertEvent**](AlertEvent.md) | List of events that trigger the alert. At least one event is required. | [default to null]
**AssociatedEntities** | [**[]AlertEntity**](AlertEntity.md) | List of entities associated with this alert. Use relation_type &#x27;all&#x27; to apply to all AutoStopping rules, or &#x27;specific&#x27; with entity_id to target specific rules. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

