/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
type AllOfv1ScaleIoVolumeSourceSecretRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid? +optional
	Name string `json:"name,omitempty"`
}