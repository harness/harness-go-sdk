package mesh

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net/url"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func testMetrics(t *testing.T) *Metrics {
	t.Helper()
	reg := prometheus.NewRegistry()
	return NewMetrics(reg)
}

// newTestCA creates a CA and a leaf SVID for spiffe://td/path with matching trust bundle.
func newTestCA(t *testing.T, spiffeURI string) (*Snapshot, *BundleRef, *ecdsa.PrivateKey) {
	t.Helper()
	caKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	caTmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "test-ca"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
	}
	caDER, err := x509.CreateCertificate(rand.Reader, caTmpl, caTmpl, &caKey.PublicKey, caKey)
	require.NoError(t, err)
	caCert, err := x509.ParseCertificate(caDER)
	require.NoError(t, err)

	leafKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	u, err := url.Parse(spiffeURI)
	require.NoError(t, err)

	leafTmpl := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: "test-leaf"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		URIs:         []*url.URL{u},
	}
	leafDER, err := x509.CreateCertificate(rand.Reader, leafTmpl, caCert, &leafKey.PublicKey, caKey)
	require.NoError(t, err)
	leafCert, err := x509.ParseCertificate(leafDER)
	require.NoError(t, err)

	svid := &Snapshot{
		Chain:      []*x509.Certificate{leafCert},
		PrivateKey: leafKey,
		Kid:        computeKid(leafCert),
		SpiffeID:   spiffeURI,
	}
	bundle := &BundleRef{
		Authorities:   []*x509.Certificate{caCert},
		BundleVersion: "v1",
	}
	return svid, bundle, leafKey
}

func TestConfigFromEnvDefaults(t *testing.T) {
	t.Setenv("MESH_IDENTITY_INBOUND_ENABLED", "")
	t.Setenv("MESH_IDENTITY_OUTBOUND_ENABLED", "")
	t.Setenv("MESH_IDENTITY_FALLBACK_ENABLED", "")
	t.Setenv("MESH_IDENTITY_AUDIENCE", "")
	t.Setenv("MESH_IDENTITY_ALLOWED_AUDIENCES", "")
	t.Setenv("SPIFFE_ENDPOINT_SOCKET", "")

	cfg := ConfigFromEnv()
	require.False(t, cfg.InboundEnabled)
	require.False(t, cfg.OutboundEnabled)
	require.True(t, cfg.FallbackEnabled)
	require.Equal(t, DefaultSPIFFESocket, cfg.SPIFFEEndpointSocket)
	require.Equal(t, []string{DefaultAllowedAudience}, cfg.AllowedAudiences)
}

func TestConfigFromEnvOverrides(t *testing.T) {
	t.Setenv("MESH_IDENTITY_INBOUND_ENABLED", "true")
	t.Setenv("MESH_IDENTITY_OUTBOUND_ENABLED", "true")
	t.Setenv("MESH_IDENTITY_OUTBOUND_ONLY", "true")
	t.Setenv("MESH_IDENTITY_FALLBACK_ENABLED", "false")
	t.Setenv("MESH_IDENTITY_REJECT_AUTH_WITHOUT_MESH_HEADER", "true")
	t.Setenv("MESH_IDENTITY_AUDIENCE", ServicePipelineService)
	t.Setenv("MESH_IDENTITY_ALLOWED_AUDIENCES", "gateway-internal-service,other-svc")
	t.Setenv("SPIFFE_ENDPOINT_SOCKET", "unix:///tmp/spire.sock")

	cfg := ConfigFromEnv()
	require.True(t, cfg.InboundEnabled)
	require.True(t, cfg.OutboundEnabled)
	require.True(t, cfg.OutboundOnly)
	require.False(t, cfg.FallbackEnabled)
	require.True(t, cfg.RejectWithoutMeshHeader)
	require.Equal(t, ServicePipelineService, cfg.Audience)
	require.Equal(t, []string{DefaultAllowedAudience, "other-svc"}, cfg.AllowedAudiences)
	require.Equal(t, "unix:///tmp/spire.sock", cfg.SPIFFEEndpointSocket)
	require.True(t, cfg.MeshActive())
}

