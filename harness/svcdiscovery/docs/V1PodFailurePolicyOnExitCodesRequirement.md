# V1PodFailurePolicyOnExitCodesRequirement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** | Restricts the check for exit codes to the container with the specified name. When null, the rule applies to all containers. When specified, it should match one the container or initContainer names in the pod template. +optional | [optional] [default to null]
**Operator** | **string** | Represents the relationship between the container exit code(s) and the specified values. Containers completed with success (exit code 0) are excluded from the requirement check. Possible values are: - In: the requirement is satisfied if at least one container exit code   (might be multiple if there are multiple containers not restricted   by the &#x27;containerName&#x27; field) is in the set of specified values. - NotIn: the requirement is satisfied if at least one container exit code   (might be multiple if there are multiple containers not restricted   by the &#x27;containerName&#x27; field) is not in the set of specified values. Additional values are considered to be added in the future. Clients should react to an unknown operator by assuming the requirement is not satisfied. | [optional] [default to null]
**Values** | **[]int32** | Specifies the set of values. Each returned container exit code (might be multiple in case of multiple containers) is checked against this set of values with respect to the operator. The list of values must be ordered and must not contain duplicates. Value &#x27;0&#x27; cannot be used for the In operator. At least one element is required. At most 255 elements are allowed. +listType&#x3D;set | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

