// HSF-40 example: Rule-based Segments service.
// Resolves workspace and lists rule-based segments (distinct from classic Segments).
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/rule_based_segments <org-id> <project-id>
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
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/rule_based_segments <org-id> <project-id>")
		fmt.Fprintln(os.Stderr, "  Resolves workspace and lists rule-based segments.")
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

	list, err := client.RuleBasedSegments.List(workspaceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "RuleBasedSegments.List: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Rule-based segments (%d):\n", len(list))
	for _, s := range list {
		tt := ""
		if s.TrafficType != nil {
			tt = " (" + s.TrafficType.Name + ")"
		} else if s.TrafficTypeID != "" {
			tt = fmt.Sprintf(" [trafficTypeId=%s]", s.TrafficTypeID)
		}
		id := s.ID
		if id == "" {
			id = "—"
		}
		fmt.Printf("  %s  id=%s  status=%s%s\n", s.Name, id, s.Status, tt)
	}
	if len(list) > 0 {
		name := list[0].Name
		def, err := client.RuleBasedSegments.Get(workspaceID, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "RuleBasedSegments.Get(%q): %v\n", name, err)
			os.Exit(1)
		}
		fmt.Printf("First segment %q definition: title=%q\n", name, def.Title)
	}
}
