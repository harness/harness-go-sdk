/*
 * Split API client
 *
 * Package skeleton and HTTP client for the Split Admin API with rate limiting and retry support.
 */
package split

import (
	"fmt"
	"net/http"

	"github.com/harness/harness-go-sdk/harness"
	"github.com/harness/harness-go-sdk/harness/helpers"
)

// Configuration holds settings for the Split API client.
type Configuration struct {
	AccountId     string            `json:"accountId,omitempty"`
	ApiKey        string            `json:"apiKey,omitempty"`
	BasePath      string            `json:"basePath,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

// DefaultSplitBasePath is the default base URL for the Split Admin API.
const DefaultSplitBasePath = "https://api.split.io"

// NewDefaultConfiguration returns a configuration with default UserAgent, BasePath, and empty DefaultHeader.
// HTTPClient is left nil so the client uses a default client with the rate-limit/retry transport.
func NewDefaultConfiguration() *Configuration {
	return &Configuration{
		BasePath:      DefaultSplitBasePath,
		DefaultHeader: make(map[string]string),
		UserAgent:     fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion),
	}
}

// NewConfiguration returns a configuration populated from the same environment variables as the rest of the SDK:
// HARNESS_ACCOUNT_ID and HARNESS_PLATFORM_API_KEY (see harness/helpers.EnvVars). BasePath, UserAgent, and DefaultHeader
// are set to the same defaults as NewDefaultConfiguration.
func NewConfiguration() *Configuration {
	cfg := NewDefaultConfiguration()
	cfg.AccountId = helpers.EnvVars.AccountId.Get()
	cfg.ApiKey = helpers.EnvVars.PlatformApiKey.Get()
	return cfg
}

// AddDefaultHeader adds a default header that will be sent with every request.
func (c *Configuration) AddDefaultHeader(key string, value string) {
	if c.DefaultHeader == nil {
		c.DefaultHeader = make(map[string]string)
	}
	c.DefaultHeader[key] = value
}
