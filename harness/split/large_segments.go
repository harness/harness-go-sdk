/*
 * Split API client - Large Segments service (HSF-41).
 * https://docs.split.io/reference/ (Large Segments)
 * Update keys: https://docs.split.io/reference/create-change-request#open-change-request-to-add-members-to-a-large-segment
 */
package split

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const largeSegmentsPath = "/internal/api/v2/large-segments"

// LargeSegmentMetadata is workspace-level large segment metadata.
type LargeSegmentMetadata struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description,omitempty"`
	TrafficType  *TrafficType `json:"trafficType,omitempty"`
	CreationTime int64        `json:"creationTime,omitempty"`
	Tags         []TagRef     `json:"tags,omitempty"`
}

// TagRef is a URN reference (id, type, name).
type TagRef struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

// LargeSegmentsService provides access to large segments.
type LargeSegmentsService struct {
	client *APIClient
}

// LargeSegmentCreateRequest is the body for creating a large segment.
type LargeSegmentCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// List returns all large segments in the workspace.
func (s *LargeSegmentsService) List(workspaceID string) ([]LargeSegmentMetadata, error) {
	u := s.client.BasePath + largeSegmentsPath + "/ws/" + workspaceID
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
		return nil, fmt.Errorf("large segments list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out []LargeSegmentMetadata
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListInEnvironment returns large segment definitions in an environment.
func (s *LargeSegmentsService) ListInEnvironment(workspaceID, environmentID string) ([]LargeSegmentMetadata, error) {
	u := s.client.BasePath + largeSegmentsPath + "/ws/" + workspaceID + "/environments/" + environmentID
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
		return nil, fmt.Errorf("large segments list in environment: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out []LargeSegmentMetadata
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns a large segment by name in the workspace.
func (s *LargeSegmentsService) Get(workspaceID, name string) (*LargeSegmentMetadata, error) {
	encodedName := url.QueryEscape(name)
	u := s.client.BasePath + largeSegmentsPath + "/ws/" + workspaceID + "/" + encodedName
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
		return nil, fmt.Errorf("large segment get: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out LargeSegmentMetadata
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create creates a large segment in the workspace for the given traffic type.
func (s *LargeSegmentsService) Create(workspaceID, trafficTypeID string, req LargeSegmentCreateRequest) (*LargeSegmentMetadata, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + largeSegmentsPath + "/ws/" + workspaceID + "/trafficTypes/" + trafficTypeID
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
		return nil, fmt.Errorf("large segment create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out LargeSegmentMetadata
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateDefinitionForEnvironment creates a large segment definition in an environment.
func (s *LargeSegmentsService) CreateDefinitionForEnvironment(environmentID, segmentName string) error {
	encodedName := url.QueryEscape(segmentName)
	u := s.client.BasePath + largeSegmentsPath + "/" + environmentID + "/" + encodedName
	req, err := http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("large segment create in environment: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	return nil
}

// DeleteInEnvironment removes the large segment from an environment.
func (s *LargeSegmentsService) DeleteInEnvironment(environmentID, segmentName string) error {
	encodedName := url.QueryEscape(segmentName)
	u := s.client.BasePath + largeSegmentsPath + "/" + environmentID + "/" + encodedName
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
		return fmt.Errorf("large segment delete in environment: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}

// Delete removes the large segment from the workspace.
func (s *LargeSegmentsService) Delete(workspaceID, segmentName string) error {
	encodedName := url.QueryEscape(segmentName)
	u := s.client.BasePath + largeSegmentsPath + "/ws/" + workspaceID + "/" + encodedName
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
		return fmt.Errorf("large segment delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}

// Change request types for updating large segment keys (UPLOAD) or removing all keys (ARCHIVE).
// See https://docs.split.io/reference/create-change-request#open-change-request-to-add-members-to-a-large-segment

const changeRequestsPath = "/internal/api/v2/changeRequests"

// LargeSegmentChangeRequestRequest is the body for creating a large-segment change request.
type LargeSegmentChangeRequestRequest struct {
	LargeSegment *LargeSegmentRef   `json:"largeSegment"`
	OperationType string            `json:"operationType"` // "UPLOAD" or "ARCHIVE"
	Title         string            `json:"title,omitempty"`
	Comment       string            `json:"comment,omitempty"`
	Approvers     []string          `json:"approvers,omitempty"`
}

// LargeSegmentRef references a large segment by name.
type LargeSegmentRef struct {
	Name string `json:"name"`
}

// LargeSegmentChangeRequestResponse is the response from creating a change request.
// For operationType UPLOAD, transactionMetadata contains the upload URL (valid ~5 minutes).
type LargeSegmentChangeRequestResponse struct {
	ID                   string                    `json:"id"`
	Status               string                    `json:"status"`
	TransactionMetadata  *ChangeRequestTxMetadata  `json:"transactionMetadata,omitempty"`
}

// ChangeRequestTxMetadata holds the upload URL and headers for large segment CSV upload.
type ChangeRequestTxMetadata struct {
	Method   string            `json:"method"`
	URL      string            `json:"url"`
	Headers  map[string][]string `json:"headers,omitempty"`
}

// LargeSegmentUploadCredentials holds the presigned upload URL and optional Host header.
// Use UploadKeysToLargeSegment to send CSV content; URL expires in ~5 minutes.
type LargeSegmentUploadCredentials struct {
	UploadURL string
	Method    string
	Host      string
	ChangeRequestID string
}

// OpenChangeRequestForLargeSegmentUpload creates a change request to add/replace keys in a large segment.
// The API returns a presigned upload URL (valid ~5 min). Call UploadKeysToLargeSegment with the returned
// credentials and your CSV body to complete the update. Full list of keys replaces the previous version.
func (s *LargeSegmentsService) OpenChangeRequestForLargeSegmentUpload(workspaceID, environmentID, segmentName, title, comment string, approvers []string) (*LargeSegmentUploadCredentials, error) {
	body := LargeSegmentChangeRequestRequest{
		LargeSegment: &LargeSegmentRef{Name: segmentName},
		OperationType: "UPLOAD",
		Title:   title,
		Comment: comment,
		Approvers: approvers,
	}
	jsonBody, _ := json.Marshal(body)
	u := s.client.BasePath + changeRequestsPath + "/ws/" + url.PathEscape(workspaceID) + "/environments/" + url.PathEscape(environmentID)
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("create change request: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var cr LargeSegmentChangeRequestResponse
	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		return nil, err
	}
	if cr.TransactionMetadata == nil || cr.TransactionMetadata.URL == "" {
		return nil, fmt.Errorf("change request response missing transactionMetadata.url")
	}
	creds := &LargeSegmentUploadCredentials{
		UploadURL:        cr.TransactionMetadata.URL,
		Method:           cr.TransactionMetadata.Method,
		ChangeRequestID:  cr.ID,
	}
	if cr.TransactionMetadata.Method == "" {
		creds.Method = http.MethodPut
	}
	if cr.TransactionMetadata.Headers != nil {
		if h := cr.TransactionMetadata.Headers["Host"]; len(h) > 0 {
			creds.Host = h[0]
		}
	}
	return creds, nil
}

// UploadKeysToLargeSegment uploads CSV content to the presigned URL from OpenChangeRequestForLargeSegmentUpload.
// The CSV must be the full list of keys (one key per line, no header) per Split CSV import format.
func (s *LargeSegmentsService) UploadKeysToLargeSegment(creds *LargeSegmentUploadCredentials, csvContent []byte) error {
	if creds.Method == "" {
		creds.Method = http.MethodPut
	}
	req, err := http.NewRequest(creds.Method, creds.UploadURL, bytes.NewReader(csvContent))
	if err != nil {
		return err
	}
	if creds.Host != "" {
		req.Header.Set("Host", creds.Host)
	}
	req.Header.Set("Content-Type", "text/csv")
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload large segment keys: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	return nil
}

// BuildCSVFromKeys builds a CSV body for large segment import: one key per line, no header.
// Keys containing comma, newline, or double-quote are quoted per CSV rules.
func BuildCSVFromKeys(keys []string) []byte {
	var buf strings.Builder
	w := csv.NewWriter(&buf)
	for _, k := range keys {
		_ = w.Write([]string{k})
	}
	w.Flush()
	return []byte(strings.TrimSpace(buf.String()))
}

// UpdateKeysInEnvironment replaces all keys in a large segment in the given environment.
// It opens a change request (UPLOAD), then uploads the CSV built from keys. Use when approval
// flows allow empty approvers or auto-approve. The full key list replaces the previous version.
func (s *LargeSegmentsService) UpdateKeysInEnvironment(workspaceID, environmentID, segmentName string, keys []string, opts *UpdateLargeSegmentKeysOptions) error {
	if opts == nil {
		opts = &UpdateLargeSegmentKeysOptions{}
	}
	creds, err := s.OpenChangeRequestForLargeSegmentUpload(workspaceID, environmentID, segmentName, opts.Title, opts.Comment, opts.Approvers)
	if err != nil {
		return err
	}
	csvBody := BuildCSVFromKeys(keys)
	return s.UploadKeysToLargeSegment(creds, csvBody)
}

// UpdateLargeSegmentKeysOptions are optional arguments for UpdateKeysInEnvironment.
type UpdateLargeSegmentKeysOptions struct {
	Title     string
	Comment   string
	Approvers []string
}
