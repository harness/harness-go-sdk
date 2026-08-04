package mesh

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

// TokenGenerator mints mesh JWTs using the workload SVID (Java MeshTokenGenerator).
type TokenGenerator struct {
	source  IdentitySource
	metrics *Metrics
}

// NewTokenGenerator creates a generator.
func NewTokenGenerator(source IdentitySource, metrics *Metrics) *TokenGenerator {
	return &TokenGenerator{source: source, metrics: metrics}
}

// SignForSelf mints an M2M JWT (sub = workload SPIFFE ID, type=SERVICE).
func (g *TokenGenerator) SignForSelf(audience string, ttl time.Duration) (string, error) {
	done := g.metrics.StartMint()
	defer done()
	token, err := g.signForSelf(audience, ttl)
	if err != nil {
		g.metrics.RecordMint(OutcomeFailure)
		return "", err
	}
	g.metrics.RecordMint(OutcomeSuccess)
	return token, nil
}

func (g *TokenGenerator) signForSelf(audience string, ttl time.Duration) (string, error) {
	svid, err := g.source.Current()
	if err != nil {
		return "", err
	}
	self, err := deriveSelf(svid)
	if err != nil {
		return "", err
	}
	claims := map[string]any{"type": string(PrincipalTypeService)}
	if self.serviceID != "" {
		claims["name"] = self.serviceID
	}
	return g.sign(svid, self.spiffeID, self.spiffeID, audience, ttl, nil, claims)
}

// SignForPrincipal mints M2M or user-delegated JWT based on principal type.
func (g *TokenGenerator) SignForPrincipal(p *Principal, audience string, ttl time.Duration, extra map[string]any) (string, error) {
	if p == nil || p.IsService() {
		return g.SignForSelf(audience, ttl)
	}
	return g.SignForUser(*p, audience, ttl, extra)
}

// SignForUser mints a user-delegated JWT with act claim.
func (g *TokenGenerator) SignForUser(user Principal, audience string, ttl time.Duration, extra map[string]any) (string, error) {
	done := g.metrics.StartMint()
	defer done()
	token, err := g.signForUser(user, audience, ttl, extra)
	if err != nil {
		g.metrics.RecordMint(OutcomeFailure)
		return "", err
	}
	g.metrics.RecordMint(OutcomeSuccess)
	return token, nil
}

func (g *TokenGenerator) signForUser(user Principal, audience string, ttl time.Duration, extra map[string]any) (string, error) {
	svid, err := g.source.Current()
	if err != nil {
		return "", err
	}
	self, err := deriveSelf(svid)
	if err != nil {
		return "", err
	}
	actor := map[string]any{"sub": self.spiffeID}
	if self.serviceID != "" {
		actor["type"] = string(PrincipalTypeService)
		actor["name"] = self.serviceID
	}
	claims := map[string]any{}
	for k, v := range extra {
		claims[k] = v
	}
	populatePrincipalClaims(claims, user)
	return g.sign(svid, self.spiffeID, user.Name, audience, ttl, actor, claims)
}

func populatePrincipalClaims(claims map[string]any, p Principal) {
	claims["type"] = string(p.Type)
	claims["name"] = p.Name
	putOrRemove(claims, "email", p.Email)
	putOrRemove(claims, "username", p.Username)
	putOrRemove(claims, "accountId", p.AccountID)
	putOrRemove(claims, "role", p.Role)
}

func putOrRemove(m map[string]any, key, value string) {
	if value != "" {
		m[key] = value
	} else {
		delete(m, key)
	}
}

type selfIdentity struct {
	spiffeID  string
	serviceID string
}

func deriveSelf(svid *Snapshot) (selfIdentity, error) {
	leaf := svid.Leaf()
	if leaf == nil {
		return selfIdentity{}, fmt.Errorf("No SVID leaf cert available from source")
	}
	spiffeURI := extractSpiffeURI(leaf)
	if spiffeURI == "" {
		if svid.SpiffeID != "" {
			spiffeURI = svid.SpiffeID
		} else {
			return selfIdentity{}, fmt.Errorf("SVID leaf cert has no spiffe:// URI SAN — cannot mint mesh JWT")
		}
	}
	serviceID := ResolveServiceIDFromSPIFFE(spiffeURI)
	return selfIdentity{spiffeID: spiffeURI, serviceID: serviceID}, nil
}

func extractSpiffeURI(cert *x509.Certificate) string {
	for _, u := range cert.URIs {
		if u != nil && u.Scheme == "spiffe" {
			return u.String()
		}
	}
	return ""
}

func (g *TokenGenerator) sign(svid *Snapshot, issuer, subject, audience string, ttl time.Duration, act map[string]any, extraClaims map[string]any) (string, error) {
	if ttl <= 0 {
		ttl = DefaultJWTTTL
	}
	now := time.Now()

	x5c := make([]string, 0, len(svid.Chain))
	for _, c := range svid.Chain {
		x5c = append(x5c, base64.StdEncoding.EncodeToString(c.Raw))
	}
	header := map[string]any{
		"typ": "JWT",
		"alg": JWTAlgorithm,
		"kid": svid.Kid,
		HeaderX5C: x5c,
	}

	payload := map[string]any{
		"iss": issuer,
		"sub": subject,
		"aud": audience,
		"iat": now.Unix(),
		"exp": now.Add(ttl).Unix(),
	}
	if act != nil {
		payload["act"] = act
	}
	for k, v := range extraClaims {
		if _, exists := payload[k]; !exists {
			payload[k] = v
		}
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("Failed to serialize JWT header: %w", err)
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Failed to serialize JWT payload: %w", err)
	}
	b64 := base64.RawURLEncoding
	signingInput := b64.EncodeToString(headerJSON) + "." + b64.EncodeToString(payloadJSON)

	ecKey, ok := svid.PrivateKey.(*ecdsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("SVID private key is not ECDSA")
	}
	sig, err := signES256(ecKey, []byte(signingInput))
	if err != nil {
		return "", fmt.Errorf("Failed to sign mesh JWT: %w", err)
	}
	return signingInput + "." + b64.EncodeToString(sig), nil
}

func signES256(key *ecdsa.PrivateKey, signingInput []byte) ([]byte, error) {
	sum := sha256.Sum256(signingInput)
	r, s, err := ecdsa.Sign(rand.Reader, key, sum[:])
	if err != nil {
		return nil, err
	}
	curveBits := key.Curve.Params().BitSize
	keyBytes := (curveBits + 7) / 8
	return append(padLeft(r.Bytes(), keyBytes), padLeft(s.Bytes(), keyBytes)...), nil
}

func padLeft(b []byte, size int) []byte {
	if len(b) >= size {
		return b[:size]
	}
	out := make([]byte, size)
	copy(out[size-len(b):], b)
	return out
}

func verifyES256(pub *ecdsa.PublicKey, signingInput, sig []byte) bool {
	curveBits := pub.Curve.Params().BitSize
	keyBytes := (curveBits + 7) / 8
	if len(sig) != 2*keyBytes {
		return false
	}
	r := new(big.Int).SetBytes(sig[:keyBytes])
	s := new(big.Int).SetBytes(sig[keyBytes:])
	sum := sha256.Sum256(signingInput)
	return ecdsa.Verify(pub, sum[:], r, s)
}
