package nextgen

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// IAC-8121: both Ansible list endpoints accept a repeated `tags` query param
// (OR logic, exact match), but the generated opts structs did not expose it, so
// no Go caller could filter by tag. The server reads the values as qp["tags"],
// so each tag has to be sent as its own tags= pair.

// newAnsibleListRecorder returns a client pointed at a server that records the
// query params of the request it receives and answers with an empty collection.
func newAnsibleListRecorder(t *testing.T, gotQuery *url.Values) *APIClient {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*gotQuery = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	}))
	t.Cleanup(server.Close)

	return newTestClient(server.URL)
}

func TestAnsibleListInventory_TagsSentAsRepeatedQueryParams(t *testing.T) {
	testCases := map[string]struct {
		opts     *AnsibleApiAnsibleListInventoryOpts
		expected []string
	}{
		"Given multiple tags - then each is sent as its own tags param": {
			opts: &AnsibleApiAnsibleListInventoryOpts{
				Tags: optional.NewInterface([]string{"env:prod", "team:platform"}),
			},
			expected: []string{"env:prod", "team:platform"},
		},
		"Given a single tag - then one tags param is sent": {
			opts: &AnsibleApiAnsibleListInventoryOpts{
				Tags: optional.NewInterface([]string{"env:prod"}),
			},
			expected: []string{"env:prod"},
		},
		"Given unset tags - then no tags param is sent": {
			opts:     &AnsibleApiAnsibleListInventoryOpts{Limit: optional.NewInt64(10)},
			expected: nil,
		},
		"Given nil opts - then no tags param is sent": {
			opts:     nil,
			expected: nil,
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			var gotQuery url.Values
			client := newAnsibleListRecorder(t, &gotQuery)

			_, resp, err := client.AnsibleApi.AnsibleListInventory(context.Background(), "org1", "proj1", "acc1", tc.opts)

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.expected, gotQuery["tags"])
		})
	}
}

func TestAnsibleListPlaybooks_TagsSentAsRepeatedQueryParams(t *testing.T) {
	testCases := map[string]struct {
		opts     *AnsibleApiAnsibleListPlaybooksOpts
		expected []string
	}{
		"Given multiple tags - then each is sent as its own tags param": {
			opts: &AnsibleApiAnsibleListPlaybooksOpts{
				Tags: optional.NewInterface([]string{"env:prod", "team:platform"}),
			},
			expected: []string{"env:prod", "team:platform"},
		},
		"Given a single tag - then one tags param is sent": {
			opts: &AnsibleApiAnsibleListPlaybooksOpts{
				Tags: optional.NewInterface([]string{"env:prod"}),
			},
			expected: []string{"env:prod"},
		},
		"Given unset tags - then no tags param is sent": {
			opts:     &AnsibleApiAnsibleListPlaybooksOpts{Limit: optional.NewInt64(10)},
			expected: nil,
		},
		"Given nil opts - then no tags param is sent": {
			opts:     nil,
			expected: nil,
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			var gotQuery url.Values
			client := newAnsibleListRecorder(t, &gotQuery)

			_, resp, err := client.AnsibleApi.AnsibleListPlaybooks(context.Background(), "org1", "proj1", "acc1", tc.opts)

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.expected, gotQuery["tags"])
		})
	}
}

// The tags filter must not disturb the params that were already supported.
func TestAnsibleListInventory_TagsCoexistWithOtherFilters(t *testing.T) {
	var gotQuery url.Values
	client := newAnsibleListRecorder(t, &gotQuery)

	_, _, err := client.AnsibleApi.AnsibleListInventory(context.Background(), "org1", "proj1", "acc1",
		&AnsibleApiAnsibleListInventoryOpts{
			Limit:          optional.NewInt64(25),
			Page:           optional.NewInt64(2),
			SearchTerm:     optional.NewString("web"),
			Sort:           optional.NewString("name"),
			IncludeDetails: optional.NewBool(true),
			Tags:           optional.NewInterface([]string{"env:prod"}),
		})

	require.NoError(t, err)
	assert.Equal(t, "25", gotQuery.Get("limit"))
	assert.Equal(t, "2", gotQuery.Get("page"))
	assert.Equal(t, "web", gotQuery.Get("searchTerm"))
	assert.Equal(t, "name", gotQuery.Get("sort"))
	assert.Equal(t, "true", gotQuery.Get("includeDetails"))
	assert.Equal(t, []string{"env:prod"}, gotQuery["tags"])
}
