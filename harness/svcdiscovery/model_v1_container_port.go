/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1ContainerPort struct {
	// Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort,omitempty"`
	// What host IP to bind the external port to. +optional
	HostIP string `json:"hostIP,omitempty"`
	// Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this. +optional
	HostPort int32 `json:"hostPort,omitempty"`
	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services. +optional
	Name string `json:"name,omitempty"`
	// Protocol for port. Must be UDP, TCP, or SCTP. Defaults to \"TCP\". +optional +default=\"TCP\"
	Protocol string `json:"protocol,omitempty"`
}
