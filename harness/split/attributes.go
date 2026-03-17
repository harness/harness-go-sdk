/*
 * Split API client - Attributes service (schema attributes for a traffic type).
 * https://docs.split.io/reference/attributes-overview
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

const schemaPath = "/internal/api/v2/schema/ws"

// Attribute represents an attribute in a traffic type schema.
type Attribute struct {
	ID                 string   `json:"id"`
	OrganizationId      string   `json:"organizationId,omitempty"`
	TrafficTypeID      string   `json:"trafficTypeId,omitempty"`
	DisplayName        string   `json:"displayName,omitempty"`
	Description        string   `json:"description,omitempty"`
	DataType           string   `json:"dataType,omitempty"` // "string", "datetime", "number", "set"
	IsSearchable       bool     `json:"isSearchable"`
	SuggestedValues    []string `json:"suggestedValues,omitempty"`
}

// AttributesService provides access to attributes for a traffic type in a workspace.
type AttributesService struct {
	client *APIClient
}

// AttributeListOptions holds optional query params for listing attributes.
type AttributeListOptions struct {
	Paginate     bool
	SearchPrefix string
	AfterMarker  string
	BeforeMarker string
	MarkerLimit  int
}

// List returns attributes for the given workspace and traffic type.
func (s *AttributesService) List(workspaceID, trafficTypeID string, opts *AttributeListOptions) ([]Attribute, error) {
	path := schemaPath + "/" + workspaceID + "/trafficTypes/" + trafficTypeID
	if opts != nil && (opts.Paginate || opts.SearchPrefix != "" || opts.AfterMarker != "" || opts.BeforeMarker != "" || opts.MarkerLimit > 0) {
		params := url.Values{}
		if opts.Paginate {
			params.Set("paginate", "true")
		}
		if opts.SearchPrefix != "" {
			params.Set("searchPrefix", opts.SearchPrefix)
		}
		if opts.AfterMarker != "" {
			params.Set("afterMarker", opts.AfterMarker)
		}
		if opts.BeforeMarker != "" {
			params.Set("beforeMarker", opts.BeforeMarker)
		}
		if opts.MarkerLimit > 0 {
			params.Set("markerLimit", fmt.Sprintf("%d", opts.MarkerLimit))
		}
		path += "?" + params.Encode()
	}
	u := s.client.BasePath + path
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("attributes list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// API may return either a raw array or an object {"objects": [...], "nextMarker": "..."}.
	var out []Attribute
	if err := json.Unmarshal(body, &out); err == nil {
		return out, nil
	}
	var result struct {
		Objects        []Attribute `json:"objects"`
		NextMarker     *string     `json:"nextMarker"`
		PreviousMarker *string     `json:"previousMarker"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Objects, nil
}

// FindByID returns the attribute with the given ID, or nil if not found.
func (s *AttributesService) FindByID(workspaceID, trafficTypeID, attributeID string, opts *AttributeListOptions) (*Attribute, error) {
	list, err := s.List(workspaceID, trafficTypeID, opts)
	if err != nil {
		return nil, err
	}
	for i := range list {
		if list[i].ID == attributeID {
			return &list[i], nil
		}
	}
	return nil, nil
}

// AttributeRequest is the body for create/update attribute.
type AttributeRequest struct {
	Identifier      string   `json:"id,omitempty"`
	DisplayName     string   `json:"displayName,omitempty"`
	Description     string   `json:"description,omitempty"`
	TrafficTypeID   string   `json:"trafficTypeId,omitempty"`
	DataType        string   `json:"dataType,omitempty"`
	IsSearchable    *bool    `json:"isSearchable,omitempty"`
	SuggestedValues []string `json:"suggestedValues,omitempty"`
}

// Create creates an attribute for the traffic type.
func (s *AttributesService) Create(workspaceID, trafficTypeID string, req AttributeRequest) (*Attribute, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + schemaPath + "/" + workspaceID + "/trafficTypes/" + trafficTypeID
	httpReq, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("attribute create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out Attribute
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update updates an attribute. AttributeID may need to be URL-encoded.
func (s *AttributesService) Update(workspaceID, trafficTypeID, attributeID string, req AttributeRequest) (*Attribute, error) {
	body, _ := json.Marshal(req)
	encodedID := url.QueryEscape(attributeID)
	u := s.client.BasePath + schemaPath + "/" + workspaceID + "/trafficTypes/" + trafficTypeID + "/" + encodedID
	httpReq, err := http.NewRequest(http.MethodPatch, u, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("attribute update: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out Attribute
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete deletes an attribute. AttributeID may need to be URL-encoded.
func (s *AttributesService) Delete(workspaceID, trafficTypeID, attributeID string) error {
	encodedID := url.QueryEscape(attributeID)
	u := s.client.BasePath + schemaPath + "/" + workspaceID + "/trafficTypes/" + trafficTypeID + "/" + encodedID
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("attribute delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}
