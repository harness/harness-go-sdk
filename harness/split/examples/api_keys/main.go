// HSF-40 example: ApiKeys service.
// Demonstrates client usage. Creating/deleting API keys requires appropriate roles and is sensitive.
// This example only resolves workspace and prints usage; use client.ApiKeys.Create/Delete in your own code.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/api_keys <org-id> <project-id>
package main

import (
	"fmt"
	"os"

	"github.com/harness/harness-go-sdk/harness/split"
)

func main() {
	cfg := split.NewConfiguration()
	if cfg.ApiKey == "" {
		fmt.Fprintln(os.Stderr, "HARNESS_PLATFORM_API_KEY is not set")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/api_keys <org-id> <project-id>")
		fmt.Fprintln(os.Stderr, "  Resolves workspace and demonstrates ApiKeys client availability.")
		os.Exit(1)
	}
	orgID := os.Args[1]
	projectID := os.Args[2]

	client := split.NewAPIClient(cfg)
	workspaceID, err := client.Workspaces.ResolveWorkspaceID(orgID, projectID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ResolveWorkspaceID: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Workspace ID: %s\n", workspaceID)
	fmt.Println("ApiKeys service is available on client.ApiKeys")
	fmt.Println("  Create: client.ApiKeys.Create(split.KeyRequest{...})")
	fmt.Println("  Delete: client.ApiKeys.Delete(keyID)")
}
