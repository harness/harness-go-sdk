package mesh

import "time"

// HTTP header carrying the SPIRE-signed application JWT.
const IdentityHeader = "X-Harness-Identity"

// Transport label values for harness_mesh_auth_transport_total.
const (
	TransportSPIRE = "spire"
	TransportHMAC  = "hmac"
)

// JWT algorithm used to sign application JWTs.
const JWTAlgorithm = "ES256"

// ClockSkewLeeway is the default clock-skew tolerance when validating exp.
const ClockSkewLeeway = 60 * time.Second

// DefaultJWTTTL is the default outbound JWT lifetime.
const DefaultJWTTTL = 60 * time.Second

// DefaultSPIFFESocket is used when SPIFFE_ENDPOINT_SOCKET is unset.
const DefaultSPIFFESocket = "unix:///run/spire/agent-sockets/spire-agent.sock"

// DefaultAllowedAudience is the default additive audience (edge gateway).
const DefaultAllowedAudience = "gateway-internal-service"

// Prometheus metric names — must match Java MeshIdentityConstants exactly.
const (
	MetricMintTotal             = "harness_mesh_jwt_mint_total"
	MetricMintDuration          = "harness_mesh_jwt_mint_duration_seconds"
	MetricValidateTotal         = "harness_mesh_jwt_validate_total"
	MetricValidateDuration      = "harness_mesh_jwt_validate_duration_seconds"
	MetricHMACFallbackTotal     = "harness_mesh_hmac_fallback_total"
	MetricBundleRefreshTotal    = "harness_mesh_bundle_refresh_total"
	MetricAuthTransportTotal    = "harness_mesh_auth_transport_total"
	MetricOutboundConfigError   = "harness_mesh_outbound_config_error_total"
)

// Outcome / label constants.
const (
	OutcomeSuccess = "success"
	OutcomeFailure = "failure"
	LabelUnknown   = "unknown"
	LabelOther     = "other"

	ConfigErrorMissingTargetService = "missing_target_service"

	ReasonMalformedToken   = "malformed_token"
	ReasonChainUntrusted   = "chain_untrusted"
	ReasonSignatureInvalid = "signature_invalid"
	ReasonExpired          = "expired"
	ReasonClaimMismatch    = "claim_mismatch"
	ReasonDecodeFailed     = "decode_failed"
	ReasonOther            = "other"
)

// HeaderX5C is the JWT header claim for the embedded X.509 cert chain.
const HeaderX5C = "x5c"

// MaxTrustDomainLabels caps distinct trust_domain label values (Java MeshMetrics).
const MaxTrustDomainLabels = 16
