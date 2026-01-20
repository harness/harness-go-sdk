// Package helpers provides common utilities and helper functions for the Harness Go SDK.
//
// This package contains shared constants, types, and utility functions used
// across the SDK for HTTP operations, environment variable handling, and
// error management.
//
// # HTTP Helpers
//
// Common HTTP headers and query parameters are provided as typed constants:
//
//	import "github.com/harness/harness-go-sdk/harness/helpers"
//
//	// Using HTTP headers
//	req.Header.Set(helpers.HTTPHeaders.ContentType.String(), helpers.HTTPHeaders.ApplicationJson.String())
//	req.Header.Set(helpers.HTTPHeaders.ApiKey.String(), apiKey)
//
//	// Using query parameters
//	params.Add(helpers.QueryParameters.AccountId.String(), accountId)
//
// # Environment Variables
//
// The package provides utilities for reading configuration from environment variables:
//
//	accountId := helpers.GetEnv("HARNESS_ACCOUNT_ID", "")
//	endpoint := helpers.GetEnv("HARNESS_ENDPOINT", "https://app.harness.io")
//
// # Error Types
//
// Common error types are defined for consistent error handling across the SDK:
//
//	if errors.Is(err, helpers.ErrNotFound) {
//	    // Handle not found case
//	}
package helpers
