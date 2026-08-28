package agent_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamExecution_FullLifecycle(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "text/event-stream", r.Header.Get("Accept"))
		assert.Equal(t, "test-key", r.Header.Get("x-api-key"))
		assert.Contains(t, r.URL.Path, "/exec-123/stream")
		assert.Equal(t, "acc1", r.URL.Query().Get("accountIdentifier"))

		flusher, ok := w.(http.Flusher)
		require.True(t, ok)

		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "event:snapshot\ndata:{\"planExecutionId\":\"exec-123\",\"status\":\"RUNNING\"}\n\n")
		flusher.Flush()

		fmt.Fprintf(w, "event:node_start\nid:node-1\ndata:{\"nodeExecutionId\":\"node-1\",\"name\":\"Build\",\"status\":\"RUNNING\"}\n\n")
		flusher.Flush()

		fmt.Fprintf(w, "event:node_complete\nid:node-1\ndata:{\"nodeExecutionId\":\"node-1\",\"name\":\"Build\",\"status\":\"SUCCEEDED\"}\n\n")
		flusher.Flush()

		fmt.Fprintf(w, "event:execution_complete\ndata:{\"status\":\"SUCCEEDED\"}\n\n")
		flusher.Flush()
	}))
	defer srv.Close()

	cfg := &nextgen.Configuration{
		BasePath: srv.URL,
		ApiKey:   "test-key",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	events, errc := nextgen.StreamExecution(ctx, cfg, "exec-123", "acc1", "org1", "proj1", nil)

	var received []nextgen.ExecutionEvent
	for ev := range events {
		received = append(received, ev)
	}

	select {
	case err := <-errc:
		require.NoError(t, err)
	default:
	}

	require.Len(t, received, 4)
	assert.Equal(t, "snapshot", received[0].Type)
	assert.Equal(t, "RUNNING", received[0].Status)
	assert.Equal(t, "node_start", received[1].Type)
	assert.Equal(t, "Build", received[1].Name)
	assert.Equal(t, "node_complete", received[2].Type)
	assert.Equal(t, "SUCCEEDED", received[2].Status)
	assert.Equal(t, "execution_complete", received[3].Type)
	assert.True(t, received[3].IsTerminal())
}

func TestStreamExecution_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	cfg := &nextgen.Configuration{BasePath: srv.URL}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	events, errc := nextgen.StreamExecution(ctx, cfg, "exec-123", "acc1", "org1", "proj1", nil)

	var received []nextgen.ExecutionEvent
	for ev := range events {
		received = append(received, ev)
	}

	err := <-errc
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "500")
	assert.Empty(t, received)
}

func TestStreamExecution_ContextCancellation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "event:snapshot\ndata:{\"status\":\"RUNNING\"}\n\n")
		flusher.Flush()

		<-r.Context().Done()
	}))
	defer srv.Close()

	cfg := &nextgen.Configuration{BasePath: srv.URL}
	ctx, cancel := context.WithCancel(context.Background())

	events, _ := nextgen.StreamExecution(ctx, cfg, "exec-123", "acc1", "org1", "proj1", nil)

	ev := <-events
	assert.Equal(t, "snapshot", ev.Type)
	cancel()

	for range events {
	}
}

func TestStreamExecution_LastEventID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "prev-event-42", r.Header.Get("Last-Event-ID"))

		flusher, _ := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "event:execution_complete\ndata:{\"status\":\"SUCCEEDED\"}\n\n")
		flusher.Flush()
	}))
	defer srv.Close()

	cfg := &nextgen.Configuration{BasePath: srv.URL}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	events, _ := nextgen.StreamExecution(ctx, cfg, "exec-123", "acc1", "org1", "proj1", &nextgen.StreamOpts{
		LastEventID: "prev-event-42",
	})

	ev := <-events
	assert.Equal(t, "execution_complete", ev.Type)
	assert.True(t, ev.IsTerminal())
}

func TestStreamExecution_Keepalive(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flusher, _ := w.(http.Flusher)
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)

		// Send a keepalive comment, then a real event
		fmt.Fprintf(w, ":keepalive\n\n")
		flusher.Flush()

		fmt.Fprintf(w, "event:execution_complete\ndata:{\"status\":\"FAILED\"}\n\n")
		flusher.Flush()
	}))
	defer srv.Close()

	cfg := &nextgen.Configuration{BasePath: srv.URL}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	events, _ := nextgen.StreamExecution(ctx, cfg, "exec-123", "acc1", "org1", "proj1", nil)

	ev := <-events
	assert.Equal(t, "execution_complete", ev.Type)
	assert.Equal(t, "FAILED", ev.Status)
}

func TestExecutionEvent_IsTerminal(t *testing.T) {
	tests := []struct {
		event    nextgen.ExecutionEvent
		terminal bool
	}{
		{nextgen.ExecutionEvent{Status: "SUCCEEDED"}, true},
		{nextgen.ExecutionEvent{Status: "FAILED"}, true},
		{nextgen.ExecutionEvent{Status: "ERRORED"}, true},
		{nextgen.ExecutionEvent{Status: "ABORTED"}, true},
		{nextgen.ExecutionEvent{Status: "EXPIRED"}, true},
		{nextgen.ExecutionEvent{Status: "RUNNING"}, false},
		{nextgen.ExecutionEvent{Type: "execution_complete"}, true},
		{nextgen.ExecutionEvent{Type: "error"}, true},
		{nextgen.ExecutionEvent{Type: "node_start", Status: "RUNNING"}, false},
		{nextgen.ExecutionEvent{Type: "node_complete", Status: "SUCCEEDED"}, false},
		{nextgen.ExecutionEvent{Type: "node_complete", Status: "FAILED"}, false},
		{nextgen.ExecutionEvent{Type: "snapshot", Status: "SUCCEEDED"}, true},
		{nextgen.ExecutionEvent{Type: "snapshot", Status: "RUNNING"}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.terminal, tt.event.IsTerminal(), "event=%+v", tt.event)
	}
}
