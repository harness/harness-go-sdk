// Package sdk provides the official Go SDK for interacting with the Harness Platform.
//
// The Harness Go SDK enables programmatic access to Harness services including
// CI/CD pipelines, feature flags, cloud cost management, chaos engineering, and more.
//
// # Installation
//
// Use go get to install the SDK:
//
//	go get github.com/harness/harness-go-sdk
//
// # Quick Start
//
// To start using the SDK, import the relevant sub-packages based on your needs:
//
//	import (
//	    "github.com/harness/harness-go-sdk/harness/nextgen"
//	    "github.com/harness/harness-go-sdk/harness/cd"
//	)
//
// # Available Packages
//
// The SDK is organized into the following packages:
//
//   - [github.com/harness/harness-go-sdk/harness/nextgen] - NextGen Platform APIs (pipelines, connectors, secrets, etc.)
//   - [github.com/harness/harness-go-sdk/harness/cd] - Continuous Delivery (FirstGen) APIs
//   - [github.com/harness/harness-go-sdk/harness/chaos] - Chaos Engineering APIs
//   - [github.com/harness/harness-go-sdk/harness/code] - Code Repository APIs
//   - [github.com/harness/harness-go-sdk/harness/helpers] - Helper utilities and common functions
//   - [github.com/harness/harness-go-sdk/harness/utils] - Utility functions
//   - [github.com/harness/harness-go-sdk/logging] - Logging utilities
//
// # Authentication
//
// Set the following environment variables to configure authentication:
//
//   - HARNESS_ACCOUNT_ID: Your Harness account identifier (required)
//   - HARNESS_API_KEY: Your API key for authentication (required)
//   - HARNESS_ENDPOINT: API endpoint URL (optional, defaults to https://app.harness.io)
//
// # Example Usage
//
// Creating a NextGen API client:
//
//	cfg := nextgen.NewConfiguration()
//	cfg.BasePath = "https://app.harness.io/gateway"
//	cfg.AddDefaultHeader("x-api-key", os.Getenv("HARNESS_API_KEY"))
//
//	client := nextgen.NewAPIClient(cfg)
//
//	// List organizations
//	orgs, resp, err := client.OrganizationApi.GetOrganizationList(
//	    context.Background(),
//	    accountIdentifier,
//	    nil,
//	)
//
// # Support
//
// For issues and feature requests, please visit:
// https://github.com/harness/harness-go-sdk/issues
//
// For more information about Harness, visit:
// https://developer.harness.io/
package sdk