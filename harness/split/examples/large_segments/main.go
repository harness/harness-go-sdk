// HSF-41 example: Large Segments service.
// Resolves workspace and lists large segments (and optionally in an environment).
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/large_segments <org-id> <project-id> [environment-id]
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
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/large_segments <org-id> <project-id> [environment-id]")
		fmt.Fprintln(os.Stderr, "  Resolves workspace, lists large segments; if environment-id is set, lists in that environment.")
		os.Exit(1)
	}
	orgID := os.Args[1]
	projectID := os.Args[2]
	envID := ""
	if len(os.Args) >= 4 {
		envID = os.Args[3]
	}

	client := split.NewAPIClient(cfg)
	workspaceID, err := client.Workspaces.ResolveWorkspaceID(orgID, projectID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ResolveWorkspaceID: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Workspace ID: %s\n", workspaceID)

	if envID != "" {
		list, err := client.LargeSegments.ListInEnvironment(workspaceID, envID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "LargeSegments.ListInEnvironment: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Large segments in environment %s (%d):\n", envID, len(list))
		for _, s := range list {
			tt := ""
			if s.TrafficType != nil {
				tt = " (" + s.TrafficType.Name + ")"
			}
			fmt.Printf("  %s%s\n", s.Name, tt)
		}
		return
	}

	list, err := client.LargeSegments.List(workspaceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LargeSegments.List: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Large segments (%d):\n", len(list))
	for _, s := range list {
		tt := ""
		if s.TrafficType != nil {
			tt = " (" + s.TrafficType.Name + ")"
		}
		fmt.Printf("  %s%s\n", s.Name, tt)
	}
	if len(list) > 0 {
		name := list[0].Name
		meta, err := client.LargeSegments.Get(workspaceID, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "LargeSegments.Get(%q): %v\n", name, err)
			os.Exit(1)
		}
		fmt.Printf("First segment %q: id=%s description=%q\n", name, meta.ID, meta.Description)
	}
}
