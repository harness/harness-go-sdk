package mesh

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

// OutboundConfig configures the mesh outbound RoundTripper.
type OutboundConfig struct {
	// TargetServiceID is the JWT aud claim (AuthorizationServiceHeader value of the callee).
	TargetServiceID string
	// Principal optionally supplies a user/service-account principal for delegated minting.
	// When nil or SERVICE, M2M SignForSelf is used.
	Principal *Principal
	// ExtraClaims are merged into delegated tokens (platform claims cannot be overridden).
	ExtraClaims map[string]any
}

// NewRoundTripper wraps next and stamps X-Harness-Identity when outbound mesh is enabled.
// It does not add or strip HMAC Authorization — wrap order should be:
//
//	mesh.NewRoundTripper(holder, cfg, hmacTransport)
//
// so HMAC still runs underneath in dual-header mode. When OutboundOnly is true, the service
// should omit its HMAC transport itself.
func NewRoundTripper(h *Holder, cfg OutboundConfig, next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = http.DefaultTransport
	}
	return &meshRoundTripper{holder: h, cfg: cfg, next: next}
}

type meshRoundTripper struct {
	holder      *Holder
	cfg         OutboundConfig
	next        http.RoundTripper
	warnedOnce  atomic.Bool
}

func (t *meshRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.holder == nil || !t.holder.OutboundEnabled() || t.holder.Generator() == nil {
		return t.next.RoundTrip(req)
	}

	metrics := t.holder.Metrics()
	audience := t.cfg.TargetServiceID
	outboundOnly := t.holder.OutboundOnly()

	if audience == "" {
		metrics.RecordOutboundConfigError(ConfigErrorMissingTargetService)
		if !t.warnedOnce.Swap(true) {
			log.Printf("mesh: no targetServiceId declared on RoundTripper (host=%s); mesh header skipped. Add OutboundConfig.TargetServiceID",
				req.URL.Host)
		}
		if outboundOnly {
			return nil, ErrMissingTargetService
		}
		return t.next.RoundTrip(req)
	}

	token, err := t.holder.Generator().SignForPrincipal(t.cfg.Principal, audience, DefaultJWTTTL, t.cfg.ExtraClaims)
	if err != nil {
		if outboundOnly {
			return nil, fmt.Errorf("mesh outbound mint failed: %w", err)
		}
		log.Printf("mesh: outbound mint failed; continuing without mesh header: %v", err)
		return t.next.RoundTrip(req)
	}

	req2 := req.Clone(req.Context())
	req2.Header = req.Header.Clone()
	req2.Header.Set(IdentityHeader, token)
	return t.next.RoundTrip(req2)
}
