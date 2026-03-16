# Split client – Jira ticket mapping (Phase 1)

Phase 1 tickets (HSF-28 Epic) are 14 SDK tasks. Commit messages and code map as follows:

| Jira   | Summary                          | Implementation |
|--------|----------------------------------|----------------|
| HSF-29 | Package skeleton and HTTP client | configuration, transport, client |
| HSF-30 | x-api-key auth and account ID    | configuration, client, examples/x_api_key_account_id |
| HSF-31 | Environments service             | environments.go, examples/environments |
| HSF-32 | TrafficTypes service             | traffic_types.go, examples/traffic_types |
| HSF-33 | Attributes service               | attributes.go, examples/attributes |
| HSF-34 | FlagSets service                 | flag_sets.go, examples/flag_sets |
| HSF-35 | Segments service                 | segments.go, examples/segments |
| HSF-36/37 | Splits and Split Definitions | splits.go, examples/splits |
| HSF-38 | Tags (used by Splits only)       | SplitTag etc. in splits.go |
| HSF-39 | Workspaces service (read-only)   | workspaces.go, examples/workspaces |
| **HSF-40** | **Rule-based Segments**      | rule_based_segments.go, client.RuleBasedSegments |
| **HSF-41** | **Large Segments**            | large_segments.go, client.LargeSegments |
| **HSF-42** | **Integration test and docs** | split_sdk_test.go, doc.go (Integration tests section) |

**Note:** The commit currently titled “feat: [HSF-40]: Split client: ApiKeys service” is the **ApiKeys** service (create/delete keys). In Jira, **HSF-40** is **Rule-based Segments**. The ApiKeys service has no separate Phase 1 ticket (Phase 2 HSF-57 is the Terraform resource). Rule-based Segments and Large Segments are implemented in this package as above; integration tests and doc live in `split_sdk_test.go` and `doc.go`.
