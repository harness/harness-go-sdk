// Package split provides the Go client for the Split Admin API (Feature Flags and Experimentation).
//
// This package includes configuration, an HTTP client with rate-limit and retry behavior,
// and uses the same environment variables as the rest of the SDK (HARNESS_ACCOUNT_ID,
// HARNESS_PLATFORM_API_KEY) when using NewConfiguration().
//
// # Overview
//
// The Split client handles 429 rate limiting by reading X-RateLimit-Reset-Seconds-Org
// and X-RateLimit-Reset-Seconds-IP, waiting, and retrying. It also retries on 5xx
// with exponential backoff. Authentication uses the x-api-key header (Harness SAT, PAT or
// Split Admin API key).
//
// # Getting Started
//
// Create a configuration and API client:
//
//	import "github.com/harness/harness-go-sdk/harness/split"
//
//	// From env: HARNESS_ACCOUNT_ID, HARNESS_PLATFORM_API_KEY
//	cfg := split.NewConfiguration()
//	client := split.NewAPIClient(cfg)
//
//	req, _ := http.NewRequest("GET", cfg.BasePath+"/internal/api/v2/workspaces", nil)
//	resp, err := client.Do(req)
//
// For custom settings (no env vars), use NewDefaultConfiguration() and set ApiKey
// and BasePath as needed.
//
// # Account ID
//
// Pass the client's AccountId on relevant Split API endpoints (e.g. as query parameter
// or header per the Split API reference). Future services (Workspaces, Segments, etc.)
// will use client.AccountId when building requests.
//
// # Integration tests
//
// Integration tests live in split_sdk_test.go and are skipped when credentials
// are missing or when running with -short. Set HARNESS_ACCOUNT_ID and
// HARNESS_PLATFORM_API_KEY (or pass -account and -api-key to the test binary)
// to run them. Example:
//
//	go test -v ./harness/split/... -run TestSplitSDK
//	go test -short ./harness/split/...   # skip integration tests
//
// # Documentation
//
// Split API reference: https://docs.split.io/reference/
package split
