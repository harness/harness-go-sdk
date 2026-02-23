# AGENTS.md

## Project Overview

The Harness Go SDK is the official Go SDK for the Harness Platform, providing programmatic access to Harness services including CI/CD pipelines, feature flags, cloud cost management, chaos engineering, and more.

## Build System

- **Build tool**: Go modules + Make
- **Build all**: `go build ./...`
- **Build specific package**: `go build ./harness/<package>`
- **Install dependencies**: `go mod download`
- **Update dependencies**: `go mod tidy`

## Testing

- **Run all tests**: `make test` or `go test ./...`
- **Run tests in folder**: `go test ./harness/<package>/...`
- **Run single test file**: `go test ./harness/<package>/<file>_test.go`
- **Run single test case**: `go test -run TestFunctionName ./harness/<package>/...`
- **Test with verbose output**: `go test -v ./...`
- **Test with coverage**: `go test -cover ./...`

## Linting & Formatting

- **Lint check**: `golangci-lint run`
- **Format code**: `go fmt ./...`
- **Vet code**: `go vet ./...`
- **Fix linting issues**: `golangci-lint run --fix`

Note: The `harness/nextgen` directory is skipped by golangci-lint (see `.golangci.yaml`).

## Git Workflow

- **Branch naming**: No strict format required
- **Commit format**: `feat|fix|techdebt|chore: [JIRA-TICKET]: Description`
  - Examples: `feat: [DBOPS-2020]: create new auth type`
  - Types: `feat`, `fix`, `techdebt`, `chore`, `chore(deps)`
  - Jira projects: DBOPS, IAC, CCM, CHAOS, CDS, IDP, PL, PIPE, etc.
- **PR title format**: Same as commit format
- **Default branch**: `main`

## DOs

- Always run tests before committing
- Run `golangci-lint run` before committing to catch linting issues
- Run `go vet ./...` before committing
- Add tests for all new public functions
- Follow existing code patterns in the codebase
- Use descriptive commit messages with Jira ticket references
- Handle errors explicitly - don't ignore returned errors
- Use meaningful variable and function names

## DON'Ts

- Never force push to main/master
- Never commit secrets, .env files, or credentials
- Never run destructive commands without confirmation
- Never skip go vet checks
- Don't add dependencies without justification
- Don't modify generated code in `harness/nextgen` directory without regenerating

## Commands to Never Run

- `git push --force origin main`
- `git push --force origin master`
- `git commit --no-verify` or `git push --no-verify` (never skip hooks)
- `rm -rf /` or any recursive delete of system directories

## Project Structure

```
harness-go-sdk/
├── harness/           # Main SDK packages
│   ├── cd/            # FirstGen Continuous Delivery APIs
│   ├── chaos/         # Chaos Engineering APIs
│   ├── code/          # Code Repository APIs
│   ├── dbops/         # Database Operations APIs
│   ├── delegate/      # Delegate management
│   ├── har/           # HAR utilities
│   ├── helpers/       # Common helper utilities
│   ├── idp/           # Internal Developer Portal APIs
│   ├── nextgen/       # NextGen Platform APIs (generated)
│   ├── po/            # Platform Operations APIs
│   ├── policymgmt/    # Policy Management APIs
│   ├── svcdiscovery/  # Service Discovery APIs
│   ├── time/          # Time utilities
│   └── utils/         # Utility functions
├── logging/           # Logging utilities
├── tools/             # Development tools
├── docs/              # Documentation
├── .github/           # GitHub configuration
├── .harness/          # Harness pipeline configuration
└── .chglog/           # Changelog configuration
```

## Important Packages

These are the core packages most frequently modified:

| Package | Description |
|---------|-------------|
| `harness/nextgen` | NextGen Platform APIs - most active, contains generated API client code |
| `harness/cd` | FirstGen CD APIs |
| `harness/chaos` | Chaos Engineering integration |
| `harness/helpers` | Shared helper functions |
| `harness/utils` | Common utilities |

## Language-Specific Guidelines

This is a Go codebase. Follow these conventions:

- Use `gofmt` for formatting
- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use error wrapping with `fmt.Errorf("context: %w", err)`
- Export only necessary types and functions
- Keep packages focused and cohesive
- Use table-driven tests where appropriate

## Dependencies

Key dependencies:
- `github.com/hashicorp/go-retryablehttp` - HTTP client with retries
- `github.com/stretchr/testify` - Testing utilities
- `github.com/sirupsen/logrus` - Structured logging
- `golang.org/x/oauth2` - OAuth2 authentication
