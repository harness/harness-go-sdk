package mesh

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Validated is the result of a successful mesh JWT validation.
type Validated struct {
	Claims     map[string]any
	SpiffeID   string
	TrustDomain string
	Header     map[string]any
	Raw        string
}

// TokenValidator validates inbound mesh JWTs (Java MeshTokenValidator).
type TokenValidator struct {
	source  IdentitySource
	metrics *Metrics

	mu          sync.Mutex
	anchorCache map[string]anchorCacheEntry
}

type anchorCacheEntry struct {
	bundleVersion string
	roots         *x509.CertPool
}

// NewTokenValidator creates a validator.
func NewTokenValidator(source IdentitySource, metrics *Metrics) *TokenValidator {
	return &TokenValidator{
		source:      source,
		metrics:     metrics,
		anchorCache: make(map[string]anchorCacheEntry),
	}
}

// Validate validates token against expectedAudience ∪ allowedAudiences.
func (v *TokenValidator) Validate(token, expectedAudience string, allowedAudiences []string) (*Validated, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, &ValidationError{Reason: ReasonDecodeFailed, Err: fmt.Errorf("malformed JWT")}
	}
	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, &ValidationError{Reason: ReasonDecodeFailed, Err: fmt.Errorf("failed to decode JWT header: %w", err)}
	}
	var header map[string]any
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, &ValidationError{Reason: ReasonDecodeFailed, Err: fmt.Errorf("failed to parse JWT header: %w", err)}
	}

	x5cRaw, ok := header[HeaderX5C]
	if !ok {
		return nil, &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("mesh JWT missing x5c header")}
	}
	x5cList, err := asStringSlice(x5cRaw)
	if err != nil || len(x5cList) == 0 {
		return nil, &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("mesh JWT missing x5c header")}
	}

	chain := make([]*x509.Certificate, 0, len(x5cList))
	for _, enc := range x5cList {
		cert, err := decodeX5C(enc)
		if err != nil {
			return nil, &ValidationError{Reason: ReasonMalformedToken, Err: err}
		}
		chain = append(chain, cert)
	}
	leaf := chain[0]
	spiffeID := extractSpiffeURI(leaf)
	if spiffeID == "" {
		return nil, &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("leaf cert has no SPIFFE ID in SAN")}
	}
	td := TrustDomainFromSPIFFE(spiffeID)

	if err := v.verifyChain(chain, spiffeID, td); err != nil {
		return nil, err
	}

	pub, ok := leaf.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("unsupported key type: %T", leaf.PublicKey)}
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, &ValidationError{Reason: ReasonDecodeFailed, Err: fmt.Errorf("failed to decode JWT payload: %w", err)}
	}
	var claims map[string]any
	if err := json.Unmarshal(payloadJSON, &claims); err != nil {
		return nil, &ValidationError{Reason: ReasonDecodeFailed, Err: fmt.Errorf("failed to parse JWT payload: %w", err)}
	}

	tokenAuds := claimAudiences(claims)
	matched := resolveAudience(tokenAuds, expectedAudience, allowedAudiences)
	if matched == "" {
		return nil, &ValidationError{
			Reason: ReasonMalformedToken,
			Err:    fmt.Errorf("token aud %v not in accepted audiences [%s, %v]", tokenAuds, expectedAudience, allowedAudiences),
		}
	}

	signingInput := parts[0] + "." + parts[1]
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, &ValidationError{Reason: ReasonDecodeFailed, Err: fmt.Errorf("failed to decode JWT signature: %w", err)}
	}
	if !verifyES256(pub, []byte(signingInput), sig) {
		return nil, &ValidationError{Reason: ReasonSignatureInvalid, Err: fmt.Errorf("signature verification failed")}
	}

	now := time.Now()
	exp, ok := claimUnix(claims, "exp")
	if !ok {
		return nil, &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("mesh JWT missing exp claim")}
	}
	if _, ok := claimUnix(claims, "iat"); !ok {
		return nil, &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("mesh JWT missing iat claim")}
	}
	if now.After(exp.Add(ClockSkewLeeway)) {
		return nil, &ValidationError{Reason: ReasonExpired, Err: fmt.Errorf("token is expired")}
	}
	// Also reject not-yet-valid beyond leeway if nbf present — not required by Java.

	iss, _ := claims["iss"].(string)
	if iss != spiffeID {
		return nil, &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("iss mismatch: leaf=%s jwt.iss=%v", spiffeID, claims["iss"])}
	}

	// Enforce aud claim matches matched (already peeked; bind after signature).
	if !audienceContains(tokenAuds, matched) {
		return nil, &ValidationError{Reason: ReasonClaimMismatch, Err: fmt.Errorf("audience mismatch")}
	}

	return &Validated{
		Claims:      claims,
		SpiffeID:    spiffeID,
		TrustDomain: td,
		Header:      header,
		Raw:         token,
	}, nil
}

