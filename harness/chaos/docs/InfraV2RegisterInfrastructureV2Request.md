# InfraV2RegisterInfrastructureV2Request

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | **map[string]string** |  | [optional] [default to null]
**CorrelationId** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**DiscoveryAgentID** | **string** |  | [optional] [default to null]
**Env** | [**[]V1EnvVar**](v1.EnvVar.md) |  | [optional] [default to null]
**EnvFrom** | [**[]V1EnvFromSource**](v1.EnvFromSource.md) |  | [optional] [default to null]
**EnvironmentID** | **string** |  | [optional] [default to null]
**Identity** | **string** |  | [optional] [default to null]
**InfraID** | **string** |  | [optional] [default to null]
**InfraNamespace** | **string** |  | [optional] [default to null]
**InfraScope** | [***InfraV2InfraScope**](infra_v2.InfraScope.md) |  | [optional] [default to null]
**InfraType** | [***InfraV2InfraType**](infra_v2.InfraType.md) |  | [optional] [default to null]
**InitContainers** | **string** |  | [optional] [default to null]
**InsecureSkipVerify** | **bool** |  | [optional] [default to null]
**K8sConnectorID** | **string** |  | [optional] [default to null]
**Label** | **map[string]string** |  | [optional] [default to null]
**Mtls** | [***InfraV2MtlsConfiguration**](infra_v2.MTLSConfiguration.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NodeSelector** | **map[string]string** |  | [optional] [default to null]
**Proxy** | [***InfraV2ProxyConfiguration**](infra_v2.ProxyConfiguration.md) |  | [optional] [default to null]
**RunAsGroup** | **int32** |  | [optional] [default to null]
**RunAsUser** | **int32** |  | [optional] [default to null]
**ServiceAccount** | **string** |  | [optional] [default to null]
**SidecarContainers** | **string** |  | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**Tolerations** | [**[]V1Toleration**](v1.Toleration.md) |  | [optional] [default to null]
**VolumeMounts** | [**[]V1VolumeMount**](v1.VolumeMount.md) |  | [optional] [default to null]
**Volumes** | [**[]V1Volume**](v1.Volume.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
