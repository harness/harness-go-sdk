// HSF-34 example: FlagSets service.
// Resolves workspace and lists flag sets for that workspace.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/flag_sets <org-id> <project-id>
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
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/flag_sets <org-id> <project-id>")
		fmt.Fprintln(os.Stderr, "  Resolves workspace and lists flag sets.")
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

	list, err := client.FlagSets.List(workspaceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FlagSets.List: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Flag sets (%d):\n", len(list))
	for _, fs := range list {
		fmt.Printf("  %s: %s", fs.ID, fs.Name)
		if fs.Description != "" {
			fmt.Printf(" — %s", fs.Description)
		}
		fmt.Println()
	}
}
