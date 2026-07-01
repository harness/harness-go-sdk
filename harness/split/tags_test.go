package split

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTagsService_AssociateTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/tags/ws/ws-1/object/my_flag/objecttype/Split", r.URL.Path)
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := NewAPIClient(cfg)

	err := client.Tags.AssociateTags("ws-1", "my_flag", ObjectTypeSplit, []string{"tagA", "tagB"})
	require.NoError(t, err)
}
