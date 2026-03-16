/*
 * Split API client - Environments service.
 * https://docs.split.io/reference/environments-overview
 */
package split

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const environmentsPath = "/internal/api/v2/environments/ws"

// Environment represents a Split environment in a workspace.
type Environment struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Production bool   `json:"production"`
}

// EnvironmentsService provides access to environments in a workspace.
type EnvironmentsService struct {
	client *APIClient
}

// List returns all environments for the given workspace ID.
func (s *EnvironmentsService) List(workspaceID string) ([]Environment, error) {
	u := s.client.BasePath + environmentsPath + "/" + workspaceID
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("environments list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out []Environment
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// FindByID returns the environment with the given ID in the workspace, or nil if not found.
func (s *EnvironmentsService) FindByID(workspaceID, id string) (*Environment, error) {
	list, err := s.List(workspaceID)
	if err != nil {
		return nil, err
	}
	for i := range list {
		if list[i].ID == id {
			return &list[i], nil
		}
	}
	return nil, nil
}

// FindByName returns the environment with the given name in the workspace, or nil if not found.
func (s *EnvironmentsService) FindByName(workspaceID, name string) (*Environment, error) {
	list, err := s.List(workspaceID)
	if err != nil {
		return nil, err
	}
	for i := range list {
		if list[i].Name == name {
			return &list[i], nil
		}
	}
	return nil, nil
}

// CreateEnvironmentRequest is the body for creating an environment.
type CreateEnvironmentRequest struct {
	Name       string `json:"name"`
	Production bool   `json:"production"`
}

// Create creates an environment in the workspace.
func (s *EnvironmentsService) Create(workspaceID string, req CreateEnvironmentRequest) (*Environment, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + environmentsPath + "/" + workspaceID
	httpReq, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("environments create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out Environment
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// envPatchOp is one JSON Patch operation (RFC 6902) for environment update.
type envPatchOp struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

// Update updates an environment (e.g. name or production flag). Uses PATCH with JSON patch body per Split API.
func (s *EnvironmentsService) Update(workspaceID, environmentID string, req CreateEnvironmentRequest) (*Environment, error) {
	var patch []envPatchOp
	patch = append(patch, envPatchOp{Op: "replace", Path: "/name", Value: req.Name})
	patch = append(patch, envPatchOp{Op: "replace", Path: "/production", Value: strconv.FormatBool(req.Production)})
	body, _ := json.Marshal(patch)
	u := s.client.BasePath + environmentsPath + "/" + workspaceID + "/" + environmentID
	httpReq, err := http.NewRequest(http.MethodPatch, u, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("environments update: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out Environment
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete deletes the environment in the workspace.
func (s *EnvironmentsService) Delete(workspaceID, environmentID string) error {
	u := s.client.BasePath + environmentsPath + "/" + workspaceID + "/" + environmentID
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotFound {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("environments delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}

const segmentsPath = "/internal/api/v2/segments"

// ListSegments returns segment IDs/names for the environment. Uses GET /segments/ws/{workspaceID}/environments/{environmentID}.
func (s *EnvironmentsService) ListSegments(workspaceID, environmentID string, offset, limit int) ([]string, int, error) {
	if limit <= 0 {
		limit = 50
	}
	u := s.client.BasePath + segmentsPath + "/ws/" + workspaceID + "/environments/" + environmentID + "?offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, 0, fmt.Errorf("list segments: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out struct {
		Objects   []struct{ ID string `json:"id"` } `json:"objects"`
		TotalCount int `json:"totalCount"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, 0, err
	}
	ids := make([]string, 0, len(out.Objects))
	for _, o := range out.Objects {
		ids = append(ids, o.ID)
	}
	return ids, out.TotalCount, nil
}

// GetSegmentKeys returns keys for a segment in an environment (paginated). Uses GET /segments/{environmentID}/{segmentName}/keys.
func (s *EnvironmentsService) GetSegmentKeys(workspaceID, environmentID, segmentName string, offset, limit int) ([]string, int, error) {
	if limit <= 0 {
		limit = 50
	}
	u := s.client.BasePath + segmentsPath + "/" + environmentID + "/" + segmentName + "/keys?offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, 0, fmt.Errorf("get segment keys: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	// API may return keys as []string or []{ "key": "..." }; support object form per terraform-provider-split.
	var out struct {
		Keys       []struct{ Key string `json:"key"` } `json:"keys"`
		TotalCount int `json:"totalCount"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, 0, err
	}
	keys := make([]string, 0, len(out.Keys))
	for _, k := range out.Keys {
		keys = append(keys, k.Key)
	}
	return keys, out.TotalCount, nil
}

// ListSegmentsAll returns all segment IDs in the environment by fetching every page (offset/limit).
func (s *EnvironmentsService) ListSegmentsAll(workspaceID, environmentID string) ([]string, error) {
	const pageSize = 50
	var all []string
	offset := 0
	for {
		ids, total, err := s.ListSegments(workspaceID, environmentID, offset, pageSize)
		if err != nil {
			return nil, err
		}
		all = append(all, ids...)
		if len(ids) == 0 || offset+len(ids) >= total {
			break
		}
		offset += len(ids)
	}
	return all, nil
}

// GetSegmentKeysAll returns all keys for a segment in an environment by fetching every page (offset/limit).
func (s *EnvironmentsService) GetSegmentKeysAll(workspaceID, environmentID, segmentName string) ([]string, error) {
	const pageSize = 50
	var all []string
	offset := 0
	for {
		keys, total, err := s.GetSegmentKeys(workspaceID, environmentID, segmentName, offset, pageSize)
		if err != nil {
			return nil, err
		}
		all = append(all, keys...)
		if len(keys) == 0 || offset+len(keys) >= total {
			break
		}
		offset += len(keys)
	}
	return all, nil
}

// AddSegmentKeysRequest is the body for add/remove segment keys (comment/title optional).
type AddSegmentKeysRequest struct {
	Keys   []string `json:"keys"`
	Comment string `json:"comment,omitempty"`
	Title  string `json:"title,omitempty"`
}

// AddSegmentKeys adds keys to a segment in an environment. Uses PUT /segments/{environmentID}/{segmentName}/uploadKeys?replace={bool}.
func (s *EnvironmentsService) AddSegmentKeys(workspaceID, environmentID, segmentName string, keys []string) error {
	return s.AddSegmentKeysWithReplace(workspaceID, environmentID, segmentName, false, keys)
}

// AddSegmentKeysWithReplace adds keys to a segment, with replace=true to replace all keys.
func (s *EnvironmentsService) AddSegmentKeysWithReplace(workspaceID, environmentID, segmentName string, replace bool, keys []string) error {
	body, _ := json.Marshal(AddSegmentKeysRequest{Keys: keys})
	replaceVal := "false"
	if replace {
		replaceVal = "true"
	}
	u := s.client.BasePath + segmentsPath + "/" + environmentID + "/" + segmentName + "/uploadKeys?replace=" + replaceVal
	req, err := http.NewRequest(http.MethodPut, u, bytes.NewReader(body))
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("add segment keys: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	return nil
}

// RemoveSegmentKeys removes keys from a segment in an environment. Uses PUT /segments/{environmentID}/{segmentName}/removeKeys.
func (s *EnvironmentsService) RemoveSegmentKeys(workspaceID, environmentID, segmentName string, keys []string) error {
	body, _ := json.Marshal(AddSegmentKeysRequest{Keys: keys})
	u := s.client.BasePath + segmentsPath + "/" + environmentID + "/" + segmentName + "/removeKeys"
	req, err := http.NewRequest(http.MethodPut, u, bytes.NewReader(body))
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("remove segment keys: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	return nil
}
