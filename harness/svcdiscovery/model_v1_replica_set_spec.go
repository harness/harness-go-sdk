/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1ReplicaSetSpec struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`
	// Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller +optional
	Replicas int32 `json:"replicas,omitempty"`
	Selector *V1LabelSelector `json:"selector,omitempty"`
	Template *V1PodTemplateSpec `json:"template,omitempty"`
}
