/*
 * Split API client - FlagSets service (API v3).
 * https://docs.split.io/reference/create-flag-set
 */
package split

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Flag set API uses v3 base path (not internal/api/v2).
const flagSetsPath = "/api/v3/flag-sets"

// FlagSet represents a flag set grouping feature flags.
type FlagSet struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Workspace   WorkspaceIDRef `json:"workspace,omitempty"`
}

// WorkspaceIDRef is a minimal workspace reference.
type WorkspaceIDRef struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// FlagSetsService provides access to flag sets (API v3).
type FlagSetsService struct {
	client *APIClient
}

// FlagSetRequest is the body for create.
type FlagSetRequest struct {
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Workspace   *WorkspaceIDRef `json:"workspace,omitempty"`
}

// List returns all flag sets for the given workspace (paginated internally).
func (s *FlagSetsService) List(workspaceID string) ([]FlagSet, error) {
	var all []FlagSet
	after := ""
	for {
		u := s.client.BasePath + flagSetsPath + "?workspace_id=" + workspaceID + "&limit=200"
		if after != "" {
			u += "&after=" + after
		}
		req, err := http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return nil, err
		}
		resp, err := s.client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			return nil, fmt.Errorf("flag sets list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
		}
		var result struct {
			Objects        []FlagSet `json:"objects"`
			NextMarker      *string  `json:"nextMarker"`
			PreviousMarker  *string  `json:"previousMarker"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()
		if len(result.Objects) > 0 {
			all = append(all, result.Objects...)
		}
		if result.NextMarker == nil || *result.NextMarker == "" {
			break
		}
		after = *result.NextMarker
	}
	return all, nil
}

// FindByID returns the flag set by ID.
func (s *FlagSetsService) FindByID(id string) (*FlagSet, error) {
	u := s.client.BasePath + flagSetsPath + "/" + id
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
		return nil, fmt.Errorf("flag set get: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out FlagSet
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// FindByName returns the flag set by name in the workspace (via List).
func (s *FlagSetsService) FindByName(workspaceID, name string) (*FlagSet, error) {
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

// Create creates a flag set.
func (s *FlagSetsService) Create(req FlagSetRequest) (*FlagSet, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + flagSetsPath
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
		return nil, fmt.Errorf("flag set create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out FlagSet
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete deletes a flag set by ID.
func (s *FlagSetsService) Delete(id string) error {
	u := s.client.BasePath + flagSetsPath + "/" + id
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
		return fmt.Errorf("flag set delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}
