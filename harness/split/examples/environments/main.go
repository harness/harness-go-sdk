// HSF-31 example: Environments service.
// Resolves workspace ID from org/project, then lists environments.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/environments [org-id] [project-id]
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
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/environments <org-id> <project-id>")
		fmt.Fprintln(os.Stderr, "  Resolves workspace and lists environments for that workspace.")
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

	list, err := client.Environments.List(workspaceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Environments.List: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Environments (%d):\n", len(list))
	for _, e := range list {
		prod := "no"
		if e.Production {
			prod = "yes"
		}
		fmt.Printf("  %s: %s (production=%s, type=%s, status=%s)\n", e.ID, e.Name, prod, e.EnvironmentType, e.Status)
		if cp := e.ChangePermissions; cp != nil {
			if cp.AreApprovalsRequired != nil {
				fmt.Printf("    approvals_required=%v", *cp.AreApprovalsRequired)
				if cp.AreApproversRestricted != nil {
					fmt.Printf("  approvers_restricted=%v", *cp.AreApproversRestricted)
				}
				fmt.Println()
			}
			if len(cp.Approvers) > 0 {
				fmt.Printf("    approvers:")
				for _, a := range cp.Approvers {
					fmt.Printf(" %s(%s/%s)", a.Name, a.Type, a.ID)
				}
				fmt.Println()
			}
			if len(cp.ApprovalSkippableBy) > 0 {
				fmt.Printf("    approval_skippable_by:")
				for _, a := range cp.ApprovalSkippableBy {
					fmt.Printf(" %s(%s/%s)", a.Name, a.Type, a.ID)
				}
				fmt.Println()
			}
			if cp.AreEditorsRestricted != nil && *cp.AreEditorsRestricted {
				fmt.Printf("    editors_restricted=true")
				for _, ed := range cp.Editors {
					fmt.Printf(" %s(%s/%s)", ed.Name, ed.Type, ed.ID)
				}
				fmt.Println()
			}
		}
		if dp := e.DataExportPermissions; dp != nil {
			if dp.AreExportersRestricted != nil && *dp.AreExportersRestricted {
				fmt.Printf("    exporters_restricted=true")
				for _, ex := range dp.Exporters {
					fmt.Printf(" %s(%s/%s)", ex.Name, ex.Type, ex.ID)
				}
				fmt.Println()
			}
		}
	}
}
