package mesh

import (
	"context"
	"strings"
)

// PrincipalType matches Java PrincipalType names used in JWT type claims.
type PrincipalType string

const (
	PrincipalTypeService        PrincipalType = "SERVICE"
	PrincipalTypeUser           PrincipalType = "USER"
	PrincipalTypeServiceAccount PrincipalType = "SERVICE_ACCOUNT"
	PrincipalTypeAPIKey         PrincipalType = "API_KEY"
)

// Principal is the authenticated identity placed on request context after mesh/HMAC auth.
type Principal struct {
	Type      PrincipalType
	Name      string
	Email     string
	Username  string
	AccountID string
	Role      string
	// SpiffeID is set for mesh-authenticated SERVICE principals (PKIX-derived).
	SpiffeID string
	// ScopedTokenID is restored from delegated JWT claims when present.
	ScopedTokenID string
	// Extra holds additional JWT claims for service-specific use.
	Extra map[string]any
}

// IsService reports whether this is an M2M service principal.
func (p Principal) IsService() bool {
	return p.Type == PrincipalTypeService || p.Type == ""
}

type ctxKey int

const principalCtxKey ctxKey = 1

// ContextWithPrincipal returns a child context carrying p.
func ContextWithPrincipal(ctx context.Context, p Principal) context.Context {
	return context.WithValue(ctx, principalCtxKey, p)
}

// PrincipalFromContext extracts the principal, if any.
func PrincipalFromContext(ctx context.Context) (Principal, bool) {
	p, ok := ctx.Value(principalCtxKey).(Principal)
	return p, ok
}

// ResolveServiceIDFromSPIFFE returns the last path segment if it is a known service ID.
// spiffeID is a full URI like spiffe://trust.domain/cell/env/NextGenManager.
func ResolveServiceIDFromSPIFFE(spiffeID string) string {
	if spiffeID == "" {
		return ""
	}
	// Strip scheme
	rest := spiffeID
	if i := strings.Index(rest, "://"); i >= 0 {
		rest = rest[i+3:]
	}
	// Drop trust domain (first segment before /)
	if i := strings.Index(rest, "/"); i >= 0 {
		rest = rest[i+1:]
	} else {
		return ""
	}
	if rest == "" {
		return ""
	}
	last := rest
	if i := strings.LastIndex(rest, "/"); i >= 0 {
		last = rest[i+1:]
	}
	if last == "" || !IsKnownServiceID(last) {
		return ""
	}
	return last
}

// TrustDomainFromSPIFFE extracts the trust domain name from a SPIFFE ID URI.
func TrustDomainFromSPIFFE(spiffeID string) string {
	if !strings.HasPrefix(spiffeID, "spiffe://") {
		return ""
	}
	rest := strings.TrimPrefix(spiffeID, "spiffe://")
	if i := strings.Index(rest, "/"); i >= 0 {
		return rest[:i]
	}
	return rest
}
