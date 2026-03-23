// Flag sets CRUD demo: create a flag set, read it by ID and by name, then delete it using the
// harness/split SDK (client.FlagSets), not raw HTTP.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/flag_sets_crud <org-id> <project-id>
//
// Optional second form: pass an explicit workspace ID instead of org + project:
//	go run ./harness/split/examples/flag_sets_crud --workspace <workspace-id>
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/harness/harness-go-sdk/harness/split"
)

func main() {
	cfg := split.NewConfiguration()
	if cfg.ApiKey == "" {
		fmt.Fprintln(os.Stderr, "HARNESS_PLATFORM_API_KEY is not set")
		os.Exit(1)
	}

	client := split.NewAPIClient(cfg)

	var workspaceID string
	switch {
	case len(os.Args) >= 3 && os.Args[1] == "--workspace":
		workspaceID = os.Args[2]
	case len(os.Args) >= 3:
		orgID, projectID := os.Args[1], os.Args[2]
		var err error
		workspaceID, err = client.Workspaces.ResolveWorkspaceID(orgID, projectID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ResolveWorkspaceID: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Workspace ID: %s\n", workspaceID)
	default:
		usage()
	}
	name := "sdk_flag_set_crud_" + strconv.FormatInt(time.Now().UnixNano(), 10)

	// Create
	fmt.Println("Creating flag set…")
	created, err := client.FlagSets.Create(split.FlagSetRequest{
		Name:        name,
		Description: "temporary flag set from harness-go-sdk example",
		Workspace:   &split.WorkspaceIDRef{Type: "workspace", ID: workspaceID},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "FlagSets.Create: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created: id=%s name=%q\n", created.ID, created.Name)

	// Confirm it appears in the workspace by name right after creation
	fmt.Println("Checking flag set is findable by name after create…")
	byName, err := client.FlagSets.FindByName(workspaceID, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FlagSets.FindByName: %v\n", err)
		_ = client.FlagSets.Delete(created.ID)
		os.Exit(1)
	}
	if byName == nil || byName.ID != created.ID {
		fmt.Fprintf(os.Stderr, "FindByName after create: want id=%s name=%q, got %v\n", created.ID, name, byName)
		_ = client.FlagSets.Delete(created.ID)
		os.Exit(1)
	}
	fmt.Printf("OK: FindByName sees id=%s name=%q\n", byName.ID, byName.Name)

	// Read
	fmt.Println("Reading flag set by ID…")
	read, err := client.FlagSets.FindByID(created.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FlagSets.FindByID: %v\n", err)
		_ = client.FlagSets.Delete(created.ID)
		os.Exit(1)
	}
	fmt.Printf("Read: id=%s name=%q description=%q\n", read.ID, read.Name, read.Description)

	// Delete
	fmt.Println("Deleting flag set…")
	if err := client.FlagSets.Delete(created.ID); err != nil {
		fmt.Fprintf(os.Stderr, "FlagSets.Delete: %v\n", err)
		os.Exit(1)
	}

	// Confirm it is gone (API should not return 200 for GET after delete)
	_, err = client.FlagSets.FindByID(created.ID)
	if err == nil {
		fmt.Fprintf(os.Stderr, "expected flag set %q to be gone after delete\n", created.ID)
		os.Exit(1)
	}
	fmt.Printf("Delete verified: flag set %q is gone (FindByID returned not found).\n", created.ID)
	fmt.Println("Done.")
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  go run ./harness/split/examples/flag_sets_crud <org-id> <project-id>")
	fmt.Fprintln(os.Stderr, "  go run ./harness/split/examples/flag_sets_crud --workspace <workspace-id>")
	os.Exit(1)
}
