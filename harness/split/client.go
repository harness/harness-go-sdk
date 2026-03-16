/*
 * Split API client
 *
 * APIClient and constructor for the Split Admin API. Services (e.g. Workspaces, Segments)
 * will be added in later tickets and will use the same HTTP client for rate-limit and retry.
 */
package split

import (
	"net/http"
)

// APIClient manages communication with the Split Admin API.
// Use NewAPIClient(cfg) to create one. The embedded HTTP client applies base URL,
// headers, auth, and rate-limit/retry behavior for all requests.
type APIClient struct {
	cfg *Configuration

	// HTTPClient is the client used for all requests. Its transport handles
	// 429 retry (after X-RateLimit-Reset-Seconds-*) and optional 5xx retry.
	HTTPClient *http.Client

	// AccountId and ApiKey are copied from config for use by services.
	// Pass AccountId on relevant Split API endpoints (e.g. query param or header per https://docs.split.io/reference/).
	AccountId string
	ApiKey    string
	BasePath  string

	// Environments provides access to environments in a workspace.
	Environments *EnvironmentsService

	// TrafficTypes provides access to traffic types in a workspace.
	TrafficTypes *TrafficTypesService

	// Attributes provides access to schema attributes for a traffic type.
	Attributes *AttributesService

	// FlagSets provides access to flag sets (API v3).
	FlagSets *FlagSetsService

	// Segments provides access to segments in a workspace.
	Segments *SegmentsService
}

// NewAPIClient creates a new Split API client from the given configuration.
// If cfg.HTTPClient is nil, a new http.Client is created with the rate-limit/retry transport.
// No Option pattern; match existing SDK style (e.g. svcdiscovery.NewAPIClient).
func NewAPIClient(cfg *Configuration) *APIClient {
	baseTransport := http.DefaultTransport
	if cfg.HTTPClient != nil && cfg.HTTPClient.Transport != nil {
		baseTransport = cfg.HTTPClient.Transport
	}
	transport := newRateLimitRetryTransport(cfg, baseTransport)

	client := &http.Client{Transport: transport}
	if cfg.HTTPClient != nil {
		client.Timeout = cfg.HTTPClient.Timeout
		client.CheckRedirect = cfg.HTTPClient.CheckRedirect
		client.Jar = cfg.HTTPClient.Jar
	}

	c := &APIClient{
		cfg:        cfg,
		HTTPClient: client,
		AccountId:  cfg.AccountId,
		ApiKey:     cfg.ApiKey,
		BasePath:   cfg.BasePath,
	}
	c.Environments = &EnvironmentsService{client: c}
	c.TrafficTypes = &TrafficTypesService{client: c}
	c.Attributes = &AttributesService{client: c}
	c.FlagSets = &FlagSetsService{client: c}
	c.Segments = &SegmentsService{client: c}
	return c
}

// Do sends an HTTP request and returns the response. It is the common entrypoint
// for future services (Workspaces, Segments, etc.) so all calls get rate-limit and retry.
func (c *APIClient) Do(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}
