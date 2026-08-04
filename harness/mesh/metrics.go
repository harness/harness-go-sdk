package mesh

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics holds Prometheus instruments for the mesh path.
type Metrics struct {
	mintTotal           *prometheus.CounterVec
	mintDuration        prometheus.Histogram
	validateTotal       *prometheus.CounterVec
	validateDuration    prometheus.Histogram
	hmacFallback        *prometheus.CounterVec
	bundleRefresh       *prometheus.CounterVec
	authTransport       *prometheus.CounterVec
	outboundConfigError *prometheus.CounterVec

	mu                 sync.Mutex
	observedTrustDomains map[string]struct{}
}

// NewMetrics registers mesh metrics on reg (use prometheus.DefaultRegisterer in production).
// If a metric is already registered, NewMetrics returns the existing collectors via MustRegister panic
// avoidance by using a private registry pattern — callers should use RegisterOnce or DefaultMetrics.
func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		observedTrustDomains: make(map[string]struct{}),
		mintTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: MetricMintTotal,
			Help: "Outbound mesh JWT mint attempts",
		}, []string{"outcome"}),
		mintDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    MetricMintDuration,
			Help:    "Outbound mesh JWT mint latency in seconds",
			Buckets: []float64{0.0005, 0.001, 0.002, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0},
		}),
		validateTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: MetricValidateTotal,
			Help: "Inbound mesh JWT validation attempts",
		}, []string{"outcome", "caller", "audience", "trust_domain"}),
		validateDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    MetricValidateDuration,
			Help:    "Inbound mesh JWT validation latency in seconds",
			Buckets: []float64{0.00005, 0.0001, 0.0005, 0.001, 0.005, 0.01, 0.025, 0.05, 0.1},
		}),
		hmacFallback: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: MetricHMACFallbackTotal,
			Help: "Mesh validation failed and the filter fell through to the HMAC path",
		}, []string{"reason"}),
		bundleRefresh: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: MetricBundleRefreshTotal,
			Help: "X.509 trust-bundle refresh events observed from the SPIRE Workload API, per trust domain",
		}, []string{"outcome", "trust_domain"}),
		authTransport: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: MetricAuthTransportTotal,
			Help: "Inbound requests by the transport that actually authenticated them (spire|hmac)",
		}, []string{"transport"}),
		outboundConfigError: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: MetricOutboundConfigError,
			Help: "Outbound-side mesh configuration errors, per reason",
		}, []string{"reason"}),
	}
	if reg != nil {
		reg.MustRegister(
			m.mintTotal,
			m.mintDuration,
			m.validateTotal,
			m.validateDuration,
			m.hmacFallback,
			m.bundleRefresh,
			m.authTransport,
			m.outboundConfigError,
		)
	}
	return m
}

var (
	defaultMetricsOnce sync.Once
	defaultMetrics     *Metrics
)

// DefaultMetrics returns process-wide metrics registered on the default Prometheus registry.
func DefaultMetrics() *Metrics {
	defaultMetricsOnce.Do(func() {
		defaultMetrics = NewMetrics(prometheus.DefaultRegisterer)
	})
	return defaultMetrics
}

// StartMint returns a function that observes mint duration when called.
func (m *Metrics) StartMint() func() {
	if m == nil {
		return func() {}
	}
	start := time.Now()
	return func() { m.mintDuration.Observe(time.Since(start).Seconds()) }
}

// RecordMint increments mint total.
func (m *Metrics) RecordMint(outcome string) {
	if m == nil {
		return
	}
	m.mintTotal.WithLabelValues(orUnknown(outcome)).Inc()
}

// StartValidate returns a function that observes validate duration when called.
func (m *Metrics) StartValidate() func() {
	if m == nil {
		return func() {}
	}
	start := time.Now()
	return func() { m.validateDuration.Observe(time.Since(start).Seconds()) }
}

// RecordValidate increments validate total with cardinality-safe trust_domain.
func (m *Metrics) RecordValidate(outcome, caller, audience, trustDomain string) {
	if m == nil {
		return
	}
	m.validateTotal.WithLabelValues(
		orUnknown(outcome),
		orUnknown(caller),
		orUnknown(audience),
		m.knownTrustDomainLabel(trustDomain),
	).Inc()
}

// RecordHMACFallback increments hmac fallback counter.
func (m *Metrics) RecordHMACFallback(reason string) {
	if m == nil {
		return
	}
	m.hmacFallback.WithLabelValues(orUnknown(reason)).Inc()
}

// RecordBundleRefresh records bundle load/rotation; success admits trust domain to label set.
func (m *Metrics) RecordBundleRefresh(outcome, trustDomain string) {
	if m == nil {
		return
	}
	var label string
	if outcome == OutcomeSuccess {
		label = m.registerTrustDomainLabel(trustDomain)
	} else {
		label = m.knownTrustDomainLabel(trustDomain)
	}
	m.bundleRefresh.WithLabelValues(orUnknown(outcome), label).Inc()
}

// RecordAuthTransport increments auth transport counter.
func (m *Metrics) RecordAuthTransport(transport string) {
	if m == nil {
		return
	}
	m.authTransport.WithLabelValues(orUnknown(transport)).Inc()
}

// RecordOutboundConfigError increments outbound config error counter.
func (m *Metrics) RecordOutboundConfigError(reason string) {
	if m == nil {
		return
	}
	m.outboundConfigError.WithLabelValues(orUnknown(reason)).Inc()
}

func (m *Metrics) registerTrustDomainLabel(trustDomain string) string {
	if trustDomain == "" {
		return LabelUnknown
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.observedTrustDomains[trustDomain]; ok {
		return trustDomain
	}
	if len(m.observedTrustDomains) >= MaxTrustDomainLabels {
		return LabelOther
	}
	m.observedTrustDomains[trustDomain] = struct{}{}
	return trustDomain
}

func (m *Metrics) knownTrustDomainLabel(trustDomain string) string {
	if trustDomain == "" {
		return LabelUnknown
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.observedTrustDomains[trustDomain]; !ok {
		return LabelUnknown
	}
	return trustDomain
}

// ResetTrustDomainLabels clears the allowlist (tests only).
func (m *Metrics) ResetTrustDomainLabels() {
	if m == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.observedTrustDomains = make(map[string]struct{})
}

func orUnknown(v string) string {
	if v == "" {
		return LabelUnknown
	}
	return v
}
