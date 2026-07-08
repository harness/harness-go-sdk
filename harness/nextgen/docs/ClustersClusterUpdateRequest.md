# ClustersClusterUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [***ClustersCluster**](clustersCluster.md) |  | [optional] [default to null]
**UpdatedFields** | **[]string** |  | [optional] [default to null]
**UpdateMask** | [***ProtobufFieldMask**](protobufFieldMask.md) |  | [optional] [default to null]
**Tags** | **map[string]string** |  | [optional] [default to null]
**Id** | [***ClustersClusterId**](clustersClusterID.md) |  | [optional] [default to null]
**SecretExpressions** | **map[string]string** | Maps credential field names (e.g. bearerToken, username, password) to Harness secret identifiers (e.g. account.my_secret). The agent resolves these at connect time instead of using plaintext credentials. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

