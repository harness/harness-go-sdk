/*
 * Split API client - Rule-based Segments service (HSF-40).
 * https://docs.split.io/reference/rule-based-segments-overview
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

const ruleBasedSegmentsPath = "/internal/api/v2/rule-based-segments"

// RuleBasedSegmentDefinition is the definition of a rule-based segment (rules, excludedKeys, etc.).
type RuleBasedSegmentDefinition struct {
	Title            string                   `json:"title,omitempty"`
	Comment          string                   `json:"comment,omitempty"`
	Rules            []RuleBasedSegmentRule   `json:"rules,omitempty"`
	ExcludedKeys     []string                 `json:"excludedKeys,omitempty"`
	ExcludedSegments []RuleBasedSegmentRef    `json:"excludedSegments,omitempty"`
}

// RuleBasedSegmentRule is a rule with condition and matchers.
type RuleBasedSegmentRule struct {
	Condition *Condition `json:"condition,omitempty"`
}

// RuleBasedSegmentRef references another segment by name and type.
type RuleBasedSegmentRef struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// RuleBasedSegmentsService provides access to rule-based segments (distinct from classic Segments).
type RuleBasedSegmentsService struct {
	client *APIClient
}

// RuleBasedSegmentMetadata is workspace-level metadata (from list/get).
type RuleBasedSegmentMetadata struct {
	Name   string        `json:"name,omitempty"`
	TrafficType *TrafficType `json:"trafficType,omitempty"`
}

// List returns all rule-based segments in the workspace.
func (s *RuleBasedSegmentsService) List(workspaceID string) ([]RuleBasedSegmentMetadata, error) {
	u := s.client.BasePath + ruleBasedSegmentsPath + "/ws/" + workspaceID
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
		return nil, fmt.Errorf("rule-based segments list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out []RuleBasedSegmentMetadata
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns a rule-based segment by name (metadata/definition).
func (s *RuleBasedSegmentsService) Get(workspaceID, name string) (*RuleBasedSegmentDefinition, error) {
	encodedName := url.QueryEscape(name)
	u := s.client.BasePath + ruleBasedSegmentsPath + "/ws/" + workspaceID + "/" + encodedName
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
		return nil, fmt.Errorf("rule-based segment get: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out RuleBasedSegmentDefinition
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateRequest is the body for creating a rule-based segment.
type RuleBasedSegmentCreateRequest struct {
	Name string `json:"name"`
}

// Create creates a rule-based segment in the workspace for the given traffic type.
func (s *RuleBasedSegmentsService) Create(workspaceID, trafficTypeID string, req RuleBasedSegmentCreateRequest) (*RuleBasedSegmentMetadata, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + ruleBasedSegmentsPath + "/ws/" + workspaceID + "/trafficTypes/" + trafficTypeID
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
		return nil, fmt.Errorf("rule-based segment create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out RuleBasedSegmentMetadata
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// EnableInEnvironment creates an empty definition for the rule-based segment in an environment.
func (s *RuleBasedSegmentsService) EnableInEnvironment(environmentID, segmentName string) (*RuleBasedSegmentDefinition, error) {
	encodedName := url.QueryEscape(segmentName)
	u := s.client.BasePath + ruleBasedSegmentsPath + "/" + environmentID + "/" + encodedName
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
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("rule-based segment enable in environment: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out RuleBasedSegmentDefinition
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateDefinition updates the rule-based segment definition in an environment (PUT).
func (s *RuleBasedSegmentsService) UpdateDefinition(environmentID, segmentName string, def RuleBasedSegmentDefinition) (*RuleBasedSegmentDefinition, error) {
	body, _ := json.Marshal(def)
	encodedName := url.QueryEscape(segmentName)
	u := s.client.BasePath + ruleBasedSegmentsPath + "/" + environmentID + "/" + encodedName
	req, err := http.NewRequest(http.MethodPut, u, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("rule-based segment update definition: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out RuleBasedSegmentDefinition
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteInEnvironment removes the rule-based segment definition from an environment.
func (s *RuleBasedSegmentsService) DeleteInEnvironment(environmentID, segmentName string) error {
	encodedName := url.QueryEscape(segmentName)
	u := s.client.BasePath + ruleBasedSegmentsPath + "/" + environmentID + "/" + encodedName
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
		return fmt.Errorf("rule-based segment delete in environment: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}

// Delete removes the rule-based segment from the workspace.
func (s *RuleBasedSegmentsService) Delete(workspaceID, segmentName string) error {
	encodedName := url.QueryEscape(segmentName)
	u := s.client.BasePath + ruleBasedSegmentsPath + "/ws/" + workspaceID + "/" + encodedName
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
		return fmt.Errorf("rule-based segment delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}
