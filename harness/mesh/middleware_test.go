package mesh

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testHolder(t *testing.T, cfg Config, src IdentitySource) *Holder {
	t.Helper()
	metrics := testMetrics(t)
	h, err := Bootstrap(context.Background(), cfg, &BootstrapOptions{Source: src, Metrics: metrics})
	require.NoError(t, err)
	t.Cleanup(func() { _ = h.Close() })
	return h
}

func TestMiddlewareMeshSuccess(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{
		InboundEnabled:  true,
		FallbackEnabled: true,
		Audience:        ServiceAccessControlService,
	}
	h := testHolder(t, cfg, src)
	token, err := h.Generator().SignForSelf(ServiceAccessControlService, DefaultJWTTTL)
	require.NoError(t, err)

	var got Principal
	handler := Middleware(h, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, ok := PrincipalFromContext(r.Context())
		require.True(t, ok)
		got = p
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://example/api", nil)
	req.Header.Set(IdentityHeader, token)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, ServiceNextGenManager, got.Name)
	require.Equal(t, PrincipalTypeService, got.Type)
}

func TestMiddlewareDelegatedUser(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{InboundEnabled: true, Audience: ServiceAccessControlService}
	h := testHolder(t, cfg, src)
	user := Principal{Type: PrincipalTypeUser, Name: "user@harness.io", Email: "user@harness.io", AccountID: "acc1"}
	token, err := h.Generator().SignForUser(user, ServiceAccessControlService, DefaultJWTTTL, nil)
	require.NoError(t, err)

	var got Principal
	handler := Middleware(h, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, ok := PrincipalFromContext(r.Context())
		require.True(t, ok)
		got = p
		w.WriteHeader(http.StatusOK)
	}))
	req := httptest.NewRequest(http.MethodGet, "http://example/api", nil)
	req.Header.Set(IdentityHeader, token)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, PrincipalTypeUser, got.Type)
	require.Equal(t, "user@harness.io", got.Name)
	require.Equal(t, "acc1", got.AccountID)
}

func TestMiddlewareRejectsDelegatedActMismatch(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, key := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	cfg := Config{InboundEnabled: true, FallbackEnabled: false, Audience: ServiceAccessControlService}
	h := testHolder(t, cfg, src)

	now := time.Now()
	token := mintCustom(t, svid.Chain[0], key, svid.Chain, nil, map[string]any{
		"iss": spiffeURI,
		"sub": "user@harness.io",
		"aud": ServiceAccessControlService,
		"iat": now.Unix(),
		"exp": now.Add(time.Minute).Unix(),
		"type": string(PrincipalTypeUser),
		"name": "user@harness.io",
		"act":  map[string]any{"sub": "spiffe://test.harness.io/qa/Impostor"},
	})

	handler := Middleware(h, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("handler should not run")
	}))
	req := httptest.NewRequest(http.MethodGet, "http://example/api", nil)
	req.Header.Set(IdentityHeader, token)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusUnauthorized, rr.Code)
	require.Contains(t, rr.Body.String(), ErrInvalidMeshToken.Error())
}

func TestMiddlewareRejectsDelegatedMissingAct(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, key := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}
	cfg := Config{InboundEnabled: true, FallbackEnabled: false, Audience: ServiceAccessControlService}
	h := testHolder(t, cfg, src)

	now := time.Now()
	token := mintCustom(t, svid.Chain[0], key, svid.Chain, nil, map[string]any{
		"iss": spiffeURI,
		"sub": "user@harness.io",
		"aud": ServiceAccessControlService,
		"iat": now.Unix(),
		"exp": now.Add(time.Minute).Unix(),
		"type": string(PrincipalTypeUser),
		"name": "user@harness.io",
	})

	handler := Middleware(h, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("handler should not run")
	}))
	req := httptest.NewRequest(http.MethodGet, "http://example/api", nil)
	req.Header.Set(IdentityHeader, token)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusUnauthorized, rr.Code)
	require.Contains(t, rr.Body.String(), ErrInvalidMeshToken.Error())
}

func TestMiddlewareNilFallbackRejectsInvalidMeshToken(t *testing.T) {
	// FallbackEnabled (default) + nil FallbackAuth must not fail-open on a garbage mesh header.
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{
		InboundEnabled:  true,
		FallbackEnabled: true,
		Audience:        ServiceAccessControlService,
	}
	h := testHolder(t, cfg, src)

	handler := Middleware(h, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("handler must not run unauthenticated after invalid mesh token")
	}))
	req := httptest.NewRequest(http.MethodGet, "http://example/api", nil)
	req.Header.Set(IdentityHeader, "not.a.valid.jwt")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusUnauthorized, rr.Code)
	require.Contains(t, rr.Body.String(), ErrInvalidMeshToken.Error())
}

func TestMiddlewareFallbackOnInvalidToken(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{
		InboundEnabled:  true,
		FallbackEnabled: true,
		Audience:        ServiceAccessControlService,
	}
	h := testHolder(t, cfg, src)

	fallbackCalled := false
	fallback := FallbackAuthFunc(func(r *http.Request) (Principal, error) {
		fallbackCalled = true
		return Principal{Type: PrincipalTypeService, Name: "hmac-caller"}, nil
	})

	var got Principal
	handler := Middleware(h, fallback)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, ok := PrincipalFromContext(r.Context())
		require.True(t, ok)
		got = p
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://example/api", nil)
	req.Header.Set(IdentityHeader, "not-a-jwt")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	require.True(t, fallbackCalled)
	require.Equal(t, "hmac-caller", got.Name)
}

