// CRUD validation: runs Create, Read, Update (where supported), Delete on each Split SDK resource type
// to verify the SDK works against the live API.
//
// For JSON request/response body examples, see the Split public API Postman collection:
// https://github.com/splitio/public-api-postman (public Admin APIs.postman_collection.json).
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key"
//	go run ./harness/split/examples/crud_validation <org-id> <project-id>
//
// Optional: RUN_LARGE_SEGMENT_KEY_UPLOAD=1 to exercise large-segment key upload (change-request + CSV).
//
// Exits 0 if all pass, 1 if any fail; prints pass/fail per resource.
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
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run ./harness/split/examples/crud_validation <org-id> <project-id>")
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

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	prefix := "sdk_crud_" + ts + "_"
	// Environment name has a 20-char API limit; use a short unique name (e.g. crud96081e = 11 chars).
	envSuffix := ts
	if len(envSuffix) >= 6 {
		envSuffix = envSuffix[len(envSuffix)-6:]
	}
	envName := "crud" + envSuffix + "e"
	var failed int

	// --- Workspaces (read-only) ---
	if !run("Workspaces", func() error {
		list, err := client.Workspaces.List(nil)
		if err != nil {
			return err
		}
		if len(list) > 0 {
			_, err = client.Workspaces.FindByName(list[0].Name, "IS")
			if err != nil {
				return err
			}
		}
		return nil
	}) {
		failed++
	}

	// --- TrafficTypes ---
	ttName := prefix + "tt"
	var trafficTypeID string
	if !run("TrafficTypes", func() error {
		list, err := client.TrafficTypes.List(workspaceID)
		if err != nil {
			return err
		}
		if len(list) == 0 {
			return fmt.Errorf("workspace has no traffic types; create one in UI first")
		}
		trafficTypeID = list[0].ID
		createdTT, err := client.TrafficTypes.Create(workspaceID, split.CreateRequest{Name: ttName})
		if err != nil {
			return err
		}
		_, err = client.TrafficTypes.FindByName(workspaceID, ttName)
		if err != nil {
			return err
		}
		_, err = client.TrafficTypes.FindByID(workspaceID, trafficTypeID)
		if err != nil {
			return err
		}
		if err := client.TrafficTypes.Delete(workspaceID, createdTT.ID); err != nil {
			return err
		}
		found, err := client.TrafficTypes.FindByName(workspaceID, ttName)
		if err != nil {
			return err
		}
		if found != nil {
			return fmt.Errorf("traffic type %q still exists after delete", ttName)
		}
		return nil
	}) {
		failed++
	}
	if trafficTypeID == "" {
		fmt.Fprintln(os.Stderr, "Skipping remaining tests: no traffic type in workspace")
		os.Exit(1)
	}

	// --- Environments ---
	var envID string
	if !run("Environments", func() error {
		list, err := client.Environments.List(workspaceID)
		if err != nil {
			return err
		}
		if len(list) > 0 {
			envID = list[0].ID // keep first existing env for Segment/RBS/LargeSegment steps
			_, err = client.Environments.FindByID(workspaceID, envID)
			if err != nil {
				return err
			}
		}
		created, err := client.Environments.Create(workspaceID, split.CreateEnvironmentRequest{Name: envName, Production: false})
		if err != nil {
			return err
		}
		_, err = client.Environments.FindByName(workspaceID, envName)
		if err != nil {
			return err
		}
		_, err = client.Environments.Update(workspaceID, created.ID, split.CreateEnvironmentRequest{Name: envName + "_updated", Production: false})
		if err != nil {
			return err
		}
		// Delete auto-created API keys for this environment before deleting the environment.
		for _, tok := range created.ApiTokens {
			_ = client.ApiKeys.Delete(tok.ID)
		}
		if err := client.Environments.Delete(workspaceID, created.ID); err != nil {
			return err
		}
		list, err = client.Environments.List(workspaceID)
		if err != nil {
			return err
		}
		for _, e := range list {
			if e.ID == created.ID {
				return fmt.Errorf("environment %q still exists after delete", created.ID)
			}
		}
		return nil
	}) {
		failed++
	}
	if envID == "" {
		fmt.Fprintln(os.Stderr, "Warning: no environment in workspace; Segment Activate/RBS/LargeSegment env steps may be skipped")
	}

	// --- ApiKeys (create then delete; creates a real key in the account) ---
	if !run("ApiKeys", func() error {
		if envID == "" {
			return fmt.Errorf("skipped: no environment (API key requires at least one environment)")
		}
		keyName := prefix + "apikey"
		created, err := client.ApiKeys.Create(split.KeyRequest{
			Name:         keyName,
			KeyType:      "server_side",
			Workspace:    &split.KeyWorkspaceRequest{Type: "workspace", Id: workspaceID},
			Environments: []split.KeyEnvironmentRequest{{Type: "environment", Id: envID}},
		})
		if err != nil {
			return err
		}
		return client.ApiKeys.Delete(created.Key)
	}) {
		failed++
	}

	// --- Attributes ---
	attrID := prefix + "attr"
	if !run("Attributes", func() error {
		created, err := client.Attributes.Create(workspaceID, trafficTypeID, split.AttributeRequest{
			Identifier:  attrID,
			DisplayName: "CRUD attr",
			DataType:    "string",
		})
		if err != nil {
			return err
		}
		_, err = client.Attributes.List(workspaceID, trafficTypeID, nil)
		if err != nil {
			return err
		}
		_, err = client.Attributes.FindByID(workspaceID, trafficTypeID, created.ID, nil)
		if err != nil {
			return err
		}
		// Exercise query-params branch (Paginate, MarkerLimit).
		_, err = client.Attributes.List(workspaceID, trafficTypeID, &split.AttributeListOptions{Paginate: true, MarkerLimit: 10})
		if err != nil {
			return err
		}
		_, err = client.Attributes.FindByID(workspaceID, trafficTypeID, created.ID, &split.AttributeListOptions{SearchPrefix: "sdk_crud"})
		if err != nil {
			return err
		}
		_, err = client.Attributes.Update(workspaceID, trafficTypeID, created.ID, split.AttributeRequest{
			Identifier:  created.ID,
			DisplayName: "CRUD attr updated",
			DataType:    "string",
		})
		if err != nil {
			return err
		}
		if err := client.Attributes.Delete(workspaceID, trafficTypeID, created.ID); err != nil {
			return err
		}
		foundAttr, err := client.Attributes.FindByID(workspaceID, trafficTypeID, created.ID, nil)
		if err != nil {
			return err
		}
		if foundAttr != nil {
			return fmt.Errorf("attribute %q still exists after delete", created.ID)
		}
		return nil
	}) {
		failed++
	}

	// --- FlagSets ---
	fsName := prefix + "fs"
	if !run("FlagSets", func() error {
		created, err := client.FlagSets.Create(split.FlagSetRequest{
			Name:      fsName,
			Workspace: &split.WorkspaceIDRef{Type: "workspace", ID: workspaceID},
		})
		if err != nil {
			return err
		}
		_, err = client.FlagSets.List(workspaceID)
		if err != nil {
			return err
		}
		_, err = client.FlagSets.FindByName(workspaceID, fsName)
		if err != nil {
			return err
		}
		_, err = client.FlagSets.FindByID(created.ID)
		if err != nil {
			return err
		}
		if err := client.FlagSets.Delete(created.ID); err != nil {
			return err
		}
		_, err = client.FlagSets.FindByID(created.ID)
		if err == nil {
			return fmt.Errorf("flag set %q still exists after delete", created.ID)
		}
		return nil
	}) {
		failed++
	}

	// --- Segments ---
	segName := prefix + "seg"
	if !run("Segments", func() error {
		_, err := client.Segments.Create(workspaceID, trafficTypeID, split.SegmentRequest{Name: segName, Description: "CRUD"})
		if err != nil {
			return err
		}
		_, err = client.Segments.List(workspaceID)
		if err != nil {
			return err
		}
		_, err = client.Segments.Get(workspaceID, segName)
		if err != nil {
			return err
		}
		if envID != "" {
			_, err = client.Segments.Activate(envID, segName)
			if err != nil {
				return err
			}
			_, _, err = client.Environments.ListSegments(workspaceID, envID, 0, 50)
			if err != nil {
				return err
			}
			_, err = client.Environments.ListSegmentsAll(workspaceID, envID)
			if err != nil {
				return err
			}
			err = client.Environments.AddSegmentKeys(workspaceID, envID, segName, []string{"sdk_crud_key"})
			if err != nil {
				return err
			}
			_, _, err = client.Environments.GetSegmentKeys(workspaceID, envID, segName, 0, 50)
			if err != nil {
				return err
			}
			_, err = client.Environments.GetSegmentKeysAll(workspaceID, envID, segName)
			if err != nil {
				return err
			}
			// Exercise replace=true: replace all keys with a single key, then restore empty for Deactivate.
			err = client.Environments.AddSegmentKeysWithReplace(workspaceID, envID, segName, true, []string{"replace_key_1"})
			if err != nil {
				return err
			}
			keysAfterReplace, err := client.Environments.GetSegmentKeysAll(workspaceID, envID, segName)
			if err != nil {
				return err
			}
			if len(keysAfterReplace) != 1 || keysAfterReplace[0] != "replace_key_1" {
				return fmt.Errorf("after replace expected [replace_key_1], got %v", keysAfterReplace)
			}
			err = client.Environments.RemoveSegmentKeys(workspaceID, envID, segName, []string{"replace_key_1"})
			if err != nil {
				return err
			}
			err = client.Segments.Deactivate(envID, segName)
			if err != nil {
				return err
			}
		}
		if err := client.Segments.Delete(workspaceID, segName); err != nil {
			return err
		}
		_, err = client.Segments.Get(workspaceID, segName)
		if err == nil {
			return fmt.Errorf("segment %q still exists after delete", segName)
		}
		return nil
	}) {
		failed++
	}

	// --- Splits ---
	splitName := prefix + "split"
	if !run("Splits", func() error {
		_, err := client.Splits.Create(workspaceID, trafficTypeID, split.SplitCreateRequest{Name: splitName, Description: "CRUD"})
		if err != nil {
			return err
		}
		_, err = client.Splits.List(workspaceID)
		if err != nil {
			return err
		}
		_, err = client.Splits.Get(workspaceID, splitName)
		if err != nil {
			return err
		}
		_, err = client.Splits.UpdateDescription(workspaceID, splitName, "CRUD updated")
		if err != nil {
			return err
		}
		_ = client.Tags.AssociateTags(workspaceID, splitName, split.ObjectTypeSplit, []string{"sdk_crud"})
		// Split definitions: add definition in an environment, read/update/remove it.
		if envID != "" {
			defOff, defHundred := "off", 100
			defReq := split.SplitDefinitionRequest{
				Treatments:        []split.Treatment{{Name: "on"}, {Name: "off"}},
				DefaultRule:       []split.Bucket{{Treatment: &defOff, Size: &defHundred}},
				DefaultTreatment:  "off",
				TrafficAllocation: 100,
			}
			_, err = client.Splits.CreateDefinition(workspaceID, splitName, envID, defReq)
			if err != nil {
				return err
			}
			_, err = client.Splits.GetDefinition(workspaceID, splitName, envID)
			if err != nil {
				return err
			}
			_, err = client.Splits.ListDefinitions(workspaceID, envID)
			if err != nil {
				return err
			}
			defReq.DefaultTreatment = "on"
			_, err = client.Splits.UpdateDefinitionFull(workspaceID, splitName, envID, defReq)
			if err != nil {
				return err
			}
			err = client.Splits.RemoveDefinition(workspaceID, splitName, envID)
			if err != nil {
				return err
			}
		}
		if err := client.Splits.Delete(workspaceID, splitName); err != nil {
			return err
		}
		_, err = client.Splits.Get(workspaceID, splitName)
		if err == nil {
			return fmt.Errorf("split %q still exists after delete", splitName)
		}
		return nil
	}) {
		failed++
	}

	// --- Rule-based Segments ---
	rbsName := prefix + "rbs"
	if !run("RuleBasedSegments", func() error {
		_, err := client.RuleBasedSegments.Create(workspaceID, trafficTypeID, split.RuleBasedSegmentCreateRequest{Name: rbsName})
		if err != nil {
			return err
		}
		_, err = client.RuleBasedSegments.List(workspaceID)
		if err != nil {
			return err
		}
		_, err = client.RuleBasedSegments.Get(workspaceID, rbsName)
		if err != nil {
			return err
		}
		if envID != "" {
			def, err := client.RuleBasedSegments.EnableInEnvironment(envID, rbsName)
			if err != nil {
				return err
			}
			if def != nil {
				def.Title = "CRUD RBS title"
				_, err = client.RuleBasedSegments.UpdateDefinition(workspaceID, envID, rbsName, *def)
				if err != nil {
					return err
				}
			}
			err = client.RuleBasedSegments.DeleteInEnvironment(envID, rbsName)
			if err != nil {
				return err
			}
		}
		if err := client.RuleBasedSegments.Delete(workspaceID, rbsName); err != nil {
			return err
		}
		_, err = client.RuleBasedSegments.Get(workspaceID, rbsName)
		if err == nil {
			return fmt.Errorf("rule-based segment %q still exists after delete", rbsName)
		}
		return nil
	}) {
		failed++
	}

	// --- Large Segments ---
	lsName := prefix + "ls"
	if !run("LargeSegments", func() error {
		_, err := client.LargeSegments.Create(workspaceID, trafficTypeID, split.LargeSegmentCreateRequest{Name: lsName, Description: "CRUD"})
		if err != nil {
			return err
		}
		_, err = client.LargeSegments.List(workspaceID)
		if err != nil {
			return err
		}
		_, err = client.LargeSegments.Get(workspaceID, lsName)
		if err != nil {
			return err
		}
		if envID != "" {
			err = client.LargeSegments.CreateDefinitionForEnvironment(envID, lsName)
			if err != nil {
				return err
			}
			_, err = client.LargeSegments.ListInEnvironment(workspaceID, envID)
			if err != nil {
				return err
			}
			if os.Getenv("RUN_LARGE_SEGMENT_KEY_UPLOAD") == "1" {
				// Exercise change-request + upload path once (do not open a second change request).
				err = client.LargeSegments.UpdateKeysInEnvironment(workspaceID, envID, lsName, []string{"key1"}, nil)
				if err != nil {
					return err
				}
			}
			// When key upload was skipped, segment is empty so DeleteInEnvironment succeeds.
			// When RUN_LARGE_SEGMENT_KEY_UPLOAD=1, a pending change request may block clearing keys (423); try delete anyway.
			err = client.LargeSegments.DeleteInEnvironment(envID, lsName)
			if err != nil {
				if os.Getenv("RUN_LARGE_SEGMENT_KEY_UPLOAD") == "1" {
					fmt.Fprintf(os.Stderr, "  Warning: DeleteInEnvironment after key upload failed (expected if CR pending): %v\n", err)
				} else {
					return err
				}
			}
		}
		if err := client.LargeSegments.Delete(workspaceID, lsName); err != nil {
			return err
		}
		_, err = client.LargeSegments.Get(workspaceID, lsName)
		if err == nil {
			return fmt.Errorf("large segment %q still exists after delete", lsName)
		}
		return nil
	}) {
		failed++
	}

	// --- Negative / error handling ---
	if !runNegative("Get nonexistent segment", func() error {
		_, err := client.Segments.Get(workspaceID, "nonexistent_segment_"+ts)
		if err != nil {
			return nil // expected error
		}
		return fmt.Errorf("expected error from Get nonexistent segment")
	}) {
		failed++
	}
	dupName := prefix + "dup"
	if !runNegative("Create duplicate traffic type", func() error {
		created, err := client.TrafficTypes.Create(workspaceID, split.CreateRequest{Name: dupName})
		if err != nil {
			return err
		}
		defer func() { _ = client.TrafficTypes.Delete(workspaceID, created.ID) }()
		_, err = client.TrafficTypes.Create(workspaceID, split.CreateRequest{Name: dupName})
		if err != nil {
			return nil // expected error (duplicate)
		}
		return fmt.Errorf("expected error when creating duplicate traffic type")
	}) {
		failed++
	}

	if failed > 0 {
		fmt.Fprintf(os.Stderr, "\n%d resource type(s) failed.\n", failed)
		os.Exit(1)
	}
	fmt.Println("\nAll CRUD checks passed.")
}

func run(name string, fn func() error) bool {
	err := fn()
	if err != nil {
		fmt.Printf("  FAIL %s: %v\n", name, err)
		return false
	}
	fmt.Printf("  OK   %s\n", name)
	return true
}

// runNegative runs fn and passes when fn returns nil (expected error was observed).
func runNegative(name string, fn func() error) bool {
	err := fn()
	if err != nil {
		fmt.Printf("  FAIL %s: %v\n", name, err)
		return false
	}
	fmt.Printf("  OK   %s (expected error)\n", name)
	return true
}
