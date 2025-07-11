/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

// TLSConfig contains the tls configuration for the prometheus probe
type AllOfv1PrometheusInputsTlsConfig struct {
	// InsecureSkipVerify flag to skip certificate checks
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
	// SecretName contains the secret name for the TLS configuration
	SecretName string `json:"secretName,omitempty"`
	// ServerName contains the server name for the TLS configuration
	ServerName string `json:"serverName,omitempty"`
}
