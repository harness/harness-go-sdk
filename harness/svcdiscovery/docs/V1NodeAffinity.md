# V1NodeAffinity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferredDuringSchedulingIgnoredDuringExecution** | [**[]V1PreferredSchedulingTerm**](v1.PreferredSchedulingTerm.md) | The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \&quot;weight\&quot; to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred. +optional | [optional] [default to null]
**RequiredDuringSchedulingIgnoredDuringExecution** | [***V1NodeSelector**](v1.NodeSelector.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

