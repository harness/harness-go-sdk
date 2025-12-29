# InfrastructureExecution

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | [**[]InstanceApproval**](InstanceApproval.md) |  | [optional] [default to null]
**CompletedInstanceIds** | **[]string** | Instances fully updated by the execution | [default to null]
**CreatedBy** | [***UserInfo**](UserInfo.md) |  | [optional] [default to null]
**Events** | [**[]ExecutionEvent**](ExecutionEvent.md) |  | [optional] [default to null]
**Id** | **string** |  | [default to null]
**InfrastructureId** | **string** |  | [default to null]
**InstanceIds** | **[]string** | Instances updated by the execution | [default to null]
**Message** | **string** | Message associated with the execution | [optional] [default to null]
**Outputs** | [**map[string]ExecutionOutput**](ExecutionOutput.md) | Outputs of the infrastructure execution | [optional] [default to null]
**PipelineUrl** | **string** |  | [optional] [default to null]
**Progress** | [***InfrastructureProgress**](InfrastructureProgress.md) |  | [default to null]
**Spec** | [***Infrastructure**](Infrastructure.md) |  | [default to null]
**SpecRevision** | **int64** |  | [default to null]
**StartedAt** | **string** | DEPRECATED: Use startedAtMs instead. RFC3339 timestamp when execution started | [default to null]
**StartedAtMs** | **int64** | Unix timestamp in milliseconds when execution started | [default to null]
**StoppedAt** | **string** | DEPRECATED: Use stoppedAtMs instead. RFC3339 timestamp when execution stopped | [optional] [default to null]
**StoppedAtMs** | **int64** | Unix timestamp in milliseconds when execution stopped | [optional] [default to null]
**TargetState** | [***State**](State.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

