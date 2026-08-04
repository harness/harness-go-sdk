package mesh

import (
	"context"
	"fmt"
	"log"
)

// BootstrapOptions customize Bootstrap (tests inject source/metrics).
type BootstrapOptions struct {
	Source  IdentitySource // if nil and mesh active, connects to Workload API
	Metrics *Metrics       // if nil, uses DefaultMetrics when mesh active
}

// Bootstrap initializes mesh identity. When both inbound and outbound are false, returns a
// no-op Holder with no SPIRE connection (Java MeshIdentityBootstrap behavior).
func Bootstrap(ctx context.Context, cfg Config, opts *BootstrapOptions) (*Holder, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	if !cfg.MeshActive() {
		h := &Holder{cfg: cfg, noop: true, metrics: nil}
		SetGlobal(h)
		return h, nil
	}

	var src IdentitySource
	var metrics *Metrics
	if opts != nil {
		src = opts.Source
		metrics = opts.Metrics
	}
	if metrics == nil {
		metrics = DefaultMetrics()
	}
	if src == nil {
		ws, err := NewWorkloadSource(ctx, cfg.SPIFFEEndpointSocket)
		if err != nil {
			return nil, err
		}
		src = ws
	}

	gen := NewTokenGenerator(src, metrics)
	var val *TokenValidator
	if cfg.InboundEnabled {
		val = NewTokenValidator(src, metrics)
	}

	h := &Holder{
		cfg:             cfg,
		source:          src,
		generator:       gen,
		validator:       val,
		metrics:         metrics,
		outboundEnabled: cfg.OutboundEnabled,
		outboundOnly:    cfg.OutboundOnly,
		inboundEnabled:  cfg.InboundEnabled,
	}
	SetGlobal(h)
	log.Printf("mesh: bootstrapped inbound=%v outbound=%v outboundOnly=%v fallback=%v audience=%q",
		cfg.InboundEnabled, cfg.OutboundEnabled, cfg.OutboundOnly, cfg.FallbackEnabled, cfg.Audience)
	return h, nil
}

// MustBootstrap is Bootstrap that panics on error.
func MustBootstrap(ctx context.Context, cfg Config, opts *BootstrapOptions) *Holder {
	h, err := Bootstrap(ctx, cfg, opts)
	if err != nil {
		panic(fmt.Sprintf("mesh.Bootstrap: %v", err))
	}
	return h
}
