# AllOfv1VolumeEmptyDir

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Medium** | [***interface{}**](interface{}.md) | What type of storage medium should back this directory. The default is \&quot;\&quot; which means to use the node&#x27;s default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir +optional | [optional] [default to null]
**SizeLimit** | [***interface{}**](interface{}.md) | Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

