# V1GitRepoVolumeSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | **string** | directory is the target directory name. Must not contain or start with &#x27;..&#x27;.  If &#x27;.&#x27; is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name. +optional | [optional] [default to null]
**Repository** | **string** | repository is the URL | [optional] [default to null]
**Revision** | **string** | revision is the commit hash for the specified revision. +optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

