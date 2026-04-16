package split_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentsService_Get(t *testing.T) {
	// The live API returns only core fields (no permissions).
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/internal/api/v2/environments/ws/ws-1/env-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"env-1","name":"Production","production":true}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	env, err := client.Environments.Get("ws-1", "env-1")
	require.NoError(t, err)
	require.NotNil(t, env)
	require.Equal(t, "env-1", env.ID)
	require.Equal(t, "Production", env.Name)
	require.True(t, env.Production)
	require.Nil(t, env.ChangePermissions, "Get does not return permissions")
	require.Nil(t, env.DataExportPermissions, "Get does not return permissions")
}

func TestEnvironmentsService_Get_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"code":404}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	_, err := client.Environments.Get("ws-1", "nonexistent")
	require.Error(t, err)
	require.Contains(t, err.Error(), "404")
}

func TestEnvironmentsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/internal/api/v2/environments/ws/ws-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"env-1","name":"Production","production":true},{"id":"env-2","name":"Staging","production":false}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	list, err := client.Environments.List("ws-1")
	require.NoError(t, err)
	require.Len(t, list, 2)
	require.Equal(t, "env-1", list[0].ID)
	require.Equal(t, "Production", list[0].Name)
	require.True(t, list[0].Production)
	require.Equal(t, "Staging", list[1].Name)
	require.False(t, list[1].Production)
}

func TestEnvironmentsService_FindByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":"env-1","name":"Production","production":true}]`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	env, err := client.Environments.FindByName("ws-1", "Production")
	require.NoError(t, err)
	require.NotNil(t, env)
	require.Equal(t, "env-1", env.ID)
	require.Equal(t, "Production", env.Name)

	env2, err := client.Environments.FindByName("ws-1", "Nonexistent")
	require.NoError(t, err)
	require.Nil(t, env2)
}

func TestEnvironmentsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/internal/api/v2/environments/ws/ws-1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":"env-new","name":"Dev","production":false}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	env, err := client.Environments.Create("ws-1", split.CreateEnvironmentRequest{Name: "Dev", Production: false})
	require.NoError(t, err)
	require.NotNil(t, env)
	require.Equal(t, "env-new", env.ID)
	require.Equal(t, "Dev", env.Name)
	require.False(t, env.Production)
}

func TestEnvironmentsService_ListSegmentsAll(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Contains(t, r.URL.Path, "/segments/ws/ws-1/environments/env-1")
		callCount++
		w.Header().Set("Content-Type", "application/json")
		// Page 1: offset=0, 2 items, total 3
		if callCount == 1 {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"objects":[{"id":"seg-a"},{"id":"seg-b"}],"totalCount":3}`))
			return
		}
		// Page 2: offset=2, 1 item
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"id":"seg-c"}],"totalCount":3}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	all, err := client.Environments.ListSegmentsAll("ws-1", "env-1")
	require.NoError(t, err)
	require.Len(t, all, 3)
	require.Equal(t, []string{"seg-a", "seg-b", "seg-c"}, all)
	require.Equal(t, 2, callCount)
}

func TestEnvironmentsService_ListSegmentsAll_totalZeroFullPageSingleRequest(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var b bytes.Buffer
		b.WriteString(`{"objects":[`)
		for i := 0; i < 50; i++ {
			if i > 0 {
				b.WriteByte(',')
			}
			_, _ = fmt.Fprintf(&b, `{"name":"seg_%d"}`, i)
		}
		b.WriteString(`],"totalCount":0}`)
		_, _ = w.Write(b.Bytes())
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	all, err := client.Environments.ListSegmentsAll("ws-1", "env-1")
	require.NoError(t, err)
	require.Len(t, all, 50)
	require.Equal(t, 1, callCount, "totalCount=0 with full page must not paginate forever")
}

func TestEnvironmentsService_ListSegments_nameOnlyObjects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Contains(t, r.URL.Path, "/segments/ws/ws-1/environments/env-1")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"objects":[{"name":"alpha"},{"name":"beta"}],"totalCount":2}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	ids, total, err := client.Environments.ListSegments("ws-1", "env-1", 0, 50)
	require.NoError(t, err)
	require.Equal(t, 2, total)
	require.Equal(t, []string{"alpha", "beta"}, ids)
}

func TestEnvironmentsService_ListSegments_itemsShape(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"items":[{"id":"id-1","name":"n1"}],"totalCount":1}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	ids, total, err := client.Environments.ListSegments("ws-1", "env-1", 0, 50)
	require.NoError(t, err)
	require.Equal(t, 1, total)
	require.Equal(t, []string{"id-1"}, ids)
}

func TestEnvironmentsService_ListSegments_FMEObjectsWithNestedFields(t *testing.T) {
	// Shape from live FME (name + nested environment/trafficType; no segment id).
	body := `{
  "objects": [
    {
      "name": "qa_accounts",
      "environment": {
        "id": "9367b3a0-c70b-11f0-9243-aa89d0eb095e",
        "name": "Production-FME"
      },
      "trafficType": {
        "id": "7e042930-d79b-11f0-a575-5259b3e62aa3",
        "name": "account"
      },
      "creationTime": 1765572368827
    },
    {
      "name": "qa_user",
      "environment": {
        "id": "9367b3a0-c70b-11f0-9243-aa89d0eb095e",
        "name": "Production-FME"
      },
      "trafficType": {
        "id": "93079240-c70b-11f0-9243-aa89d0eb095e",
        "name": "user"
      },
      "creationTime": 1764950097892
    }
  ],
  "offset": 0,
  "limit": 20,
  "totalCount": 2
}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	ids, total, err := client.Environments.ListSegments("ws-1", "env-1", 0, 20)
	require.NoError(t, err)
	require.Equal(t, 2, total)
	require.Equal(t, []string{"qa_accounts", "qa_user"}, ids)
}

