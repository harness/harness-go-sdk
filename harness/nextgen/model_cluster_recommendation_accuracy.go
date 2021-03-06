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

type ClusterRecommendationAccuracy struct {
	Cpu          float64 `json:"cpu,omitempty"`
	MasterPrice  float64 `json:"masterPrice,omitempty"`
	Memory       float64 `json:"memory,omitempty"`
	Nodes        int64   `json:"nodes,omitempty"`
	RegularNodes int64   `json:"regularNodes,omitempty"`
	RegularPrice float64 `json:"regularPrice,omitempty"`
	SpotNodes    int64   `json:"spotNodes,omitempty"`
	SpotPrice    float64 `json:"spotPrice,omitempty"`
	TotalPrice   float64 `json:"totalPrice,omitempty"`
	WorkerPrice  float64 `json:"workerPrice,omitempty"`
	Zone         string  `json:"zone,omitempty"`
}
