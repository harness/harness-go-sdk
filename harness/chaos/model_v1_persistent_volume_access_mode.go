/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1PersistentVolumeAccessMode string

// List of v1.PersistentVolumeAccessMode
const (
	READ_WRITE_ONCE_V1PersistentVolumeAccessMode V1PersistentVolumeAccessMode = "ReadWriteOnce"
	READ_ONLY_MANY_V1PersistentVolumeAccessMode V1PersistentVolumeAccessMode = "ReadOnlyMany"
	READ_WRITE_MANY_V1PersistentVolumeAccessMode V1PersistentVolumeAccessMode = "ReadWriteMany"
)