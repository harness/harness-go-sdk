/*
 * Split API client - TrafficTypes service.
 * GET https://api.split.io/internal/api/v2/trafficTypes/ws/{workspace-id}
 */
package split

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	trafficTypesPath    = "/internal/api/v2/trafficTypes/ws"
	trafficTypesPathID  = "/internal/api/v2/trafficTypes" // Delete uses ID only (no workspace in path)
)

// TrafficType represents a Split traffic type.
type TrafficType struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	DisplayAttributeID string `json:"displayAttributeId,omitempty"`
}

// TrafficTypesService provides access to traffic types in a workspace.
type TrafficTypesService struct {
	client *APIClient
}

// List returns all traffic types for the given workspace ID.
func (s *TrafficTypesService) List(workspaceID string) ([]TrafficType, error) {
	u := s.client.BasePath + trafficTypesPath + "/" + workspaceID
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
		return nil, fmt.Errorf("traffic types list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out []TrafficType
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// FindByID returns the traffic type with the given ID in the workspace, or nil if not found.
func (s *TrafficTypesService) FindByID(workspaceID, id string) (*TrafficType, error) {
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

// FindByName returns the traffic type with the given name in the workspace, or nil if not found.
func (s *TrafficTypesService) FindByName(workspaceID, name string) (*TrafficType, error) {
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

// CreateRequest is the body for creating a traffic type.
type CreateRequest struct {
	Name               string `json:"name"`
	DisplayAttributeID string `json:"displayAttributeId,omitempty"`
}

// Create creates a traffic type in the workspace. Uses POST to the Split API.
func (s *TrafficTypesService) Create(workspaceID string, req CreateRequest) (*TrafficType, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + trafficTypesPath + "/" + workspaceID
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
		return nil, fmt.Errorf("traffic types create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out TrafficType
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete deletes the traffic type by ID. API uses DELETE /trafficTypes/{id} (no workspace in path).
func (s *TrafficTypesService) Delete(workspaceID, trafficTypeID string) error {
	u := s.client.BasePath + trafficTypesPathID + "/" + trafficTypeID
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
		return fmt.Errorf("traffic types delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}
