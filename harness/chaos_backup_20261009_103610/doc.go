// Package chaos provides the Go client for Harness Chaos Engineering APIs.
//
// This package contains auto-generated API clients for interacting with
// Harness Chaos Engineering services, enabling you to run chaos experiments,
// manage chaos infrastructure, and integrate chaos testing into your pipelines.
//
// # Overview
//
// Harness Chaos Engineering helps you build resilient systems by testing how
// applications respond to failures. This package provides programmatic access
// to create and manage chaos experiments, infrastructure, and probes.
//
// # Getting Started
//
// Create a configuration and API client:
//
//	import "github.com/harness/harness-go-sdk/harness/chaos"
//
//	cfg := chaos.NewConfiguration()
//	cfg.BasePath = "https://app.harness.io/gateway/chaos/manager/api"
//	cfg.AddDefaultHeader("x-api-key", os.Getenv("HARNESS_API_KEY"))
//
//	client := chaos.NewAPIClient(cfg)
//
// # Key Features
//
// The chaos package enables you to:
//
//   - Register and manage chaos infrastructure
//   - Create and run chaos experiments
//   - Define chaos probes for validation
//   - Configure chaos hubs for experiment templates
//   - Integrate with CI/CD pipelines
//
// # Available API Services
//
//   - ChaosSdkApi - Core chaos operations (infrastructure, experiments)
//   - DefaultApi - Additional chaos functionality (notes, bulk operations)
//
// # Example: Register Chaos Infrastructure
//
//	infra := chaos.RegisterInfraRequest{
//	    Name:                  "my-k8s-cluster",
//	    EnvironmentIdentifier: envId,
//	    PlatformName:          "Kubernetes",
//	}
//	resp, httpResp, err := client.ChaosSdkApi.RegisterInfraV2(ctx, infra)
//
// # More Information
//
// For Chaos Engineering documentation, visit:
// https://developer.harness.io/docs/chaos-engineering/
package chaos