func TestConfigValidateAudienceRequired(t *testing.T) {
	cfg := Config{InboundEnabled: true}
	require.ErrorIs(t, cfg.Validate(), ErrAudienceRequired)
	cfg.Audience = ServiceNextGenManager
	require.NoError(t, cfg.Validate())
}

func TestResolveServiceIDFromSPIFFE(t *testing.T) {
	require.Equal(t, ServiceNextGenManager, ResolveServiceIDFromSPIFFE("spiffe://harness.io/qa/NextGenManager"))
	require.Equal(t, ServiceAccessControlService, ResolveServiceIDFromSPIFFE("spiffe://td/accessControlService"))
	require.Equal(t, "", ResolveServiceIDFromSPIFFE("spiffe://td/unknown-svc"))
	require.Equal(t, "harness.io", TrustDomainFromSPIFFE("spiffe://harness.io/qa/NextGenManager"))
}

func TestMintValidateRoundTrip(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{
		SVID:    svid,
		Bundles: map[string]*BundleRef{td: bundle},
	}
	metrics := testMetrics(t)
	gen := NewTokenGenerator(src, metrics)
	val := NewTokenValidator(src, metrics)

	token, err := gen.SignForSelf(ServiceAccessControlService, DefaultJWTTTL)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	v, err := val.Validate(token, ServiceAccessControlService, nil)
	require.NoError(t, err)
	require.Equal(t, spiffeURI, v.SpiffeID)
	require.Equal(t, string(PrincipalTypeService), v.Claims["type"])
	require.Equal(t, ServiceNextGenManager, v.Claims["name"])
}

func TestValidateRejectsWrongAudience(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	metrics := testMetrics(t)
	gen := NewTokenGenerator(src, metrics)
	val := NewTokenValidator(src, metrics)

	token, err := gen.SignForSelf(ServicePipelineService, DefaultJWTTTL)
	require.NoError(t, err)

	_, err = val.Validate(token, ServiceAccessControlService, nil)
	require.Error(t, err)
	require.Equal(t, ReasonMalformedToken, ReasonFrom(err))
}

func TestValidateAcceptsAllowedAudience(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	metrics := testMetrics(t)
	gen := NewTokenGenerator(src, metrics)
	val := NewTokenValidator(src, metrics)

	token, err := gen.SignForSelf(DefaultAllowedAudience, DefaultJWTTTL)
	require.NoError(t, err)

	_, err = val.Validate(token, ServiceAccessControlService, []string{DefaultAllowedAudience})
	require.NoError(t, err)
}

func TestSignForUserActClaim(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	metrics := testMetrics(t)
	gen := NewTokenGenerator(src, metrics)
	val := NewTokenValidator(src, metrics)

	user := Principal{Type: PrincipalTypeUser, Name: "user@harness.io", Email: "user@harness.io", AccountID: "acc"}
	token, err := gen.SignForUser(user, ServiceAccessControlService, DefaultJWTTTL, nil)
	require.NoError(t, err)

	v, err := val.Validate(token, ServiceAccessControlService, nil)
	require.NoError(t, err)
	act, ok := v.Claims["act"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, spiffeURI, act["sub"])
}

func TestSignForPrincipalRoutes(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	gen := NewTokenGenerator(src, testMetrics(t))
	val := NewTokenValidator(src, testMetrics(t))

	token, err := gen.SignForPrincipal(nil, ServiceAccessControlService, DefaultJWTTTL, nil)
	require.NoError(t, err)
	v, err := val.Validate(token, ServiceAccessControlService, nil)
	require.NoError(t, err)
	require.Equal(t, string(PrincipalTypeService), v.Claims["type"])

	user := &Principal{Type: PrincipalTypeUser, Name: "u@harness.io", Email: "u@harness.io", Role: "admin"}
	token, err = gen.SignForPrincipal(user, ServiceAccessControlService, DefaultJWTTTL, nil)
	require.NoError(t, err)
	v, err = val.Validate(token, ServiceAccessControlService, nil)
	require.NoError(t, err)
	require.Equal(t, string(PrincipalTypeUser), v.Claims["type"])
	require.Equal(t, "admin", v.Claims["role"])
	require.False(t, user.IsService())
	svc := &Principal{Type: PrincipalTypeService, Name: ServiceNextGenManager}
	require.True(t, svc.IsService())
}