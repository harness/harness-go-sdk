/*
 * Split API client - Segments service.
 * https://docs.split.io/reference/segments-overview
 */
package split

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const segmentsBasePath = "/internal/api/v2/segments"

// Segment represents a Split segment (workspace-level definition).
type Segment struct {
	Name         string       `json:"name"`
	Description  string       `json:"description,omitempty"`
	Environment  *Environment `json:"environment,omitempty"`
	TrafficType  *TrafficType `json:"trafficType,omitempty"`
	CreationTime int64        `json:"creationTime,omitempty"`
}

// SegmentListResult is the response when listing segments.
type SegmentListResult struct {
	Objects    []Segment `json:"objects"`
	TotalCount int      `json:"totalCount"`
}

// SegmentsService provides access to segments in a workspace.
type SegmentsService struct {
	client *APIClient
}

// SegmentRequest is the body for creating a segment.
type SegmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// List returns all segments in the workspace.
func (s *SegmentsService) List(workspaceID string) (*SegmentListResult, error) {
	u := s.client.BasePath + segmentsBasePath + "/ws/" + workspaceID
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
		return nil, fmt.Errorf("segments list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out SegmentListResult
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get returns a segment by name in the workspace.
func (s *SegmentsService) Get(workspaceID, name string) (*Segment, error) {
	u := s.client.BasePath + segmentsBasePath + "/ws/" + workspaceID + "/" + url.QueryEscape(name)
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
		return nil, fmt.Errorf("segment get: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out Segment
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create creates a segment (does not configure it in any environment).
func (s *SegmentsService) Create(workspaceID, trafficTypeID string, req SegmentRequest) (*Segment, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + segmentsBasePath + "/ws/" + workspaceID + "/trafficTypes/" + trafficTypeID
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
		return nil, fmt.Errorf("segment create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out Segment
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete deletes a segment (unconfigures from all environments).
func (s *SegmentsService) Delete(workspaceID, segmentName string) error {
	u := s.client.BasePath + segmentsBasePath + "/ws/" + workspaceID + "/" + url.QueryEscape(segmentName)
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
		return fmt.Errorf("segment delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}

// Activate activates a segment in an environment.
func (s *SegmentsService) Activate(environmentID, segmentName string) (*Segment, error) {
	u := s.client.BasePath + segmentsBasePath + "/" + environmentID + "/" + url.QueryEscape(segmentName)
	req, err := http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("segment activate: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out Segment
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Deactivate deactivates a segment in an environment.
func (s *SegmentsService) Deactivate(environmentID, segmentName string) error {
	u := s.client.BasePath + segmentsBasePath + "/" + environmentID + "/" + url.QueryEscape(segmentName)
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
		return fmt.Errorf("segment deactivate: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}
