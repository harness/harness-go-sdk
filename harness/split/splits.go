/*
 * Split API client - Splits and Split Definitions service.
 * https://docs.split.io/reference/splits-overview
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

const splitsPath = "/internal/api/v2/splits/ws"

// Split is a feature flag / toggle.
type Split struct {
	ID                     string             `json:"id"`
	Name                   string             `json:"name"`
	Description            string             `json:"description,omitempty"`
	Tags                   []SplitTag         `json:"tags,omitempty"`
	CreationTime           int64              `json:"creationTime,omitempty"`
	RolloutStatusTimestamp int64              `json:"rolloutStatusTimestamp,omitempty"`
	TrafficType            *TrafficType       `json:"trafficType,omitempty"`
	RolloutStatus          *SplitRolloutStatus `json:"rolloutStatus,omitempty"`
}

// SplitTag is a tag on a split.
type SplitTag struct {
	Name string `json:"name"`
}

// SplitRolloutStatus represents rollout status.
type SplitRolloutStatus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SplitsListResult is the response when listing splits.
type SplitsListResult struct {
	Objects    []Split `json:"objects"`
	TotalCount int     `json:"totalCount"`
}

// SplitsService provides access to splits and split definitions.
type SplitsService struct {
	client *APIClient
}

// SplitCreateRequest is the body for creating a split.
type SplitCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// List returns splits in the workspace (with optional query params).
func (s *SplitsService) List(workspaceID string) (*SplitsListResult, error) {
	u := s.client.BasePath + splitsPath + "/" + workspaceID
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
		return nil, fmt.Errorf("splits list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out SplitsListResult
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get returns a split by ID or name.
func (s *SplitsService) Get(workspaceID, splitID string) (*Split, error) {
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/" + splitID
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
		return nil, fmt.Errorf("split get: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out Split
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create creates a split in the workspace for the given traffic type.
func (s *SplitsService) Create(workspaceID, trafficTypeID string, req SplitCreateRequest) (*Split, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/trafficTypes/" + trafficTypeID
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
		return nil, fmt.Errorf("split create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out Split
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateDescription updates a split's description. Split name is required.
func (s *SplitsService) UpdateDescription(workspaceID, splitName, description string) (*Split, error) {
	encodedName := url.QueryEscape(splitName)
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/" + encodedName + "/updateDescription"
	httpReq, err := http.NewRequest(http.MethodPut, u, bytes.NewReader([]byte(description)))
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
		return nil, fmt.Errorf("split update description: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out Split
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete deletes a split (by name). Unconfigures from all environments.
func (s *SplitsService) Delete(workspaceID, splitName string) error {
	encodedName := url.QueryEscape(splitName)
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/" + encodedName
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
		return fmt.Errorf("split delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}

// SplitDefinition is the configuration of a split in an environment.
// See https://docs.split.io/reference/feature-flag-definition and create/update definition endpoints.
type SplitDefinition struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	DefaultTreatment   string        `json:"defaultTreatment"`
	BaselineTreatment  string        `json:"baselineTreatment,omitempty"`
	Environment        *Environment  `json:"environment,omitempty"`
	TrafficType        *TrafficType  `json:"trafficType,omitempty"`
	Killed             bool          `json:"killed"`
	Treatments         []Treatment   `json:"treatments"`
	Rules              []Rule        `json:"rules"`
	DefaultRule        []Bucket      `json:"defaultRule"`
	TrafficAllocation  int           `json:"trafficAllocation"`
	CreationTime       int           `json:"creationTime,omitempty"`
	LastUpdateTime     int           `json:"lastUpdateTime,omitempty"`
}

// Treatment is a state of a split (e.g. on/off).
type Treatment struct {
	Name           string   `json:"name,omitempty"`
	Configurations string   `json:"configurations,omitempty"`
	Description    string   `json:"description,omitempty"`
	Keys           []string `json:"keys,omitempty"`
	Segments       []string `json:"segments,omitempty"`
}

// Rule is a condition and buckets.
type Rule struct {
	Condition *Condition `json:"condition"`
	Buckets   []Bucket   `json:"buckets"`
}

// Bucket is a treatment and size.
type Bucket struct {
	Treatment *string `json:"treatment"`
	Size      *int    `json:"size"`
}

// Condition has matchers (AND combiner).
type Condition struct {
	Combiner string    `json:"combiner,omitempty"`
	Matchers []Matcher `json:"matchers,omitempty"`
}

// Matcher selects a subset of users.
type Matcher struct {
	Negate    *bool    `json:"negate,omitempty"`
	Type      *string  `json:"type,omitempty"`
	Attribute *string  `json:"attribute,omitempty"`
	String    *string  `json:"string,omitempty"`
	Bool      *bool    `json:"bool,omitempty"`
	Strings   []string `json:"strings,omitempty"`
	Number    *int     `json:"number,omitempty"`
	Date      *int     `json:"date,omitempty"`
	Between   any      `json:"between,omitempty"`
	Depends   any      `json:"depends,omitempty"`
}

// SplitDefinitionRequest is the body for create/update definition.
// Required: treatments, defaultTreatment, defaultRule. Optional: rules, comment, title, trafficAllocation, baselineTreatment.
// See https://docs.split.io/reference/create-feature-flag-definition-in-environment.
type SplitDefinitionRequest struct {
	Treatments         []Treatment `json:"treatments"`
	Rules              []Rule      `json:"rules"`
	DefaultRule        []Bucket    `json:"defaultRule"`
	DefaultTreatment   string      `json:"defaultTreatment"`
	BaselineTreatment  string      `json:"baselineTreatment,omitempty"`
	Comment            string      `json:"comment,omitempty"`
	Title              string      `json:"title,omitempty"`
	TrafficAllocation  int         `json:"trafficAllocation"`
}

// ListDefinitions returns split definitions in an environment.
func (s *SplitsService) ListDefinitions(workspaceID, environmentID string) (*SplitDefinitionsResult, error) {
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/environments/" + environmentID
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
		return nil, fmt.Errorf("split definitions list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out SplitDefinitionsResult
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// SplitDefinitionsResult is the response when listing definitions.
type SplitDefinitionsResult struct {
	Objects    []SplitDefinition `json:"objects"`
	TotalCount int               `json:"totalCount"`
}

// GetDefinition returns a split definition by name and environment.
func (s *SplitsService) GetDefinition(workspaceID, splitName, environmentID string) (*SplitDefinition, error) {
	encodedName := url.QueryEscape(splitName)
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/" + encodedName + "/environments/" + environmentID
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
		return nil, fmt.Errorf("split definition get: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out SplitDefinition
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateDefinition creates a split definition in an environment.
func (s *SplitsService) CreateDefinition(workspaceID, splitName, environmentID string, req SplitDefinitionRequest) (*SplitDefinition, error) {
	body, _ := json.Marshal(req)
	encodedName := url.QueryEscape(splitName)
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/" + encodedName + "/environments/" + environmentID
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
		return nil, fmt.Errorf("split definition create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out SplitDefinition
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateDefinitionFull does a full PUT update of a split definition.
func (s *SplitsService) UpdateDefinitionFull(workspaceID, splitName, environmentID string, req SplitDefinitionRequest) (*SplitDefinition, error) {
	body, _ := json.Marshal(req)
	encodedName := url.QueryEscape(splitName)
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/" + encodedName + "/environments/" + environmentID
	httpReq, err := http.NewRequest(http.MethodPut, u, bytes.NewReader(body))
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
		return nil, fmt.Errorf("split definition update: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out SplitDefinition
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RemoveDefinition removes a split definition from an environment.
func (s *SplitsService) RemoveDefinition(workspaceID, splitName, environmentID string) error {
	encodedName := url.QueryEscape(splitName)
	u := s.client.BasePath + splitsPath + "/" + workspaceID + "/" + encodedName + "/environments/" + environmentID
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
		return fmt.Errorf("split definition remove: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}
