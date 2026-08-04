package mesh

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// mintCustom builds an ES256 mesh JWT with full control over claims/header (test helper).
// Pass header=nil to auto-populate alg/typ/kid/x5c from the leaf chain. A non-nil header is used as-is.
func mintCustom(t *testing.T, leaf *x509.Certificate, key *ecdsa.PrivateKey, chain []*x509.Certificate, header map[string]any, claims map[string]any) string {
	t.Helper()
	if header == nil {
		x5c := make([]string, 0, len(chain))
		for _, c := range chain {
			x5c = append(x5c, base64.StdEncoding.EncodeToString(c.Raw))
		}
		header = map[string]any{
			"alg":     JWTAlgorithm,
			"typ":     "JWT",
			"kid":     computeKid(leaf),
			HeaderX5C: x5c,
		}
	}
	hj, err := json.Marshal(header)
	require.NoError(t, err)
	pj, err := json.Marshal(claims)
	require.NoError(t, err)
	b64 := base64.RawURLEncoding
	input := b64.EncodeToString(hj) + "." + b64.EncodeToString(pj)
	sum := sha256.Sum256([]byte(input))
	r, s, err := ecdsa.Sign(rand.Reader, key, sum[:])
	require.NoError(t, err)
	kb := (key.Curve.Params().BitSize + 7) / 8
	sig := append(padLeft(r.Bytes(), kb), padLeft(s.Bytes(), kb)...)
	return input + "." + b64.EncodeToString(sig)
}

func TestValidateRejectsExpiredBeyondLeeway(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, key := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	val := NewTokenValidator(src, testMetrics(t))

	now := time.Now()
	// Expired 90s ago — beyond ClockSkewLeeway (60s).
	token := mintCustom(t, svid.Chain[0], key, svid.Chain, nil, map[string]any{
		"iss": spiffeURI,
		"sub": spiffeURI,
		"aud": ServiceAccessControlService,
		"iat": now.Add(-120 * time.Second).Unix(),
		"exp": now.Add(-90 * time.Second).Unix(),
		"type": string(PrincipalTypeService),
		"name": ServiceNextGenManager,
	})

	_, err := val.Validate(token, ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonExpired, ReasonFrom(err))
}

func TestValidateAcceptsExpiredWithinLeeway(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, key := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	val := NewTokenValidator(src, testMetrics(t))

	now := time.Now()
	// Expired 30s ago — within ClockSkewLeeway (60s).
	token := mintCustom(t, svid.Chain[0], key, svid.Chain, nil, map[string]any{
		"iss": spiffeURI,
		"sub": spiffeURI,
		"aud": ServiceAccessControlService,
		"iat": now.Add(-90 * time.Second).Unix(),
		"exp": now.Add(-30 * time.Second).Unix(),
		"type": string(PrincipalTypeService),
	})

	_, err := val.Validate(token, ServiceAccessControlService, nil)
	require.NoError(t, err)
}

func TestValidateRejectsUntrustedChain(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, _, key := newTestCA(t, spiffeURI)
	// Different CA in the trust bundle — leaf will not chain.
	_, otherBundle, _ := newTestCA(t, "spiffe://other.td/OtherService")
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: otherBundle}}
	val := NewTokenValidator(src, testMetrics(t))

	now := time.Now()
	token := mintCustom(t, svid.Chain[0], key, svid.Chain, nil, map[string]any{
		"iss": spiffeURI,
		"sub": spiffeURI,
		"aud": ServiceAccessControlService,
		"iat": now.Unix(),
		"exp": now.Add(time.Minute).Unix(),
	})

	_, err := val.Validate(token, ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonChainUntrusted, ReasonFrom(err))
}

func TestValidateRejectsIssMismatch(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, key := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	val := NewTokenValidator(src, testMetrics(t))

	now := time.Now()
	token := mintCustom(t, svid.Chain[0], key, svid.Chain, nil, map[string]any{
		"iss": "spiffe://test.harness.io/qa/Impostor",
		"sub": spiffeURI,
		"aud": ServiceAccessControlService,
		"iat": now.Unix(),
		"exp": now.Add(time.Minute).Unix(),
	})

	_, err := val.Validate(token, ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonMalformedToken, ReasonFrom(err))
	require.Contains(t, err.Error(), "iss mismatch")
}

