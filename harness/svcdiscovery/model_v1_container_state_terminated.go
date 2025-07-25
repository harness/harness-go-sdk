/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1ContainerStateTerminated struct {
	// Container's ID in the format '<type>://<container_id>' +optional
	ContainerID string `json:"containerID,omitempty"`
	// Exit status from the last termination of the container
	ExitCode int32 `json:"exitCode,omitempty"`
	// Time at which the container last terminated +optional
	FinishedAt string `json:"finishedAt,omitempty"`
	// Message regarding the last termination of the container +optional
	Message string `json:"message,omitempty"`
	// (brief) reason from the last termination of the container +optional
	Reason string `json:"reason,omitempty"`
	// Signal from the last termination of the container +optional
	Signal int32 `json:"signal,omitempty"`
	// Time at which previous execution of the container started +optional
	StartedAt string `json:"startedAt,omitempty"`
}