func (v *TokenValidator) verifyChain(chain []*x509.Certificate, spiffeID, td string) error {
	roots, err := v.anchorsFor(td)
	if err != nil {
		return err
	}
	leaf := chain[0]
	intermediates := x509.NewCertPool()
	for _, c := range chain[1:] {
		intermediates.AddCert(c)
	}
	opts := x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermediates,
		CurrentTime:   time.Now(),
		KeyUsages:    []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	}
	if _, err := leaf.Verify(opts); err != nil {
		return &ValidationError{Reason: ReasonChainUntrusted, Err: fmt.Errorf("PKIX chain validation failed for %s: %w", spiffeID, err)}
	}
	return nil
}

func (v *TokenValidator) anchorsFor(td string) (*x509.CertPool, error) {
	ref, err := v.source.BundleForTrustDomain(td)
	if err != nil {
		v.metrics.RecordBundleRefresh(OutcomeFailure, td)
		return nil, &ValidationError{Reason: ReasonChainUntrusted, Err: fmt.Errorf("no X.509 trust bundle for trust domain %s: %w", td, err)}
	}
	if ref == nil || len(ref.Authorities) == 0 {
		v.metrics.RecordBundleRefresh(OutcomeFailure, td)
		return nil, &ValidationError{Reason: ReasonChainUntrusted, Err: fmt.Errorf("empty X.509 trust bundle for trust domain %s", td)}
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	if cached, ok := v.anchorCache[td]; ok && cached.bundleVersion == ref.BundleVersion {
		return cached.roots, nil
	}
	pool := x509.NewCertPool()
	for _, a := range ref.Authorities {
		pool.AddCert(a)
	}
	v.anchorCache[td] = anchorCacheEntry{bundleVersion: ref.BundleVersion, roots: pool}
	v.metrics.RecordBundleRefresh(OutcomeSuccess, td)
	return pool, nil
}

// ValidationError carries a stable fallback reason label.
type ValidationError struct {
	Reason string
	Err    error
}

func (e *ValidationError) Error() string {
	if e.Err == nil {
		return e.Reason
	}
	return e.Err.Error()
}

func (e *ValidationError) Unwrap() error { return e.Err }

// ReasonFrom maps an error to hmac_fallback reason labels (Java reasonFrom).
func ReasonFrom(err error) string {
	var ve *ValidationError
	if errors.As(err, &ve) && ve.Reason != "" {
		return ve.Reason
	}
	return ReasonOther
}

func decodeX5C(enc string) (*x509.Certificate, error) {
	der, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		// some stacks use raw std without padding issues — try RawStd
		der, err = base64.RawStdEncoding.DecodeString(enc)
		if err != nil {
			return nil, fmt.Errorf("failed to decode x5c cert entry: %w", err)
		}
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, fmt.Errorf("failed to decode x5c cert entry: %w", err)
	}
	return cert, nil
}

func asStringSlice(v any) ([]string, error) {
	switch t := v.(type) {
	case []string:
		return t, nil
	case []any:
		out := make([]string, 0, len(t))
		for _, el := range t {
			s, ok := el.(string)
			if !ok {
				return nil, fmt.Errorf("x5c entry not string")
			}
			out = append(out, s)
		}
		return out, nil
	default:
		return nil, fmt.Errorf("x5c not a list")
	}
}

func claimAudiences(claims map[string]any) []string {
	raw, ok := claims["aud"]
	if !ok {
		return nil
	}
	switch t := raw.(type) {
	case string:
		return []string{t}
	case []any:
		out := make([]string, 0, len(t))
		for _, el := range t {
			if s, ok := el.(string); ok {
				out = append(out, s)
			}
		}
		return out
	case []string:
		return t
	default:
		return nil
	}
}

func resolveAudience(tokenAuds []string, primary string, extra []string) string {
	for _, aud := range tokenAuds {
		if aud == primary {
			return aud
		}
	}
	for _, allowed := range extra {
		if allowed == "" {
			continue
		}
		for _, aud := range tokenAuds {
			if aud == allowed {
				return aud
			}
		}
	}
	return ""
}

func audienceContains(auds []string, want string) bool {
	for _, a := range auds {
		if a == want {
			return true
		}
	}
	return false
}

func claimUnix(claims map[string]any, key string) (time.Time, bool) {
	raw, ok := claims[key]
	if !ok {
		return time.Time{}, false
	}
	switch t := raw.(type) {
	case float64:
		return time.Unix(int64(t), 0), true
	case json.Number:
		i, err := t.Int64()
		if err != nil {
			return time.Time{}, false
		}
		return time.Unix(i, 0), true
	case int64:
		return time.Unix(t, 0), true
	case int:
		return time.Unix(int64(t), 0), true
	default:
		return time.Time{}, false
	}
}
