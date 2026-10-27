package code

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/require"
)

func newForkTestClient(t *testing.T, handler http.HandlerFunc) *APIClient {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cfg := NewConfiguration()
	cfg.BasePath = srv.URL
	return NewAPIClient(cfg)
}

func TestRepositoryApiService_ForkCreate(t *testing.T) {
	var (
		gotMethod string
		gotPath   string
		gotQuery  map[string]string
		gotBody   RepoIdentifierForkBody
		gotAPIKey string
	)

	c := newForkTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotQuery = map[string]string{
			"accountIdentifier": r.URL.Query().Get("accountIdentifier"),
			"orgIdentifier":     r.URL.Query().Get("orgIdentifier"),
			"projectIdentifier": r.URL.Query().Get("projectIdentifier"),
		}
		gotAPIKey = r.Header.Get("x-api-key")
		require.NoError(t, json.NewDecoder(r.Body).Decode(&gotBody))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"identifier":"my_fork","fork_id":42,"importing":true,"default_branch":"main"}`))
	})

	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "pat"})
	repo, resp, err := c.RepositoryApi.ForkCreate(ctx, "acct", "upstream",
		&RepositoryApiForkCreateOpts{
			Body: optional.NewInterface(&RepoIdentifierForkBody{
				ParentRef:  "acct/org/proj",
				Identifier: "my_fork",
				ForkBranch: "release",
			}),
			OrgIdentifier:     optional.NewString("org"),
			ProjectIdentifier: optional.NewString("proj"),
		})

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	require.Equal(t, http.MethodPost, gotMethod)
	require.Equal(t, "/repos/upstream/fork", gotPath)
	require.Equal(t, map[string]string{
		"accountIdentifier": "acct",
		"orgIdentifier":     "org",
		"projectIdentifier": "proj",
	}, gotQuery)
	require.Equal(t, "pat", gotAPIKey)
	require.Equal(t, RepoIdentifierForkBody{ParentRef: "acct/org/proj", Identifier: "my_fork", ForkBranch: "release"}, gotBody)

	require.Equal(t, "my_fork", repo.Identifier)
	require.EqualValues(t, 42, repo.ForkId)
	require.True(t, repo.Importing)
}

func TestRepositoryApiService_ForkCreate_OmitsEmptyOptionalFields(t *testing.T) {
	var raw map[string]interface{}

	c := newForkTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, json.NewDecoder(r.Body).Decode(&raw))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"identifier":"f"}`))
	})

	_, _, err := c.RepositoryApi.ForkCreate(context.Background(), "acct", "upstream",
		&RepositoryApiForkCreateOpts{
			Body: optional.NewInterface(&RepoIdentifierForkBody{ParentRef: "acct", Identifier: "f"}),
		})
	require.NoError(t, err)
	require.NotContains(t, raw, "fork_branch", "empty fork_branch must not be sent so all branches are forked")
}

func TestRepositoryApiService_ForkCreate_ErrorKeepsStatusAndBody(t *testing.T) {
	c := newForkTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"Can not fork an empty repository."}`))
	})

	_, resp, err := c.RepositoryApi.ForkCreate(context.Background(), "acct", "upstream",
		&RepositoryApiForkCreateOpts{
			Body: optional.NewInterface(&RepoIdentifierForkBody{ParentRef: "acct", Identifier: "f"}),
		})

	require.Error(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)

	swaggerErr, ok := err.(GenericSwaggerError)
	require.True(t, ok, "error must be a code.GenericSwaggerError so provider helpers can classify it")
	require.Contains(t, string(swaggerErr.Body()), "Can not fork an empty repository.")
}
