/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type EcsRecommendationDto struct {
	Id              string                       `json:"id,omitempty"`
	ClusterName     string                       `json:"clusterName,omitempty"`
	ServiceArn      string                       `json:"serviceArn,omitempty"`
	ServiceName     string                       `json:"serviceName,omitempty"`
	Current         map[string]string            `json:"current,omitempty"`
	PercentileBased map[string]map[string]string `json:"percentileBased,omitempty"`
	LastDayCost     *Cost                        `json:"lastDayCost,omitempty"`
	CpuHistogram    *HistogramExp                `json:"cpuHistogram,omitempty"`
	MemoryHistogram *HistogramExp                `json:"memoryHistogram,omitempty"`
}
