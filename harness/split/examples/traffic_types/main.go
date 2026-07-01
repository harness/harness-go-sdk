// HSF-32 example: TrafficTypes service.
// Resolves workspace ID from org/project (if provided), then lists traffic types.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key-or-sat"
//	go run ./harness/split/examples/traffic_types [org-id] [project-id]
//	If org-id and project-id are provided, resolves workspace and lists traffic types; otherwise prints usage.
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
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/traffic_types <org-id> <project-id>")
		fmt.Fprintln(os.Stderr, "  Resolves workspace and lists traffic types for that workspace.")
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

	list, err := client.TrafficTypes.List(workspaceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "TrafficTypes.List: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Traffic types (%d):\n", len(list))
	for _, tt := range list {
		fmt.Printf("  %s: %s (displayAttributeId=%s)\n", tt.ID, tt.Name, tt.DisplayAttributeID)
	}
}
