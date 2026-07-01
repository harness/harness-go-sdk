// HSF-30 example: x-api-key auth and account ID.
// Uses NewConfiguration() (HARNESS_ACCOUNT_ID, HARNESS_PLATFORM_API_KEY) and
// demonstrates that the client sends x-api-key and AccountId is available for requests.
//
//	export HARNESS_ACCOUNT_ID="your-account-id"
//	export HARNESS_PLATFORM_API_KEY="your-split-api-key-or-sat"
//	go run ./harness/split/examples/x_api_key_account_id
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/harness/harness-go-sdk/harness/split"
)

func main() {
	cfg := split.NewConfiguration()
	if cfg.ApiKey == "" {
		fmt.Fprintln(os.Stderr, "HARNESS_PLATFORM_API_KEY is not set")
		os.Exit(1)
	}
	if cfg.AccountId == "" {
		fmt.Fprintln(os.Stderr, "HARNESS_ACCOUNT_ID is not set")
		os.Exit(1)
	}

	client := split.NewAPIClient(cfg)

	// GET workspaces with account context (x-api-key and accountId from config)
	u := cfg.BasePath + "/internal/api/v2/workspaces?limit=5"
	if client.AccountId != "" {
		parsed, _ := url.Parse(u)
		q := parsed.Query()
		q.Set("organizationIdentifier", client.AccountId)
		parsed.RawQuery = q.Encode()
		u = parsed.String()
	}
	req, err := http.NewRequest(http.MethodGet, u, nil)
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
	fmt.Printf("AccountID: %s\n", client.AccountId)
	fmt.Printf("Status: %d %s\n", resp.StatusCode, resp.Status)
	fmt.Printf("Body: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
}
