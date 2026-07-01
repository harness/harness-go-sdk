// Package code provides the Go client for Harness Code Repository APIs.
//
// This package contains auto-generated API clients for interacting with
// Harness Code Repository services, enabling you to manage repositories,
// branches, pull requests, and other source code management operations.
//
// # Overview
//
// Harness Code is a source code management solution built into the Harness
// Platform. This package provides programmatic access to repositories,
// webhooks, and code review features.
//
// # Getting Started
//
// Create a configuration and API client:
//
//	import "github.com/harness/harness-go-sdk/harness/code"
//
//	cfg := code.NewConfiguration()
//	cfg.BasePath = "https://app.harness.io/gateway/code/api"
//	cfg.AddDefaultHeader("x-api-key", os.Getenv("HARNESS_API_KEY"))
//
//	client := code.NewAPIClient(cfg)
//
// # Key Features
//
// The code package enables you to:
//
//   - Manage repositories
//   - Handle branches and commits
//   - Create and manage pull requests
//   - Configure webhooks
//   - Access repository metadata
//
// # More Information
//
// For Code Repository documentation, visit:
// https://developer.harness.io/docs/code-repository/
package code
