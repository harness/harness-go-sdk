// Segments in an environment: create a workspace segment, activate it in an environment,
// read by name (Segments.Get) and by environment list id (Environments.ListSegmentsAll),
// then deactivate and delete using harness/split only.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/segments_env_crud <org-id> <project-id> [environment-id]
//
// If environment-id is omitted, the first environment in the workspace is used.
//
//	go run ./harness/split/examples/segments_env_crud --workspace <workspace-id> --environment <environment-id>
package main

import (
	"fmt"
	"os"
	"slices"
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

	workspaceID, envID, err := resolveWorkspaceAndEnv(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Workspace ID: %s\n", workspaceID)
	fmt.Printf("Environment ID: %s\n", envID)

	ttList, err := client.TrafficTypes.List(workspaceID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "TrafficTypes.List: %v\n", err)
		os.Exit(1)
	}
	if len(ttList) == 0 {
		fmt.Fprintln(os.Stderr, "workspace has no traffic types; create one in the UI first")
		os.Exit(1)
	}
	trafficTypeID := ttList[0].ID
	fmt.Printf("Using traffic type: %s (%s)\n", trafficTypeID, ttList[0].Name)

	segName := "sdk_seg_env_" + strconv.FormatInt(time.Now().UnixNano(), 10)

	// Create (workspace-level definition)
	fmt.Println("Creating segment…")
	created, err := client.Segments.Create(workspaceID, trafficTypeID, split.SegmentRequest{
		Name:        segName,
		Description: "temporary segment from harness-go-sdk example",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Segments.Create: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created: name=%q\n", created.Name)

	// Confirm findable by name in workspace (GET by name)
	fmt.Println("Checking segment is readable by name after create…")
	byName, err := client.Segments.Get(workspaceID, segName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Segments.Get by name: %v\n", err)
		_ = client.Segments.Delete(workspaceID, segName)
		os.Exit(1)
	}
	if byName == nil || byName.Name != segName {
		fmt.Fprintf(os.Stderr, "Get by name: want name=%q, got %v\n", segName, byName)
		_ = client.Segments.Delete(workspaceID, segName)
		os.Exit(1)
	}
	fmt.Printf("OK: Get by name sees name=%q description=%q\n", byName.Name, byName.Description)

	// Activate in environment
	fmt.Println("Activating segment in environment…")
	if _, err := client.Segments.Activate(envID, segName); err != nil {
		fmt.Fprintf(os.Stderr, "Segments.Activate: %v\n", err)
		_ = client.Segments.Delete(workspaceID, segName)
		os.Exit(1)
	}

	// Read by id: list API returns segment identifiers per environment (often the segment name).
	fmt.Println("Checking segment appears in environment list (read by list id)…")
	ids, err := client.Environments.ListSegmentsAll(workspaceID, envID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Environments.ListSegmentsAll: %v\n", err)
		_ = client.Segments.Deactivate(envID, segName)
		_ = client.Segments.Delete(workspaceID, segName)
		os.Exit(1)
	}
	if !slices.Contains(ids, segName) {
		// Some tenants return an opaque id; surface what we got for debugging.
		fmt.Fprintf(os.Stderr, "environment list does not include segment name %q; ids sample: %v\n", segName, truncateIDs(ids, 10))
		_ = client.Segments.Deactivate(envID, segName)
		_ = client.Segments.Delete(workspaceID, segName)
		os.Exit(1)
	}
	fmt.Printf("OK: environment list contains id/name %q\n", segName)

	// Deactivate (remove from environment)
	fmt.Println("Deactivating segment in environment…")
	if err := client.Segments.Deactivate(envID, segName); err != nil {
		fmt.Fprintf(os.Stderr, "Segments.Deactivate: %v\n", err)
		_ = client.Segments.Delete(workspaceID, segName)
		os.Exit(1)
	}
	idsAfter, err := client.Environments.ListSegmentsAll(workspaceID, envID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListSegmentsAll after deactivate: %v\n", err)
		_ = client.Segments.Delete(workspaceID, segName)
		os.Exit(1)
	}
	if slices.Contains(idsAfter, segName) {
		fmt.Fprintf(os.Stderr, "expected segment %q to be gone from environment after deactivate\n", segName)
		_ = client.Segments.Delete(workspaceID, segName)
		os.Exit(1)
	}
	fmt.Printf("OK: segment no longer listed in environment after deactivate\n")

	// Delete workspace definition
	fmt.Println("Deleting segment…")
	if err := client.Segments.Delete(workspaceID, segName); err != nil {
		fmt.Fprintf(os.Stderr, "Segments.Delete: %v\n", err)
		os.Exit(1)
	}
	_, err = client.Segments.Get(workspaceID, segName)
	if err == nil {
		fmt.Fprintf(os.Stderr, "expected segment %q to be gone after delete\n", segName)
		os.Exit(1)
	}
	fmt.Printf("Delete verified: segment %q is gone (Get returned not found).\n", segName)
	fmt.Println("Done.")
}

func resolveWorkspaceAndEnv(client *split.APIClient) (workspaceID, envID string, err error) {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--workspace" {
		if len(args) != 4 || args[2] != "--environment" {
			fmt.Fprintln(os.Stderr, "With --workspace, use: --workspace <workspace-id> --environment <environment-id>")
			usage()
		}
		workspaceID, envID = args[1], args[3]
		if env, e := client.Environments.FindByID(workspaceID, envID); e != nil {
			return "", "", fmt.Errorf("Environments.FindByID: %w", e)
		} else if env == nil {
			return "", "", fmt.Errorf("environment %q not found in workspace %q", envID, workspaceID)
		}
		return workspaceID, envID, nil
	}
	if len(args) < 2 {
		usage()
	}
	orgID, projectID := args[0], args[1]
	workspaceID, err = client.Workspaces.ResolveWorkspaceID(orgID, projectID)
	if err != nil {
		return "", "", fmt.Errorf("ResolveWorkspaceID: %w", err)
	}
	if len(args) >= 3 {
		envID = args[2]
		if env, e := client.Environments.FindByID(workspaceID, envID); e != nil {
			return "", "", fmt.Errorf("Environments.FindByID: %w", e)
		} else if env == nil {
			return "", "", fmt.Errorf("environment %q not found in workspace %q", envID, workspaceID)
		}
		return workspaceID, envID, nil
	}
	list, err := client.Environments.List(workspaceID)
	if err != nil {
		return "", "", fmt.Errorf("Environments.List: %w", err)
	}
	if len(list) == 0 {
		return "", "", fmt.Errorf("no environments in workspace; create one or pass environment-id")
	}
	envID = list[0].ID
	return workspaceID, envID, nil
}

func truncateIDs(ids []string, max int) []string {
	if len(ids) <= max {
		return ids
	}
	out := make([]string, max)
	copy(out, ids[:max])
	return append(out, fmt.Sprintf("…+%d more", len(ids)-max))
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  go run ./harness/split/examples/segments_env_crud <org-id> <project-id> [environment-id]")
	fmt.Fprintln(os.Stderr, "  go run ./harness/split/examples/segments_env_crud --workspace <workspace-id> --environment <environment-id>")
	os.Exit(1)
}
