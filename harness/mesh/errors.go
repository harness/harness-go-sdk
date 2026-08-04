package mesh

import "errors"

var (
	// ErrAudienceRequired is returned when inbound is enabled without audience.
	ErrAudienceRequired = errors.New("meshIdentity.audience must be set when meshIdentity.inboundEnabled is true")

	// ErrInvalidMeshToken is returned when mesh validation fails and fallback is disabled.
	ErrInvalidMeshToken = errors.New("Invalid mesh token")

	// ErrMeshHeaderRequired is returned when reject-without-header is on and header is absent.
	ErrMeshHeaderRequired = errors.New("Mesh identity header required")

	// ErrMissingTargetService is returned in outboundOnly mode when TargetServiceID is empty.
	ErrMissingTargetService = errors.New("mesh outbound: no targetServiceId declared")

	// ErrMeshDisabled is returned when outbound mesh is required but holder has outbound off.
	ErrMeshDisabled = errors.New("mesh outbound is not enabled")
)
