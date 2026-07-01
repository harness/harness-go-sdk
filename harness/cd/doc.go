// Package cd provides the Go client for Harness FirstGen (Continuous Delivery) APIs.
//
// This package contains clients for interacting with Harness FirstGen services
// including applications, services, environments, workflows, pipelines, and more.
//
// # Overview
//
// The cd package provides a high-level API client for Harness FirstGen. Unlike
// the nextgen package which is auto-generated, this package provides hand-crafted
// clients with convenience methods for common operations.
//
// # Getting Started
//
// Create an API client with environment variables:
//
//	import "github.com/harness/harness-go-sdk/harness/cd"
//
//	client, err := cd.NewClient(&cd.Config{
//	    AccountId: os.Getenv("HARNESS_ACCOUNT_ID"),
//	    APIKey:    os.Getenv("HARNESS_API_KEY"),
//	    Endpoint:  "https://app.harness.io",
//	})
//
// # Available Clients
//
// The ApiClient provides access to various service clients:
//
//   - ApplicationClient - Manage applications
//   - ServiceClient - Manage services
//   - CloudProviderClient - Manage cloud providers
//   - ConfigAsCodeClient - Config-as-code operations
//   - ConnectorClient - Manage connectors
//   - DelegateClient - Manage delegates
//   - ExecutionClient - Pipeline executions
//   - SecretClient - Manage secrets
//   - SSOClient - Single sign-on configuration
//   - UserClient - User management
//   - TagClient - Tagging operations
//   - TriggerClient - Pipeline triggers
//   - ApprovalClient - Approval workflows
//
// # Example: Get Application by Name
//
//	app, err := client.ApplicationClient.GetApplicationByName("my-app")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(app.Id, app.Name)
//
// # Example: Create a Service
//
//	svc := &cd.Service{
//	    Name:           "my-service",
//	    ApplicationId:  appId,
//	    DeploymentType: cd.DeploymentTypes.Kubernetes,
//	    ArtifactType:   cd.ArtifactTypes.Docker,
//	}
//	newService, err := client.ServiceClient.CreateService(svc)
//
// # Configuration Options
//
// The Config struct supports the following options:
//
//   - AccountId: Your Harness account identifier (required)
//   - APIKey: Your API key for authentication (required)
//   - Endpoint: API endpoint URL (default: https://app.harness.io)
//   - HTTPClient: Custom retryable HTTP client
//   - DebugLogging: Enable debug logging
//   - Logger: Custom logger instance
//
// # More Information
//
// For FirstGen API documentation, visit:
// https://developer.harness.io/docs/first-gen/
package cd
