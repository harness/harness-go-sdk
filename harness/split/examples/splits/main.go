// HSF-36/37 example: Splits and Split Definitions service.
// Resolves workspace and lists splits; then lists definitions for the first environment.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/splits <org-id> <project-id>
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
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/splits <org-id> <project-id>")
		fmt.Fprintln(os.Stderr, "  Resolves workspace, lists splits, then lists split definitions for first environment.")
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

	splitsResult, err := client.Splits.List(workspaceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Splits.List: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Splits (%d total):\n", splitsResult.TotalCount)
	for _, s := range splitsResult.Objects {
		fmt.Printf("  %s: %s", s.ID, s.Name)
		if s.Description != "" {
			fmt.Printf(" — %s", s.Description)
		}
		fmt.Println()
	}

	envs, err := client.Environments.List(workspaceID)
	if err != nil || len(envs) == 0 {
		fmt.Println("(Skipping definitions: no environments or list failed)")
		return
	}
	envID := envs[0].ID
	fmt.Printf("\nSplit definitions in environment %s (%s):\n", envs[0].Name, envID)
	defs, err := client.Splits.ListDefinitions(workspaceID, envID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Splits.ListDefinitions: %v\n", err)
		return
	}
	for _, d := range defs.Objects {
		fmt.Printf("  %s: defaultTreatment=%s, trafficAllocation=%d\n", d.Name, d.DefaultTreatment, d.TrafficAllocation)
	}
}
