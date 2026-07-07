package idp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/require"
)

func TestEntityWriteDeleteIncludeHarnessAccountQueryParam(t *testing.T) {
	const accountID = "account-123"

	tests := []struct {
		name       string
		method     string
		path       string
		invoke     func(*APIClient) error
		statusCode int
	}{
		{
			name:   "UpdateEntity",
			method: http.MethodPut,
			path:   "/v1/entities/account/component/example",
			invoke: func(client *APIClient) error {
				_, _, err := client.EntitiesApi.UpdateEntity(
					context.Background(),
					EntityUpdateRequest{},
					"account",
					"component",
					"example",
					&EntitiesApiUpdateEntityOpts{
						HarnessAccount: optional.NewString(accountID),
					},
				)
				return err
			},
			statusCode: http.StatusOK,
		},
		{
			name:   "DeleteEntity",
			method: http.MethodDelete,
			path:   "/v1/entities/account/component/example",
			invoke: func(client *APIClient) error {
				_, err := client.EntitiesApi.DeleteEntity(
					context.Background(),
					"account",
					"component",
					"example",
					&EntitiesApiDeleteEntityOpts{
						HarnessAccount: optional.NewString(accountID),
					},
				)
				return err
			},
			statusCode: http.StatusNoContent,
		},
		{
			name:   "UpdateEntityVersion",
			method: http.MethodPut,
			path:   "/v1/entities/account/component/example/versions/v1",
			invoke: func(client *APIClient) error {
				_, _, err := client.EntitiesApi.UpdateEntityVersion(
					context.Background(),
					EntityVersionUpdateRequest{},
					"account",
					"component",
					"example",
					"v1",
					&EntitiesApiUpdateEntityVersionOpts{
						HarnessAccount: optional.NewString(accountID),
					},
				)
				return err
			},
			statusCode: http.StatusOK,
		},
		{
			name:   "DeleteEntityVersion",
			method: http.MethodDelete,
			path:   "/v1/entities/account/component/example/versions/v1",
			invoke: func(client *APIClient) error {
				_, err := client.EntitiesApi.DeleteEntityVersion(
					context.Background(),
					"account",
					"component",
					"example",
					"v1",
					&EntitiesApiDeleteEntityVersionOpts{
						HarnessAccount: optional.NewString(accountID),
					},
				)
				return err
			},
			statusCode: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, tt.method, r.Method)
				require.Equal(t, tt.path, r.URL.Path)
				require.Equal(t, accountID, r.Header.Get("Harness-Account"))
				require.Equal(t, accountID, r.URL.Query().Get("accountIdentifier"))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte("{}"))
			}))
			defer server.Close()

			cfg := NewConfiguration()
			cfg.BasePath = server.URL
			client := NewAPIClient(cfg)

			require.NoError(t, tt.invoke(client))
		})
	}
}
