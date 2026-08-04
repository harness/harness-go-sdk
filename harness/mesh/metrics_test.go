package mesh

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
)

func TestMetricsTrustDomainCap(t *testing.T) {
	m := testMetrics(t)
	for i := 0; i < MaxTrustDomainLabels+5; i++ {
		m.RecordBundleRefresh(OutcomeSuccess, fmt.Sprintf("td-%d", i))
	}
	other := testutil.ToFloat64(m.bundleRefresh.WithLabelValues(OutcomeSuccess, LabelOther))
	require.Greater(t, other, 0.0)
}

func TestMetricsMintAndOutbound(t *testing.T) {
	m := testMetrics(t)
	done := m.StartMint()
	done()
	m.RecordMint(OutcomeSuccess)
	m.RecordOutboundConfigError(ConfigErrorMissingTargetService)
	m.RecordAuthTransport(TransportSPIRE)
	m.RecordHMACFallback(ReasonExpired)

	require.Equal(t, 1.0, testutil.ToFloat64(m.mintTotal.WithLabelValues(OutcomeSuccess)))
	require.Equal(t, 1.0, testutil.ToFloat64(m.outboundConfigError.WithLabelValues(ConfigErrorMissingTargetService)))
	require.Equal(t, 1.0, testutil.ToFloat64(m.authTransport.WithLabelValues(TransportSPIRE)))
	require.Equal(t, 1.0, testutil.ToFloat64(m.hmacFallback.WithLabelValues(ReasonExpired)))
}
