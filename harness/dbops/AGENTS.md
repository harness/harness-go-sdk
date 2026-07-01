# AGENTS.md - Database Operations (dbops)

## Purpose

The `dbops` package provides Go SDK client bindings for the Harness Database Operations (DB DevOps) service. This package contains auto-generated code from OpenAPI/Swagger specifications.

## Code Generation

This package is **auto-generated** using Swagger Codegen. Do NOT manually edit the generated files.

- **Generator**: Swagger Codegen
- **Source**: OpenAPI Spec 3 for the DB Service
- **API version**: 1.0
- **Config files**: `.swagger-codegen-ignore`, `.swagger-codegen/`

## Key Files

| File | Purpose |
|------|---------|
| `client.go` | Main API client configuration and HTTP handling |
| `configuration.go` | Client configuration (base path, headers, HTTP client) |
| `api_database_instance.go` | Database instance CRUD operations |
| `api_database_schema.go` | Database schema management |
| `api_deployed_state.go` | Deployment state tracking |
| `api_migration_state.go` | Migration state management |
| `api_pipeline_step.go` | Pipeline step integration |
| `api_log_ingest.go` | Log ingestion APIs |
| `api_overview.go` | Overview/summary APIs |
| `model_*.go` | Data models for API requests/responses |

## API Services

- `DatabaseInstanceApiService` - Manage database instances
- `DatabaseSchemaApiService` - Manage database schemas
- `DeployedStateApiService` - Track deployed states
- `MigrationStateApiService` - Manage migration states
- `PipelineStepApiService` - Pipeline integration
- `LogIngestApiService` - Ingest logs
- `OverviewApiService` - Get overviews

## Testing

- **Run tests**: `go test ./harness/dbops/...`
- **Test with verbose**: `go test -v ./harness/dbops/...`

## DOs

- Update OpenAPI spec and regenerate when API changes
- Only edit `client_utils.go` for custom utility functions
- Follow generated code patterns when adding utilities
- Test API changes against actual Harness DB service

## DON'Ts

- Do NOT manually edit `api_*.go` files (auto-generated)
- Do NOT manually edit `model_*.go` files (auto-generated)
- Do NOT edit `client.go` or `configuration.go` (auto-generated)
- Do NOT commit `.swagger-codegen/` directory changes unless intentional

## Regenerating Code

To regenerate the client code:

1. Update the OpenAPI specification
2. Run Swagger Codegen with appropriate configuration
3. Review generated changes
4. Run tests to verify functionality

## Related Jira Projects

- DBOPS - Database Operations feature tickets