func TestEnvironmentsService_GetSegmentKeysAll(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Contains(t, r.URL.Path, "/segments/env-1/my_seg/keys")
		callCount++
		w.Header().Set("Content-Type", "application/json")
		if callCount == 1 {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"keys":[{"key":"k1"},{"key":"k2"}],"totalCount":3}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"keys":[{"key":"k3"}],"totalCount":3}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	all, err := client.Environments.GetSegmentKeysAll("ws-1", "env-1", "my_seg")
	require.NoError(t, err)
	require.Len(t, all, 3)
	require.Equal(t, []string{"k1", "k2", "k3"}, all)
	require.Equal(t, 2, callCount)
}

func TestEnvironmentsService_List_WithPermissions(t *testing.T) {
	respBody := `[{
		"id":"env-1",
		"type":"environment",
		"name":"Production",
		"production":true,
		"environmentType":"production",
		"orgId":"org-123",
		"status":"ACTIVE",
		"workspaceIds":["ws-1"],
		"creationTime":1658765348027,
		"permissioningEnabled":true,
		"changePermissions":{
			"allowKills":true,
			"areApprovalsRequired":true,
			"areApproversRestricted":true,
			"approvers":[{"id":"u-1","name":"Alice","type":"user"}],
			"areEditorsRestricted":false,
			"editors":[],
			"approvalSkippableBy":[{"id":"g-1","name":"Administrators","type":"group"}]
		},
		"dataExportPermissions":{
			"areExportersRestricted":true,
			"exporters":[{"id":"u-2","name":"Bob","type":"user"}]
		}
	}]`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(respBody))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	list, err := client.Environments.List("ws-1")
	require.NoError(t, err)
	require.Len(t, list, 1)
	env := list[0]
	require.Equal(t, "env-1", env.ID)
	require.Equal(t, "environment", env.Type)
	require.Equal(t, "production", env.EnvironmentType)
	require.Equal(t, "org-123", env.OrgId)
	require.Equal(t, "ACTIVE", env.Status)
	require.Equal(t, []string{"ws-1"}, env.WorkspaceIds)
	require.Equal(t, int64(1658765348027), env.CreationTime)
	require.True(t, env.PermissioningEnabled)

	require.NotNil(t, env.ChangePermissions)
	require.NotNil(t, env.ChangePermissions.AllowKills)
	require.True(t, *env.ChangePermissions.AllowKills)
	require.NotNil(t, env.ChangePermissions.AreApprovalsRequired)
	require.True(t, *env.ChangePermissions.AreApprovalsRequired)
	require.NotNil(t, env.ChangePermissions.AreApproversRestricted)
	require.True(t, *env.ChangePermissions.AreApproversRestricted)
	require.Len(t, env.ChangePermissions.Approvers, 1)
	require.Equal(t, "u-1", env.ChangePermissions.Approvers[0].ID)
	require.Equal(t, "Alice", env.ChangePermissions.Approvers[0].Name)
	require.Equal(t, "user", env.ChangePermissions.Approvers[0].Type)
	require.NotNil(t, env.ChangePermissions.AreEditorsRestricted)
	require.False(t, *env.ChangePermissions.AreEditorsRestricted)
	require.Len(t, env.ChangePermissions.ApprovalSkippableBy, 1)
	require.Equal(t, "g-1", env.ChangePermissions.ApprovalSkippableBy[0].ID)
	require.Equal(t, "Administrators", env.ChangePermissions.ApprovalSkippableBy[0].Name)
	require.Equal(t, "group", env.ChangePermissions.ApprovalSkippableBy[0].Type)

	require.NotNil(t, env.DataExportPermissions)
	require.NotNil(t, env.DataExportPermissions.AreExportersRestricted)
	require.True(t, *env.DataExportPermissions.AreExportersRestricted)
	require.Len(t, env.DataExportPermissions.Exporters, 1)
	require.Equal(t, "u-2", env.DataExportPermissions.Exporters[0].ID)
}

