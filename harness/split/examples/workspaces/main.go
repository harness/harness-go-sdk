// HSF-39 example: Workspaces service (read-only).
// Lists workspaces, FindByName, FindByOrganizationAndProject, ResolveWorkspaceID.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key-or-sat"
//	go run ./harness/split/examples/workspaces [org-id] [project-id]
//	If org-id and project-id are provided, resolves workspace ID and lists; otherwise lists first 10 workspaces.
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

	client := split.NewAPIClient(cfg)

	if len(os.Args) >= 3 {
		orgID := os.Args[1]
		projectID := os.Args[2]
		workspaceID, err := client.Workspaces.ResolveWorkspaceID(orgID, projectID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ResolveWorkspaceID: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("ResolveWorkspaceID(%q, %q) = %s\n", orgID, projectID, workspaceID)
		return
	}

	ws, err := client.Workspaces.List(&split.ListOpts{Limit: 10})
	if err != nil {
		fmt.Fprintf(os.Stderr, "List: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Workspaces (%d):\n", len(ws))
	for _, w := range ws {
		fmt.Printf("  %s: %s (org=%s project=%s)\n", w.ID, w.Name, w.OrganizationIdentifier, w.ProjectIdentifier)
	}
}
