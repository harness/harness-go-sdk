package nextgen

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/harness-go-sdk/logging"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
)

// NewAgentClient creates an APIClient configured for AI agent usage:
// - Rate limit aware retry (respects Retry-After and X-RateLimit-Reset headers)
// - Higher retry count for transient failures
// - Appropriate timeouts
//
// This is the recommended constructor for programmatic/agent usage.
func NewAgentClient(cfg *Configuration) *APIClient {
	if cfg == nil {
		cfg = NewConfiguration()
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = newAgentHTTPClient(cfg)
	}
	return NewAPIClient(cfg)
}

func newAgentHTTPClient(cfg *Configuration) *retryablehttp.Client {
	logger := cfg.Logger
	if logger == nil {
		logger = logging.NewLogger()
		if helpers.EnvVars.TfLog.Get() == "DEBUG" {
			logger.SetLevel(log.DebugLevel)
		}
	}

	httpClient := retryablehttp.NewClient()
	httpClient.Logger = retryablehttp.LeveledLogger(&logging.Logger{Logger: logger})
	httpClient.HTTPClient.Transport = logging.NewTransport(harness.SDKName, logger, cleanhttp.DefaultPooledClient().Transport)
	httpClient.RetryMax = maxRateLimitRetries
	httpClient.CheckRetry = RateLimitCheckRetry
	httpClient.Backoff = RateLimitBackoff

	if cfg.UserAgent == "" {
		cfg.UserAgent = fmt.Sprintf("%s-%s (agent)", harness.SDKName, harness.SDKVersion)
	}
	if cfg.BasePath == "" {
		cfg.BasePath = helpers.EnvVars.Endpoint.GetWithDefault(utils.BaseUrl)
	}

	return httpClient
}
