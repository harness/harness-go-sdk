// Simple example that uses the split client to call the Split API (e.g. list workspaces).
//
// Set env vars (same as rest of SDK) then run:
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key-or-sat"
//	go run ./harness/split/examples
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/harness/harness-go-sdk/harness/split"
)

func main() {
	cfg := split.NewConfiguration()
	if cfg.ApiKey == "" {
		fmt.Fprintln(os.Stderr, "HARNESS_PLATFORM_API_KEY is not set")
		os.Exit(1)
	}

	client := split.NewAPIClient(cfg)

	// GET workspaces (Split Admin API: internal/api/v2/workspaces)
	url := cfg.BasePath + "/internal/api/v2/workspaces?limit=5"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "new request: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %d %s\n", resp.StatusCode, resp.Status)
	fmt.Printf("Body: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
}
