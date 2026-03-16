// HSF-33 example: Attributes service.
// Resolves workspace, lists traffic types, then lists attributes for the first traffic type.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/attributes <org-id> <project-id>
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
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/attributes <org-id> <project-id>")
		fmt.Fprintln(os.Stderr, "  Resolves workspace, lists traffic types, then lists attributes for the first traffic type.")
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

	ttList, err := client.TrafficTypes.List(workspaceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "TrafficTypes.List: %v\n", err)
		os.Exit(1)
	}
	if len(ttList) == 0 {
		fmt.Println("No traffic types in workspace; nothing to list attributes for.")
		return
	}
	trafficTypeID := ttList[0].ID
	fmt.Printf("Listing attributes for traffic type: %s (%s)\n", ttList[0].Name, trafficTypeID)

	attrs, err := client.Attributes.List(workspaceID, trafficTypeID, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Attributes.List: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Attributes (%d):\n", len(attrs))
	for _, a := range attrs {
		fmt.Printf("  %s: %s (dataType=%s, isSearchable=%v)\n", a.ID, a.DisplayName, a.DataType, a.IsSearchable)
	}
}
