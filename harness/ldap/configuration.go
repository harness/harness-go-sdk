// Package ldap provides a hand-written client for the Harness next-generation
// LDAP settings API (spec-server v1, served by ng-manager):
//
//	/v1/ldap-settings
//
// This is the NEW account-level LDAP settings surface gated by the
// PL_ENABLE_NG_LDAP_SETTINGS feature flag (NGLdapSettings store). It is
// unrelated to the legacy /ng/api/authentication-settings/ldap/settings API,
// which serves the CG LdapSettings store. Requests authenticate with the
// x-api-key header and carry the account in the Harness-Account header,
// mirroring the harness/idp and harness/po clients.
package ldap

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/harness/harness-go-sdk/harness"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/harness-go-sdk/logging"
	"github.com/hashicorp/go-retryablehttp"
)

type Configuration struct {
	AccountId     string            `json:"accountId,omitempty"`
	ApiKey        string            `json:"apiKey,omitempty"`
	BasePath      string            `json:"basePath,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *retryablehttp.Client
	Logger        *log.Logger
	DebugLogging  bool
}

func NewConfiguration() *Configuration {
	logger := logging.NewLogger()
	if helpers.EnvVars.TfLog.Get() == "DEBUG" {
		logger.SetLevel(log.DebugLevel)
	}
	cfg := &Configuration{
		AccountId:     helpers.EnvVars.AccountId.Get(),
		ApiKey:        helpers.EnvVars.PlatformApiKey.Get(),
		BasePath:      helpers.EnvVars.Endpoint.GetWithDefault(utils.BaseUrl),
		DefaultHeader: make(map[string]string),
		HTTPClient:    utils.GetDefaultHttpClient(logger),
		Logger:        logger,
		UserAgent:     fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion),
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
