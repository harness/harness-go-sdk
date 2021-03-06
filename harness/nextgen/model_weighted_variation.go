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

// A variation and the weighting it should receive as part of a percentage rollout
type WeightedVariation struct {
	// The variation identifier
	Variation string `json:"variation"`
	// The weight to be given to the variation in percent
	Weight int32 `json:"weight"`
}
