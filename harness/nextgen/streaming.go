package nextgen

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// ExecutionEvent represents a single SSE event from the execution stream.
type ExecutionEvent struct {
	// Type is the SSE event name (e.g. "snapshot", "node_start", "node_complete", "execution_complete", "error").
	Type string `json:"_event_type"`

	// ID is the SSE event ID (for reconnection via Last-Event-ID).
	ID string `json:"_event_id,omitempty"`

	// Fields from snapshot events
	PlanExecutionID string `json:"planExecutionId,omitempty"`
	Status          string `json:"status,omitempty"`
	StartTs         *int64 `json:"startTs,omitempty"`
	EndTs           *int64 `json:"endTs,omitempty"`

	// Fields from node status events
	NodeExecutionID string `json:"nodeExecutionId,omitempty"`
	Identifier      string `json:"identifier,omitempty"`
	Name            string `json:"name,omitempty"`
	StepType        string `json:"stepType,omitempty"`

	// Error message (present when Type == "error")
	Message string `json:"message,omitempty"`

	// Raw holds the full JSON payload for fields not explicitly modeled.
	Raw json.RawMessage `json:"-"`
}

// IsTerminal returns true if this event signals the entire execution has finished.
// Node-level events (node_start, node_complete) are never terminal — only
// execution_complete, error, or snapshot events with a terminal status qualify.
func (e *ExecutionEvent) IsTerminal() bool {
	if e.Type == "execution_complete" || e.Type == "error" {
		return true
	}
	if e.Type == "node_start" || e.Type == "node_complete" || e.Type == "node_status" {
		return false
	}
	switch e.Status {
	case "SUCCEEDED", "FAILED", "ERRORED", "ABORTED", "EXPIRED":
		return true
	}
	return false
}

// StreamOpts configures the SSE streaming connection.
type StreamOpts struct {
	// LastEventID enables reconnection from a known event ID.
	LastEventID string

	// HTTPClient overrides the default HTTP client for the stream connection.
	// If nil, http.DefaultClient is used (retryablehttp is not suitable for long-lived streams).
	HTTPClient *http.Client
}

// StreamExecution connects to the execution SSE endpoint and returns a channel
// of ExecutionEvents. The channel is closed when the execution reaches a terminal
// state, the context is cancelled, or the connection drops.
//
// The caller must cancel the context to clean up the connection when done.
func StreamExecution(ctx context.Context, cfg *Configuration, planExecutionID string, accountID, orgID, projectID string, opts *StreamOpts) (<-chan ExecutionEvent, <-chan error) {
	events := make(chan ExecutionEvent, 32)
	errc := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errc)

		httpClient := http.DefaultClient
		if opts != nil && opts.HTTPClient != nil {
			httpClient = opts.HTTPClient
		}

		url := fmt.Sprintf("%s/pipeline/api/pipelines/execution/%s/stream?accountIdentifier=%s&orgIdentifier=%s&projectIdentifier=%s",
			strings.TrimSuffix(cfg.BasePath, "/"), planExecutionID, accountID, orgID, projectID)

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			errc <- fmt.Errorf("building SSE request: %w", err)
			return
		}
		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("Cache-Control", "no-cache")
		if cfg.ApiKey != "" {
			req.Header.Set("x-api-key", cfg.ApiKey)
		}
		for k, v := range cfg.DefaultHeader {
			req.Header.Set(k, v)
		}
		if opts != nil && opts.LastEventID != "" {
			req.Header.Set("Last-Event-ID", opts.LastEventID)
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			errc <- fmt.Errorf("connecting to SSE: %w", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errc <- fmt.Errorf("SSE endpoint returned HTTP %d", resp.StatusCode)
			return
		}

		scanner := bufio.NewScanner(resp.Body)
		var eventType, eventID, data string

		for scanner.Scan() {
			if ctx.Err() != nil {
				return
			}
			line := scanner.Text()

			if line == "" {
				if data != "" {
					event := parseSSEData(eventType, eventID, data)
					select {
					case events <- event:
					case <-ctx.Done():
						return
					}
					if event.IsTerminal() {
						return
					}
				}
				eventType, eventID, data = "", "", ""
				continue
			}

			if strings.HasPrefix(line, "event:") {
				eventType = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
			} else if strings.HasPrefix(line, "id:") {
				eventID = strings.TrimSpace(strings.TrimPrefix(line, "id:"))
			} else if strings.HasPrefix(line, "data:") {
				data += strings.TrimPrefix(line, "data:")
			} else if strings.HasPrefix(line, ":") {
				// Comment / keepalive, ignore
			}
		}

		if err := scanner.Err(); err != nil && ctx.Err() == nil {
			errc <- fmt.Errorf("reading SSE stream: %w", err)
		}
	}()

	return events, errc
}

func parseSSEData(eventType, eventID, data string) ExecutionEvent {
	var event ExecutionEvent
	raw := []byte(strings.TrimSpace(data))
	event.Raw = raw
	_ = json.Unmarshal(raw, &event)
	event.Type = eventType
	event.ID = eventID
	return event
}

// WaitForExecution polls the execution status endpoint until the execution reaches
// a terminal state. Use this as a fallback when SSE streaming is not available.
//
// pollInterval controls how often to check; defaults to 5s if zero.
func WaitForExecution(ctx context.Context, client *APIClient, accountID, orgID, projectID, planExecutionID string, pollInterval time.Duration) (*PipelineExecutionSummary, error) {
	if pollInterval == 0 {
		pollInterval = 5 * time.Second
	}

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resp, httpResp, err := client.ExecutionDetailsApi.GetExecutionDetailV2(ctx, accountID, orgID, projectID, planExecutionID, nil)
			if err != nil {
				return nil, fmt.Errorf("polling execution status: %w", err)
			}
			if httpResp != nil {
				httpResp.Body.Close()
			}
			if resp.Data == nil || resp.Data.PipelineExecutionSummary == nil {
				continue
			}
			summary := resp.Data.PipelineExecutionSummary
			if isTerminalStatus(summary.Status) {
				return summary, nil
			}
		}
	}
}

func isTerminalStatus(status string) bool {
	switch status {
	case "Success", "Failed", "Errored", "Aborted", "Expired",
		"SUCCEEDED", "FAILED", "ERRORED", "ABORTED", "EXPIRED":
		return true
	}
	return false
}
