package nextgen

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antihax/optional"
)

func TestFailure_UnmarshalJSON_Code(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{"string code", `{"code":"INVALID_REQUEST","message":"m"}`, "INVALID_REQUEST"},
		{"numeric code", `{"code":400,"message":"m"}`, "400"},
		{"null code", `{"code":null,"message":"m"}`, ""},
		{"missing code", `{"message":"m"}`, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var failure Failure
			if err := json.Unmarshal([]byte(tt.body), &failure); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if failure.Code != tt.want {
				t.Errorf("Code = %q, want %q", failure.Code, tt.want)
			}
			if failure.Message != "m" {
				t.Errorf("Message = %q, want %q", failure.Message, "m")
			}
		})
	}
}

// TestCreateInfrastructure_NumericFailureCode reproduces CDS-135362 / terraform-provider-harness#1300:
// a 400 with Dropwizard's {"code":400,...} body must surface the API message, not the encoding/json error.
func TestCreateInfrastructure_NumericFailureCode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code":400,"message":"Unable to process JSON","details":"the real reason"}`))
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.BasePath = srv.URL
	c := NewAPIClient(cfg)

	_, _, err := c.InfrastructuresApi.CreateInfrastructure(context.Background(), "acc",
		&InfrastructuresApiCreateInfrastructureOpts{Body: optional.NewInterface(InfrastructureRequest{Identifier: "x"})})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	want := "Unable to process JSON"
	if err.Error() != want {
		t.Errorf("Error() = %q, want %q", err.Error(), want)
	}
}