func TestValidateRejectsBadSignature(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	gen := NewTokenGenerator(src, testMetrics(t))
	val := NewTokenValidator(src, testMetrics(t))

	token, err := gen.SignForSelf(ServiceAccessControlService, DefaultJWTTTL)
	require.NoError(t, err)

	parts := splitJWT(t, token)
	// Flip last signature byte.
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	require.NoError(t, err)
	sig[len(sig)-1] ^= 0xff
	tampered := parts[0] + "." + parts[1] + "." + base64.RawURLEncoding.EncodeToString(sig)

	_, err = val.Validate(tampered, ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonSignatureInvalid, ReasonFrom(err))
}

func TestValidateRejectsMissingX5C(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, key := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	val := NewTokenValidator(src, testMetrics(t))

	now := time.Now()
	header := map[string]any{"alg": JWTAlgorithm, "typ": "JWT", "kid": "x"}
	token := mintCustom(t, svid.Chain[0], key, svid.Chain, header, map[string]any{
		"iss": spiffeURI,
		"sub": spiffeURI,
		"aud": ServiceAccessControlService,
		"iat": now.Unix(),
		"exp": now.Add(time.Minute).Unix(),
	})

	_, err := val.Validate(token, ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonMalformedToken, ReasonFrom(err))
}

func TestValidateRejectsMalformedJWT(t *testing.T) {
	val := NewTokenValidator(&StaticSource{}, testMetrics(t))
	_, err := val.Validate("not-a-jwt", ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonDecodeFailed, ReasonFrom(err))
}

func TestValidateRejectsMissingExp(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, key := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	val := NewTokenValidator(src, testMetrics(t))

	now := time.Now()
	token := mintCustom(t, svid.Chain[0], key, svid.Chain, nil, map[string]any{
		"iss": spiffeURI,
		"sub": spiffeURI,
		"aud": ServiceAccessControlService,
		"iat": now.Unix(),
	})

	_, err := val.Validate(token, ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonMalformedToken, ReasonFrom(err))
}

func TestReasonFrom(t *testing.T) {
	require.Equal(t, ReasonExpired, ReasonFrom(&ValidationError{Reason: ReasonExpired}))
	require.Equal(t, ReasonOther, ReasonFrom(nil))
	require.Equal(t, ReasonOther, ReasonFrom(errString("x")))
}

type errString string

func (e errString) Error() string { return string(e) }

func splitJWT(t *testing.T, token string) []string {
	t.Helper()
	parts := make([]string, 3)
	i, j := 0, 0
	for n := 0; n < 2; n++ {
		for j < len(token) && token[j] != '.' {
			j++
		}
		require.Less(t, j, len(token))
		parts[n] = token[i:j]
		j++
		i = j
	}
	parts[2] = token[i:]
	return parts
}

func TestNewTestCAHasSPIFFEURI(t *testing.T) {
	// sanity: leaf without SPIFFE should be rejected when crafted
	caKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	caTmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "ca"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	caDER, err := x509.CreateCertificate(rand.Reader, caTmpl, caTmpl, &caKey.PublicKey, caKey)
	require.NoError(t, err)
	caCert, err := x509.ParseCertificate(caDER)
	require.NoError(t, err)

	leafKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	u, err := url.Parse("https://example.com/not-spiffe")
	require.NoError(t, err)
	leafTmpl := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: "leaf"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		URIs:         []*url.URL{u},
	}
	leafDER, err := x509.CreateCertificate(rand.Reader, leafTmpl, caCert, &leafKey.PublicKey, caKey)
	require.NoError(t, err)
	leaf, err := x509.ParseCertificate(leafDER)
	require.NoError(t, err)

	bundle := &BundleRef{Authorities: []*x509.Certificate{caCert}, BundleVersion: "v1"}
	src := &StaticSource{
		SVID:    &Snapshot{Chain: []*x509.Certificate{leaf}, PrivateKey: leafKey},
		Bundles: map[string]*BundleRef{"test.harness.io": bundle},
	}
	val := NewTokenValidator(src, testMetrics(t))
	now := time.Now()
	token := mintCustom(t, leaf, leafKey, []*x509.Certificate{leaf}, nil, map[string]any{
		"iss": "spiffe://test.harness.io/x",
		"sub": "spiffe://test.harness.io/x",
		"aud": ServiceAccessControlService,
		"iat": now.Unix(),
		"exp": now.Add(time.Minute).Unix(),
	})
	_, err = val.Validate(token, ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonMalformedToken, ReasonFrom(err))
}
