# Harness Go SDK

[![Latest Release](https://img.shields.io/github/v/release/harness/harness-go-sdk)](https://github.com/harness/harness-go-sdk/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/harness/harness-go-sdk.svg)](https://pkg.go.dev/github.com/harness/harness-go-sdk)
[![License](https://img.shields.io/github/license/harness/harness-go-sdk)](LICENSE.md)

The official Go SDK for the [Harness Platform](https://harness.io). This SDK provides programmatic access to Harness services including CI/CD pipelines, feature flags, cloud cost management, chaos engineering, and more.

## Disclaimer

This product is not supported by the Harness Customer support team. If you have any questions please open a [new issue](https://github.com/harness/harness-go-sdk/issues/new) or join our slack [channel](https://harnesscommunity.slack.com/archives/C02G9CUNF1S).

## Installation

```bash
go get github.com/harness/harness-go-sdk
```

To update to the latest version:

```bash
go get -u github.com/harness/harness-go-sdk@latest
```

## Package Overview

The SDK is organized into the following packages:

| Package | Description | Documentation |
|---------|-------------|---------------|
| [`harness/nextgen`](https://pkg.go.dev/github.com/harness/harness-go-sdk/harness/nextgen) | NextGen Platform APIs (pipelines, connectors, secrets, GitOps, etc.) | [API Docs](https://apidocs.harness.io/) |
| [`harness/cd`](https://pkg.go.dev/github.com/harness/harness-go-sdk/harness/cd) | FirstGen Continuous Delivery APIs | [FirstGen Docs](https://developer.harness.io/docs/first-gen/) |
| [`harness/chaos`](https://pkg.go.dev/github.com/harness/harness-go-sdk/harness/chaos) | Chaos Engineering APIs | [Chaos Docs](https://developer.harness.io/docs/chaos-engineering/) |
| [`harness/code`](https://pkg.go.dev/github.com/harness/harness-go-sdk/harness/code) | Code Repository APIs | [Code Docs](https://developer.harness.io/docs/code-repository/) |
| [`harness/helpers`](https://pkg.go.dev/github.com/harness/harness-go-sdk/harness/helpers) | Common helper utilities | - |
| [`harness/utils`](https://pkg.go.dev/github.com/harness/harness-go-sdk/harness/utils) | Utility functions | - |
| [`logging`](https://pkg.go.dev/github.com/harness/harness-go-sdk/logging) | Logging utilities | - |

## Quick Start

### NextGen Platform (Recommended)

Most users should use the `nextgen` package for interacting with the Harness Platform:

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/harness/harness-go-sdk/harness/nextgen"
)

func main() {
    // Create configuration
    cfg := nextgen.NewConfiguration()
    cfg.BasePath = "https://app.harness.io/gateway"
    cfg.AddDefaultHeader("x-api-key", os.Getenv("HARNESS_API_KEY"))

    // Create client
    client := nextgen.NewAPIClient(cfg)

    // List organizations
    accountId := os.Getenv("HARNESS_ACCOUNT_ID")
    resp, _, err := client.OrganizationApi.GetOrganizationList(
        context.Background(),
        accountId,
        &nextgen.OrganizationApiGetOrganizationListOpts{},
    )
    if err != nil {
        panic(err)
    }

    for _, org := range resp.Data.Content {
        fmt.Println(org.Organization.Name)
    }
}
```

### FirstGen (Legacy)

For FirstGen Harness, use the `cd` package:

```go
package main

import (
    "fmt"
    "os"

    "github.com/harness/harness-go-sdk/harness/cd"
)

func main() {
    // Create client
    client, err := cd.NewClient(&cd.Config{
        AccountId: os.Getenv("HARNESS_ACCOUNT_ID"),
        APIKey:    os.Getenv("HARNESS_API_KEY"),
        Endpoint:  "https://app.harness.io",
    })
    if err != nil {
        panic(err)
    }

    // Get application by name
    app, err := client.ApplicationClient.GetApplicationByName("my-app")
    if err != nil {
        panic(err)
    }
    fmt.Println(app.Id, app.Name)
}
```

## Configuration

### Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `HARNESS_ACCOUNT_ID` | Yes | Your Harness account identifier |
| `HARNESS_API_KEY` | Yes | API key for authentication |
| `HARNESS_ENDPOINT` | No | API endpoint (default: `https://app.harness.io`) |
| `HARNESS_BEARER_TOKEN` | No | Bearer token for certain FirstGen APIs (deprecated) |

### Generating an API Key

1. Go to the [Harness Platform](https://app.harness.io/)
2. Navigate to **Account Settings** → **Access Control** → **API Keys**
3. Click **+ API Key** and create a new key
4. Save the token securely - it won't be shown again

## Common Examples

### Create a Connector

```go
connector := nextgen.Connector{
    Name:       "my-github-connector",
    Identifier: "my_github_connector",
    Type_:      "Github",
    Spec: map[string]interface{}{
        "url": "https://github.com/myorg",
        "authentication": map[string]interface{}{
            "type": "Http",
            "spec": map[string]interface{}{
                "type": "UsernameToken",
                "spec": map[string]interface{}{
                    "username":     "myuser",
                    "tokenRef":     "account.github_token",
                },
            },
        },
    },
}

resp, _, err := client.ConnectorsApi.CreateConnector(
    context.Background(),
    connector,
    accountIdentifier,
    nil,
)
```

### Create a Secret

```go
secret := nextgen.Secret{
    Name:       "my-secret",
    Identifier: "my_secret",
    Type_:      "SecretText",
    Spec: map[string]interface{}{
        "valueType":       "Inline",
        "value":           "my-secret-value",
        "secretManagerIdentifier": "harnessSecretManager",
    },
}

resp, _, err := client.SecretsApi.CreateSecret(
    context.Background(),
    secret,
    accountIdentifier,
    nil,
)
```

### List Pipelines

```go
resp, _, err := client.PipelinesApi.GetPipelineList(
    context.Background(),
    accountIdentifier,
    orgIdentifier,
    projectIdentifier,
    &nextgen.PipelinesApiGetPipelineListOpts{
        PageIndex: optional.NewInt32(0),
        PageSize:  optional.NewInt32(20),
    },
)
```

## Error Handling

API errors include the HTTP response for additional context:

```go
resp, httpResp, err := client.OrganizationApi.GetOrganization(
    context.Background(),
    orgIdentifier,
    accountIdentifier,
    nil,
)
if err != nil {
    if httpResp != nil {
        fmt.Printf("HTTP Status: %d\n", httpResp.StatusCode)
        body, _ := io.ReadAll(httpResp.Body)
        fmt.Printf("Response: %s\n", string(body))
    }
    return err
}
```

## Pagination

List endpoints support pagination:

```go
opts := &nextgen.OrganizationApiGetOrganizationListOpts{
    PageIndex: optional.NewInt32(0),
    PageSize:  optional.NewInt32(50),
}
resp, _, err := client.OrganizationApi.GetOrganizationList(ctx, accountId, opts)

// Check for more pages
if resp.Data.TotalPages > 1 {
    // Fetch additional pages...
}
```

## Advanced Configuration

### Custom HTTP Client

```go
import (
    "net/http"
    "time"

    retryablehttp "github.com/hashicorp/go-retryablehttp"
)

// Create a retryable HTTP client
retryClient := retryablehttp.NewClient()
retryClient.RetryMax = 3
retryClient.RetryWaitMin = 1 * time.Second
retryClient.RetryWaitMax = 5 * time.Second

cfg := nextgen.NewConfiguration()
cfg.HTTPClient = retryClient.StandardClient()
```

### Debug Logging

Enable debug logging for troubleshooting:

```go
import "github.com/harness/harness-go-sdk/logging"

logger := logging.NewLogger()
logger.SetLevel(log.DebugLevel)

// Use with FirstGen client
client, _ := cd.NewClient(&cd.Config{
    AccountId:    accountId,
    APIKey:       apiKey,
    DebugLogging: true,
    Logger:       logger,
})
```

## Go Modules

This SDK uses Go modules. To get a specific version:

```bash
go get github.com/harness/harness-go-sdk@v0.7.0
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Resources

- [Harness Developer Hub](https://developer.harness.io/)
- [API Reference](https://apidocs.harness.io/)
- [Harness Community Slack](https://harnesscommunity.slack.com/)

## License

This SDK is licensed under the Apache 2.0 License. See [LICENSE.md](LICENSE.md) for details.
