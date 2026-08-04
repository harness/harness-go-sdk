package mesh

import "sync"

// Holder is the process-wide mesh identity handle (Java MeshIdentityHolder).
type Holder struct {
	mu sync.RWMutex

	cfg       Config
	source    IdentitySource
	generator *TokenGenerator
	validator *TokenValidator
	metrics   *Metrics

	outboundEnabled bool
	outboundOnly    bool
	inboundEnabled  bool
	noop            bool
}

// Config returns a copy of the holder's config.
func (h *Holder) Config() Config {
	if h == nil {
		return Config{}
	}
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.cfg
}

// OutboundEnabled reports whether outbound mesh stamping is on.
func (h *Holder) OutboundEnabled() bool {
	return h != nil && h.outboundEnabled
}

// OutboundOnly reports whether HMAC headers should be omitted by the service (mesh-only stage).
func (h *Holder) OutboundOnly() bool {
	return h != nil && h.outboundOnly
}

// InboundEnabled reports whether inbound mesh validation is on.
func (h *Holder) InboundEnabled() bool {
	return h != nil && h.inboundEnabled
}

// IsNoop reports whether mesh was bootstrapped disabled (no SPIRE connection).
func (h *Holder) IsNoop() bool {
	return h == nil || h.noop
}

// Generator returns the token generator (nil if noop).
func (h *Holder) Generator() *TokenGenerator {
	if h == nil {
		return nil
	}
	return h.generator
}

// Validator returns the token validator (nil if noop / inbound off).
func (h *Holder) Validator() *TokenValidator {
	if h == nil {
		return nil
	}
	return h.validator
}

// Metrics returns the metrics sink.
func (h *Holder) Metrics() *Metrics {
	if h == nil {
		return nil
	}
	return h.metrics
}

// Source returns the identity source (nil if noop).
func (h *Holder) Source() IdentitySource {
	if h == nil {
		return nil
	}
	return h.source
}

// Close closes the underlying Workload API source.
func (h *Holder) Close() error {
	if h == nil || h.source == nil {
		return nil
	}
	return h.source.Close()
}

var (
	globalMu     sync.RWMutex
	globalHolder *Holder
)

// SetGlobal publishes the holder for packages that cannot take an explicit dependency
// (mirrors Java MeshIdentityHolder.set).
func SetGlobal(h *Holder) {
	globalMu.Lock()
	defer globalMu.Unlock()
	globalHolder = h
}

// Global returns the process-wide holder, or nil.
func Global() *Holder {
	globalMu.RLock()
	defer globalMu.RUnlock()
	return globalHolder
}
