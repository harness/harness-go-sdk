// Package nextgen provides the Go client for the Harness NextGen Platform APIs.
//
// This package contains auto-generated API clients for interacting with Harness
// NextGen services including pipelines, connectors, secrets, organizations,
// projects, environments, services, and more.
//
// # Overview
//
// The NextGen package provides comprehensive access to the Harness Platform's
// modern API surface. It is generated from the OpenAPI Specification v3.0 and
// includes clients for all major Harness modules.
//
// # Getting Started
//
// Create a configuration and API client:
//
//	import "github.com/harness/harness-go-sdk/harness/nextgen"
//
//	cfg := nextgen.NewConfiguration()
//	cfg.BasePath = "https://app.harness.io/gateway"
//	cfg.AddDefaultHeader("x-api-key", "your-api-key")
//
//	client := nextgen.NewAPIClient(cfg)
//
// # Authentication
//
// All API calls require authentication via an API key passed in the x-api-key header:
//
//	cfg.AddDefaultHeader("x-api-key", os.Getenv("HARNESS_API_KEY"))
//
// # Available API Services
//
// The client provides access to numerous API services. Here are the most commonly used:
//
// Core Platform:
//   - OrganizationApi - Manage organizations
//   - ProjectApi - Manage projects
//   - ConnectorsApi - Manage connectors (Git, Docker, Cloud providers, etc.)
//   - SecretsApi - Manage secrets and secret managers
//   - UserGroupApi - Manage user groups and permissions
//   - RoleAssignmentsApi - Manage role assignments
//
// Continuous Delivery:
//   - PipelinesApi - Create and manage CD pipelines
//   - ServicesApi - Manage services
//   - EnvironmentsApi - Manage environments
//   - InfrastructureApi - Manage infrastructure definitions
//
// Continuous Integration:
//   - InputSetsApi - Manage pipeline input sets
//   - TriggersApi - Manage pipeline triggers
//
// Feature Flags:
//   - FeatureFlagsApi - Manage feature flags
//   - TargetGroupsApi - Manage target groups
//   - TargetsApi - Manage targets
//
// Cloud Cost Management:
//   - CloudCostBudgetsApi - Manage budgets
//   - CloudCostPerspectivesApi - Manage perspectives
//   - CloudCostAutoStoppingRulesApi - Manage auto-stopping rules
//
// GitOps:
//   - AgentsApi - Manage GitOps agents
//   - ClustersApi - Manage clusters
//   - ApplicationsApi - Manage applications
//   - RepositoriesApi - Manage repositories
//
// # Example: List Organizations
//
//	resp, httpResp, err := client.OrganizationApi.GetOrganizationList(
//	    context.Background(),
//	    accountIdentifier,
//	    &nextgen.OrganizationApiGetOrganizationListOpts{},
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, org := range resp.Data.Content {
//	    fmt.Println(org.Organization.Name)
//	}
//
// # Example: Create a Connector
//
//	connector := nextgen.Connector{
//	    Name:       "my-github-connector",
//	    Identifier: "my_github_connector",
//	    Type_:      "Github",
//	    Spec: map[string]interface{}{
//	        "url":            "https://github.com/myorg",
//	        "authentication": map[string]interface{}{...},
//	    },
//	}
//
//	resp, httpResp, err := client.ConnectorsApi.CreateConnector(
//	    context.Background(),
//	    connector,
//	    accountIdentifier,
//	    nil,
//	)
//
// # Error Handling
//
// API errors are returned as Go errors. The HTTP response is also returned
// for additional context:
//
//	resp, httpResp, err := client.OrganizationApi.GetOrganization(...)
//	if err != nil {
//	    if httpResp != nil {
//	        log.Printf("HTTP Status: %d", httpResp.StatusCode)
//	    }
//	    return err
//	}
//
// # Pagination
//
// List endpoints support pagination via optional parameters:
//
//	opts := &nextgen.OrganizationApiGetOrganizationListOpts{
//	    PageIndex: optional.NewInt32(0),
//	    PageSize:  optional.NewInt32(50),
//	}
//	resp, _, err := client.OrganizationApi.GetOrganizationList(ctx, accountId, opts)
//
// # More Information
//
// For complete API documentation, visit:
// https://apidocs.harness.io/
//
// For SDK issues and contributions:
// https://github.com/harness/harness-go-sdk
package nextgen