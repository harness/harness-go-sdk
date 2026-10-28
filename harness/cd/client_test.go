package cd

import (
	"flag"
	"os"
	"sync"
	"testing"

	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() && helpers.EnvVars.ApiKey.Get() == "" {
		// When running unit tests in short mode without live FirstGen CD credentials,
		// filter to unit tests so credential-free tests (TestClientRequire*) run and contribute coverage.
		if f := flag.Lookup("test.run"); f == nil || f.Value.String() == "" {
			flag.Set("test.run", "^TestClient")
		}
	}
	os.Exit(m.Run())
}

var configureClient sync.Once
var apiClient *ApiClient

func getClient() *ApiClient {
	configureClient.Do(func() {
		var err error

		apiClient, err = NewClient(DefaultConfig())

		if err != nil {
			panic(err)
		}
	})

	return apiClient
}

func GetUnauthorizedClient() *ApiClient {
	cfg := DefaultConfig()
	cfg.APIKey = "INVALID_API_KEY"

	c, _ := NewClient(cfg)
	return c
}

func TestClientRequireApiKey_Config(t *testing.T) {
	cfg := DefaultConfig()
	cfg.AccountId = "ACCOUNT_ID"
	cfg.APIKey = ""

	_, err := NewClient(cfg)

	require.Error(t, err, InvalidConfigError{})
	require.Equal(t, "ApiKey", err.(InvalidConfigError).Field)

	cfg.APIKey = "APIKEY"
	c, err := NewClient(cfg)

	require.NoError(t, err)
	require.Equal(t, c.Configuration.APIKey, cfg.APIKey)
}

func TestClientRequireAccountId_Config(t *testing.T) {
	cfg := DefaultConfig()
	cfg.AccountId = ""
	cfg.APIKey = "APIKEY"

	_, err := NewClient(cfg)

	require.Error(t, err, InvalidConfigError{})
	require.Equal(t, "AccountId", err.(InvalidConfigError).Field)

	cfg.AccountId = "ACCOUNT_ID"
	c, err := NewClient(cfg)

	require.NoError(t, err)
	require.Equal(t, c.Configuration.AccountId, cfg.AccountId)
}