func TestMiddlewareRejectWithoutHeader(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{
		InboundEnabled:          true,
		FallbackEnabled:         true,
		RejectWithoutMeshHeader: true,
		Audience:                ServiceAccessControlService,
	}
	h := testHolder(t, cfg, src)

	handler := Middleware(h, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://example/api", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestMiddlewareNoFallbackOnFailure(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{
		InboundEnabled:  true,
		FallbackEnabled: false,
		Audience:        ServiceAccessControlService,
	}
	h := testHolder(t, cfg, src)

	handler := Middleware(h, FallbackAuthFunc(func(r *http.Request) (Principal, error) {
		return Principal{}, errors.New("should not be called")
	}))(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "http://example/api", nil)
	req.Header.Set(IdentityHeader, "bad")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestBootstrapNoop(t *testing.T) {
	h, err := Bootstrap(context.Background(), Config{}, nil)
	require.NoError(t, err)
	require.True(t, h.IsNoop())
	require.Nil(t, h.Generator())
}

func TestRoundTripperStampsHeader(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{OutboundEnabled: true}
	h := testHolder(t, cfg, src)

	var sawHeader string
	next := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		sawHeader = req.Header.Get(IdentityHeader)
		return &http.Response{StatusCode: 200, Body: http.NoBody, Request: req}, nil
	})

	rt := NewRoundTripper(h, OutboundConfig{TargetServiceID: ServiceAccessControlService}, next)
	req, err := http.NewRequest(http.MethodGet, "http://access-control/api", nil)
	require.NoError(t, err)
	resp, err := rt.RoundTrip(req)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.NotEmpty(t, sawHeader)
}

func TestRoundTripperMissingTargetDualMode(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{OutboundEnabled: true, OutboundOnly: false}
	h := testHolder(t, cfg, src)

	called := false
	next := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		called = true
		require.Empty(t, req.Header.Get(IdentityHeader))
		return &http.Response{StatusCode: 200, Body: http.NoBody, Request: req}, nil
	})

	rt := NewRoundTripper(h, OutboundConfig{}, next)
	req, _ := http.NewRequest(http.MethodGet, "http://svc/", nil)
	_, err := rt.RoundTrip(req)
	require.NoError(t, err)
	require.True(t, called)
}

func TestRoundTripperMissingTargetMeshOnly(t *testing.T) {
	spiffeURI := "spiffe://test.harness.io/qa/NextGenManager"
	svid, bundle, _ := newTestCA(t, spiffeURI)
	td := TrustDomainFromSPIFFE(spiffeURI)
	src := &StaticSource{SVID: svid, Bundles: map[string]*BundleRef{td: bundle}}

	cfg := Config{OutboundEnabled: true, OutboundOnly: true}
	h := testHolder(t, cfg, src)

	rt := NewRoundTripper(h, OutboundConfig{}, http.DefaultTransport)
	req, _ := http.NewRequest(http.MethodGet, "http://svc/", nil)
	_, err := rt.RoundTrip(req)
	require.ErrorIs(t, err, ErrMissingTargetService)
}

func TestRoundTripperPassthroughWhenOutboundDisabled(t *testing.T) {
	h, err := Bootstrap(context.Background(), Config{}, nil)
	require.NoError(t, err)

	called := false
	next := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		called = true
		require.Empty(t, req.Header.Get(IdentityHeader))
		return &http.Response{StatusCode: 204, Body: http.NoBody, Request: req}, nil
	})
	rt := NewRoundTripper(h, OutboundConfig{TargetServiceID: ServiceAccessControlService}, next)
	req, _ := http.NewRequest(http.MethodGet, "http://svc/", nil)
	_, err = rt.RoundTrip(req)
	require.NoError(t, err)
	require.True(t, called)
}

func TestRoundTripperMintFailDualContinues(t *testing.T) {
	src := &failingSource{err: errors.New("no svid")}
	cfg := Config{OutboundEnabled: true, OutboundOnly: false}
	h := testHolder(t, cfg, src)

	called := false
	next := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		called = true
		require.Empty(t, req.Header.Get(IdentityHeader))
		return &http.Response{StatusCode: 200, Body: http.NoBody, Request: req}, nil
	})
	rt := NewRoundTripper(h, OutboundConfig{TargetServiceID: ServiceAccessControlService}, next)
	req, _ := http.NewRequest(http.MethodGet, "http://svc/", nil)
	_, err := rt.RoundTrip(req)
	require.NoError(t, err)
	require.True(t, called)
}

func TestRoundTripperMintFailMeshOnlyErrors(t *testing.T) {
	src := &failingSource{err: errors.New("no svid")}
	cfg := Config{OutboundEnabled: true, OutboundOnly: true}
	h := testHolder(t, cfg, src)

	rt := NewRoundTripper(h, OutboundConfig{TargetServiceID: ServiceAccessControlService}, http.DefaultTransport)
	req, _ := http.NewRequest(http.MethodGet, "http://svc/", nil)
	_, err := rt.RoundTrip(req)
	require.Error(t, err)
	require.Contains(t, err.Error(), "mesh outbound mint failed")
}

type failingSource struct{ err error }

func (f *failingSource) Current() (*Snapshot, error)                       { return nil, f.err }
func (f *failingSource) BundleForTrustDomain(string) (*BundleRef, error)   { return nil, f.err }
func (f *failingSource) Close() error                                      { return nil }

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
