/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1LifecycleHandler struct {
	Exec *V1ExecAction `json:"exec,omitempty"`
	HttpGet *V1HttpGetAction `json:"httpGet,omitempty"`
	TcpSocket *V1TcpSocketAction `json:"tcpSocket,omitempty"`
}
