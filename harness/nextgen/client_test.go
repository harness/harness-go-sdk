package nextgen

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
)

var configureClient sync.Once
var client *APIClient

func getClientWithContext() (*APIClient, context.Context) {
	configureClient.Do(func() {
		cfg := NewConfiguration()
		client = NewAPIClient(cfg)
	})

	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: client.ApiKey})
	return client, ctx
}

// TestGenericSwaggerError_Error_ResponseMessagesOnly asserts that Error()
// surfaces the message from the responseMessages envelope when the API
// returns a body that has no top-level "message" field (e.g. the response
// shape from POST /authz/api/roleassignments). Regression test for PL-72255.
func TestGenericSwaggerError_Error_ResponseMessagesOnly(t *testing.T) {
	body := []byte(`{
		"metaData": null,
		"resource": null,
		"responseMessages": [
			{"code": "INVALID_ARGUMENT", "level": "ERROR", "message": "identifier may not be empty"}
		]
	}`)

	var failure Failure
	if err := json.Unmarshal(body, &failure); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	err := GenericSwaggerError{
		body:  body,
		error: "400 Bad Request",
		model: failure,
	}

	got := err.Error()
	want := "identifier may not be empty"
	if got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

// TestGenericSwaggerError_Error_TopLevelMessageWins asserts that when the API
// returns a body with both a top-level "message" and a "responseMessages"
// array, the top-level message continues to be returned (backward compat
// for endpoints that already work today, e.g. /authz/api/roles).
func TestGenericSwaggerError_Error_TopLevelMessageWins(t *testing.T) {
	body := []byte(`{
		"status": "ERROR",
		"code": "INVALID_ARGUMENT",
		"message": "top-level message",
		"correlationId": "abc-123",
		"responseMessages": [
			{"code": "INVALID_ARGUMENT", "level": "ERROR", "message": "envelope message"}
		]
	}`)

	var failure Failure
	if err := json.Unmarshal(body, &failure); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	err := GenericSwaggerError{
		body:  body,
		error: "400 Bad Request",
		model: failure,
	}

	got := err.Error()
	want := "top-level message"
	if got != want {
		t.Errorf("Error() = %q, want %q (top-level message should win)", got, want)
	}
}
