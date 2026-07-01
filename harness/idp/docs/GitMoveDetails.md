# GitMoveDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | Name of the branch. When moving an inline pipeline or a remote pipeline, this branch is where the remote entity is created or fetched. | [default to null]
**FilePath** | **string** | File path of the entity in the repository. | [default to null]
**CommitMessage** | **string** | Merge commit message. | [optional] [default to null]
**BaseBranch** | **string** | Default branch name. This checks out the branch titled branch_name | [optional] [default to null]
**ConnectorRef** | **string** | Harness connector id used for entity CRUD operations | [default to null]
**RepoName** | **string** | Name of the repository. | [default to null]
**IsHarnessCodeRepo** | **bool** | Is Git Experience repo harness code. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

