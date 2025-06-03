# V1Taint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute. | [optional] [default to null]
**Key** | **string** | Required. The taint key to be applied to a node. | [optional] [default to null]
**TimeAdded** | **string** | TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints. +optional | [optional] [default to null]
**Value** | **string** | The taint value corresponding to the taint key. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

