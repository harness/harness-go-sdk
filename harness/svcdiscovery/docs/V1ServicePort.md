# V1ServicePort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProtocol** | **string** | The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. +optional | [optional] [default to null]
**Name** | **string** | The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the &#x27;name&#x27; field in the EndpointPort. Optional if only one ServicePort is defined on this service. +optional | [optional] [default to null]
**NodePort** | **int32** | The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport +optional | [optional] [default to null]
**Port** | **int32** | The port that will be exposed by this service. | [optional] [default to null]
**Protocol** | **string** | The IP protocol for this port. Supports \&quot;TCP\&quot;, \&quot;UDP\&quot;, and \&quot;SCTP\&quot;. Default is TCP. +default&#x3D;\&quot;TCP\&quot; +optional | [optional] [default to null]
**TargetPort** | [***IntstrIntOrString**](intstr.IntOrString.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