func TestEnvironmentsService_Create_WithApprovals(t *testing.T) {
	var capturedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		capturedBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"env-new","name":"Staging","production":false,"changePermissions":{"areApprovalsRequired":true,"areApproversRestricted":true,"approvers":[{"id":"u-1","name":"Alice","type":"user"}]}}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	boolTrue := true
	env, err := client.Environments.Create("ws-1", split.CreateEnvironmentRequest{
		Name:       "Staging",
		Production: false,
		ChangePermissions: &split.ChangePermissions{
			AreApprovalsRequired:   &boolTrue,
			AreApproversRestricted: &boolTrue,
			Approvers:              []split.PermissionEntity{{ID: "u-1", Name: "Alice", Type: "user"}},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, env)
	require.Equal(t, "env-new", env.ID)
	require.NotNil(t, env.ChangePermissions)
	require.NotNil(t, env.ChangePermissions.AreApprovalsRequired)
	require.True(t, *env.ChangePermissions.AreApprovalsRequired)
	require.Len(t, env.ChangePermissions.Approvers, 1)

	var reqBody map[string]interface{}
	require.NoError(t, json.Unmarshal(capturedBody, &reqBody))
	require.Contains(t, reqBody, "changePermissions")
}

func TestEnvironmentsService_Create_WithApprovalSkippableBy(t *testing.T) {
	var capturedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"env-skip","name":"Dev","production":false}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	boolTrue := true
	_, err := client.Environments.Create("ws-1", split.CreateEnvironmentRequest{
		Name:       "Dev",
		Production: false,
		ChangePermissions: &split.ChangePermissions{
			AreApprovalsRequired:   &boolTrue,
			AreApproversRestricted: &boolTrue,
			ApprovalSkippableBy:    []split.PermissionEntity{{ID: "g-1", Name: "Administrators", Type: "group"}},
		},
	})
	require.NoError(t, err)

	var reqBody map[string]interface{}
	require.NoError(t, json.Unmarshal(capturedBody, &reqBody))
	cp := reqBody["changePermissions"].(map[string]interface{})
	require.Contains(t, cp, "approvalSkippableBy")
	skipBy := cp["approvalSkippableBy"].([]interface{})
	require.Len(t, skipBy, 1)
	require.Equal(t, "Administrators", skipBy[0].(map[string]interface{})["name"])
}

func TestEnvironmentsService_Update_ChangePermissions(t *testing.T) {
	var capturedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		capturedBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"env-1","name":"Production","production":true,"changePermissions":{"areApprovalsRequired":true}}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	boolTrue := true
	boolFalse := false
	env, err := client.Environments.Update("ws-1", "env-1", split.UpdateEnvironmentRequest{
		ChangePermissions: &split.ChangePermissions{
			AreApprovalsRequired:   &boolTrue,
			AreApproversRestricted: &boolFalse,
			Approvers:              []split.PermissionEntity{},
			AreEditorsRestricted:   &boolTrue,
			Editors:                []split.PermissionEntity{{ID: "u-1", Name: "Editor", Type: "user"}},
			ApprovalSkippableBy:    []split.PermissionEntity{{ID: "g-1", Name: "Administrators", Type: "group"}},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, env)

	var patchOps []map[string]interface{}
	require.NoError(t, json.Unmarshal(capturedBody, &patchOps))
	require.Len(t, patchOps, 1)
	require.Equal(t, "replace", patchOps[0]["op"])
	require.Equal(t, "/changePermissions", patchOps[0]["path"])
	val := patchOps[0]["value"].(map[string]interface{})
	require.Equal(t, true, val["areApprovalsRequired"])
}

func TestEnvironmentsService_Update_NameOnly(t *testing.T) {
	var capturedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"env-1","name":"NewName","production":true}`))
	}))
	defer server.Close()

	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = server.URL
	client := split.NewAPIClient(cfg)

	name := "NewName"
	_, err := client.Environments.Update("ws-1", "env-1", split.UpdateEnvironmentRequest{Name: &name})
	require.NoError(t, err)

	var patchOps []map[string]interface{}
	require.NoError(t, json.Unmarshal(capturedBody, &patchOps))
	require.Len(t, patchOps, 1)
	require.Equal(t, "replace", patchOps[0]["op"])
	require.Equal(t, "/name", patchOps[0]["path"])
	require.Equal(t, "NewName", patchOps[0]["value"])
}

func TestEnvironmentsService_Update_NoFields(t *testing.T) {
	cfg := split.NewDefaultConfiguration()
	cfg.BasePath = "http://unused"
	client := split.NewAPIClient(cfg)

	_, err := client.Environments.Update("ws-1", "env-1", split.UpdateEnvironmentRequest{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "no fields to update")
}
