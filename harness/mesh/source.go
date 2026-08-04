package mesh

import (
	"context"
	"crypto"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

// Snapshot is an immutable view of the workload's current X509-SVID.
type Snapshot struct {
	Chain      []*x509.Certificate // leaf-first
	PrivateKey crypto.Signer
	Kid        string
	SpiffeID   string
}

// Leaf returns chain[0] or nil.
func (s *Snapshot) Leaf() *x509.Certificate {
	if s == nil || len(s.Chain) == 0 {
		return nil
	}
	return s.Chain[0]
}

// BundleRef is a trust-domain bundle plus an identity token used for cache invalidation.
type BundleRef struct {
	Authorities   []*x509.Certificate
	BundleVersion string // opaque identity; changes when underlying bundle object changes
}

// IdentitySource supplies SVIDs and trust bundles (SPIRE Workload API or test fake).
type IdentitySource interface {
	Current() (*Snapshot, error)
	BundleForTrustDomain(td string) (*BundleRef, error)
	Close() error
}

// WorkloadSource wraps go-spiffe workloadapi.X509Source.
type WorkloadSource struct {
	src *workloadapi.X509Source
}

// NewWorkloadSource connects to the SPIRE Agent Workload API.
func NewWorkloadSource(ctx context.Context, socketPath string) (*WorkloadSource, error) {
	opts := []workloadapi.X509SourceOption{}
	if socketPath != "" {
		opts = append(opts, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)))
	}
	src, err := workloadapi.NewX509Source(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize SPIFFE Workload API source. Is the SPIRE Agent socket mounted, or SPIFFE_ENDPOINT_SOCKET set?: %w", err)
	}
	return &WorkloadSource{src: src}, nil
}

// Current returns the latest SVID snapshot.
func (w *WorkloadSource) Current() (*Snapshot, error) {
	svid, err := w.src.GetX509SVID()
	if err != nil {
		return nil, err
	}
	chain := svid.Certificates
	if len(chain) == 0 {
		return nil, fmt.Errorf("empty SVID chain")
	}
	leaf := chain[0]
	return &Snapshot{
		Chain:      append([]*x509.Certificate(nil), chain...),
		PrivateKey: svid.PrivateKey,
		Kid:        computeKid(leaf),
		SpiffeID:   svid.ID.String(),
	}, nil
}

// BundleForTrustDomain returns trust anchors for td.
func (w *WorkloadSource) BundleForTrustDomain(td string) (*BundleRef, error) {
	domain, err := spiffeid.TrustDomainFromString(td)
	if err != nil {
		return nil, err
	}
	bundle, err := w.src.GetX509BundleForTrustDomain(domain)
	if err != nil {
		return nil, err
	}
	if bundle == nil {
		return nil, fmt.Errorf("nil bundle for trust domain %s", td)
	}
	auths := bundle.X509Authorities()
	if len(auths) == 0 {
		return nil, fmt.Errorf("empty X.509 trust bundle for trust domain %s", td)
	}
	return &BundleRef{
		Authorities:   append([]*x509.Certificate(nil), auths...),
		BundleVersion: fmt.Sprintf("%p", bundle),
	}, nil
}

// Close closes the Workload API source.
func (w *WorkloadSource) Close() error {
	if w == nil || w.src == nil {
		return nil
	}
	return w.src.Close()
}

func computeKid(cert *x509.Certificate) string {
	sum := sha256.Sum256(cert.Raw)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

// StaticSource is a test IdentitySource with a fixed SVID and trust anchors.
type StaticSource struct {
	SVID    *Snapshot
	Bundles map[string]*BundleRef // trust domain -> bundle
}

// Current implements IdentitySource.
func (s *StaticSource) Current() (*Snapshot, error) {
	if s == nil || s.SVID == nil {
		return nil, fmt.Errorf("no SVID")
	}
	return s.SVID, nil
}

// BundleForTrustDomain implements IdentitySource.
func (s *StaticSource) BundleForTrustDomain(td string) (*BundleRef, error) {
	if s == nil || s.Bundles == nil {
		return nil, fmt.Errorf("no X.509 trust bundle for trust domain %s", td)
	}
	b, ok := s.Bundles[td]
	if !ok || b == nil || len(b.Authorities) == 0 {
		return nil, fmt.Errorf("no X.509 trust bundle for trust domain %s", td)
	}
	return b, nil
}

// Close implements IdentitySource.
func (s *StaticSource) Close() error { return nil }
