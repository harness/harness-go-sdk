// examples/risk_scan/main.go
//
// Creates a chaos risk scan with the same inputs the UI uses.
//
// Infrastructure scan (default). infrastructure id and environment id are the
// same values as harness_chaos_infrastructure_v2 infra_id and environment_id.
// The client reads the discovery agent stored on that infrastructure.
// A lookup failure is returned. The scan is not sent to some other agent.
//
//	export HARNESS_API_KEY=pat.xxxx
//	export INFRA_ID=<infra_id>
//	export ENVIRONMENT_ID=<environment id of that chaos infrastructure>
//	export DETECTION_METHOD=ai
//	go run ./harness/chaos/examples/risk_scan
//
// DETECTION_METHOD selects the UI dropdown. ai is AI Scan. ruleBased is
// Rule-based Scan. AI and RiskAnalyzer are accepted as the same choices.
// Leave it unset and the API uses AI Scan.
//
// Pipeline scan. Use a successful execution that deployed with a Kubernetes CD stage.
//
//	export SCAN_KIND=pipeline
//	export PIPELINE_ID=<pipeline identifier>
//	export PIPELINE_EXECUTION_ID=<planExecutionId>
//	go run ./harness/chaos/examples/risk_scan
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
)

const gateway = "https://app.harness.io/gateway"

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func requireEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("set %s", key)
	}
	return value
}

func explain(step string, resp *http.Response, err error) {
	status := "no response"
	if resp != nil {
		status = resp.Status
	}
	fmt.Printf("%s failed (%s): %v\n", step, status, err)
	var apiErr chaos.GenericSwaggerError
	if errors.As(err, &apiErr) && len(apiErr.Body()) > 0 {
		fmt.Printf("response body: %s\n", apiErr.Body())
	}
}

func printJSON(label string, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("%s: %+v\n", label, v)
		return
	}
	fmt.Printf("%s:\n%s\n", label, b)
}

func main() {
	apiKey := requireEnv("HARNESS_API_KEY")
	scope := chaos.Scope{
		AccountID: requireEnv("HARNESS_ACCOUNT_ID"),
		OrgID:     requireEnv("HARNESS_ORG_ID"),
		ProjectID: requireEnv("HARNESS_PROJECT_ID"),
	}

	cfg := chaos.NewConfiguration()
	cfg.ApiKey = apiKey
	cfg.AccountId = scope.AccountID
	cfg.BasePath = gateway + "/chaos/manager/api"
	cfg.HTTPClient = &http.Client{Timeout: 60 * time.Second}
	api, ctx := chaos.NewAPIClient(cfg).WithAuthContext(context.Background())
	scans := chaos.NewScanClient(api)

	kind := envOr("SCAN_KIND", "infrastructure")
	var scanID string
	var scan *chaos.Scan
	var resp *http.Response
	var err error

	switch kind {
	case "pipeline":
		pipelineID := requireEnv("PIPELINE_ID")
		executionID := requireEnv("PIPELINE_EXECUTION_ID")
		scanID = envOr("SCAN_ID", fmt.Sprintf("%s_sdk_%d", pipelineID, time.Now().Unix()))
		fmt.Printf("=== Creating pipeline scan %s for %s execution %s ===\n", scanID, pipelineID, executionID)
		scan, resp, err = scans.CreatePipelineScan(ctx, scope, chaos.PipelineScanRequest{
			Identity:            scanID,
			Name:                fmt.Sprintf("%s sdk scan", pipelineID),
			Description:         "created from harness-go-sdk examples/risk_scan",
			Tags:                []string{"sdk-test"},
			DetectionMethod:     os.Getenv("DETECTION_METHOD"),
			PipelineIdentity:    pipelineID,
			PipelineExecutionID: executionID,
		})
	default:
		infraID := requireEnv("INFRA_ID")
		environmentID := requireEnv("ENVIRONMENT_ID")
		scanID = envOr("SCAN_ID", fmt.Sprintf("%s_sdk_%d", infraID, time.Now().Unix()))
		fmt.Printf("=== Creating infrastructure scan %s for %s in environment %s ===\n", scanID, infraID, environmentID)
		scan, resp, err = scans.CreateInfrastructureScan(ctx, scope, chaos.InfrastructureScanRequest{
			Identity:         scanID,
			Name:             fmt.Sprintf("%s sdk scan", infraID),
			Description:      "created from harness-go-sdk examples/risk_scan",
			Tags:             []string{"sdk-test"},
			DetectionMethod:  os.Getenv("DETECTION_METHOD"),
			InfrastructureID: infraID,
			EnvironmentID:    environmentID,
		})
	}
	if err != nil {
		explain("create", resp, err)
		os.Exit(1)
	}
	method := scan.DetectionMethod
	if method == "" {
		method = os.Getenv("DETECTION_METHOD")
	}
	if method == "" {
		method = "AI (API default)"
	}
	fmt.Printf("created: identity=%s status=%s scanType=%s detectionMethod=%s\n", scan.Identity, scan.Status, scan.ScanType, method)
	if scan.Source.Infrastructure != nil {
		fmt.Printf("discovery agent=%s environment=%s\n", scan.Source.Infrastructure.AgentIdentity, scan.Source.Infrastructure.EnvironmentIdentity)
	}

	fmt.Println("\n=== Polling scan status ===")
	deadline := time.Now().Add(20 * time.Minute)
	for {
		scan, resp, err = scans.Get(ctx, scope, scanID)
		if err != nil {
			explain("get", resp, err)
			os.Exit(1)
		}
		fmt.Printf("%s status=%s\n", time.Now().Format("15:04:05"), scan.Status)

		switch scan.Status {
		case chaos.ScanStatusCompleted, chaos.ScanStatusErrored, chaos.ScanStatusAborted:
		default:
			if time.Now().After(deadline) {
				log.Fatalf("scan %s still %s after 20 minutes", scanID, scan.Status)
			}
			time.Sleep(15 * time.Second)
			continue
		}
		break
	}

	printJSON("\nfinal scan", scan)
	if scan.Status != chaos.ScanStatusCompleted {
		fmt.Printf("\nscan ended with %s: %s\n", scan.Status, scan.ErrorMessage)
		if scan.CollectedData != nil {
			for _, e := range scan.CollectedData.Errors {
				fmt.Printf("collection error [%s]: %s\n", e.Source, e.Message)
			}
		}
		os.Exit(1)
	}

	fmt.Println("\n=== Fetching report ===")
	report, resp, err := scans.GetReport(ctx, scope, scanID)
	if err != nil {
		explain("report", resp, err)
		os.Exit(1)
	}
	fmt.Printf("risks=%d occurrences=%d bySeverity=%v\n",
		report.Summary.TotalRisks, report.Summary.TotalOccurrences, report.Summary.BySeverity)
	for i, r := range report.ScannedRisks {
		fmt.Printf("%d. [%s] %s (service=%s, occurrences=%d)\n", i+1, r.Severity, r.Name, r.ServiceName, len(r.Occurrences))
	}
}
