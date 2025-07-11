/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1VolumeMount struct {
	// Path within the container at which the volume should be mounted.  Must not contain ':'.
	MountPath string `json:"mountPath,omitempty"`
	// mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. +optional
	MountPropagation string `json:"mountPropagation,omitempty"`
	// This must match the Name of a Volume.
	Name string `json:"name,omitempty"`
	// Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root). +optional
	SubPath string `json:"subPath,omitempty"`
	// Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to \"\" (volume's root). SubPathExpr and SubPath are mutually exclusive. +optional
	SubPathExpr string `json:"subPathExpr,omitempty"`
}
