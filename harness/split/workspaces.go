/*
 * Split API client - Workspaces service (read-only).
 * GET https://api.split.io/internal/api/v2/workspaces
 */
package split

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const workspacesPath = "/internal/api/v2/workspaces"

// Workspace represents a Split project (workspace).
type Workspace struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	Type                     string `json:"type"`
	OrganizationIdentifier   string `json:"organizationIdentifier,omitempty"`
	ProjectIdentifier        string `json:"projectIdentifier,omitempty"`
	RequiresTitleAndComments bool   `json:"requiresTitleAndComments"`
}

// WorkspacesListResponse is the response from GET workspaces.
type WorkspacesListResponse struct {
	Objects    []Workspace `json:"objects"`
	Offset     int         `json:"offset"`
	Limit      int         `json:"limit"`
	TotalCount int         `json:"totalCount"`
}

// WorkspacesService provides read-only access to Split workspaces (projects).
type WorkspacesService struct {
	client *APIClient
}

// ListOpts holds optional parameters for List.
type ListOpts struct {
	Offset                   int
	Limit                    int
	Name                     string
	NameOp                   string // IS, STARTS_WITH, CONTAINS
	OrganizationIdentifier   string
	ProjectIdentifier        string
}

// List returns workspaces with optional filters. Pagination: if no limit set, fetches all pages (limit 100 per request; API max is 100).
func (s *WorkspacesService) List(opts *ListOpts) ([]Workspace, error) {
	var all []Workspace
	offset := 0
	limit := 100
	if opts != nil {
		if opts.Offset > 0 {
			offset = opts.Offset
		}
		if opts.Limit > 0 {
			limit = opts.Limit
		}
	}
	for {
		resp, err := s.listPage(offset, limit, opts)
		if err != nil {
			return nil, err
		}
		all = append(all, resp.Objects...)
		if len(resp.Objects) < limit || offset+len(resp.Objects) >= resp.TotalCount {
			break
		}
		offset += len(resp.Objects)
		if opts != nil && opts.Limit > 0 {
			break
		}
	}
	return all, nil
}

func (s *WorkspacesService) listPage(offset, limit int, opts *ListOpts) (*WorkspacesListResponse, error) {
	u := s.client.BasePath + workspacesPath
	parsed, _ := url.Parse(u)
	q := parsed.Query()
	q.Set("offset", fmt.Sprintf("%d", offset))
	q.Set("limit", fmt.Sprintf("%d", limit))
	if opts != nil {
		if opts.Name != "" {
			q.Set("name", opts.Name)
		}
		if opts.NameOp != "" {
			q.Set("nameOp", opts.NameOp)
		}
		if opts.OrganizationIdentifier != "" {
			q.Set("organizationIdentifier", opts.OrganizationIdentifier)
		}
		if opts.ProjectIdentifier != "" {
			q.Set("projectIdentifier", opts.ProjectIdentifier)
		}
	}
	parsed.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, parsed.String(), nil)
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
		return nil, fmt.Errorf("workspaces list: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	var out WorkspacesListResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// FindByName returns workspaces matching the name filter. nameOp: IS (default), STARTS_WITH, CONTAINS.
func (s *WorkspacesService) FindByName(name, nameOp string) ([]Workspace, error) {
	opts := &ListOpts{Name: name, NameOp: nameOp}
	if nameOp == "" {
		opts.NameOp = "IS"
	}
	return s.List(opts)
}

// FindByOrganizationAndProject returns workspaces matching the Harness org and project identifiers.
func (s *WorkspacesService) FindByOrganizationAndProject(organizationIdentifier, projectIdentifier string) ([]Workspace, error) {
	return s.List(&ListOpts{
		OrganizationIdentifier: organizationIdentifier,
		ProjectIdentifier:      projectIdentifier,
	})
}

// ResolveWorkspaceID returns the workspace ID for the given Harness organization and project identifiers.
// It calls FindByOrganizationAndProject and returns the first workspace ID, or an error if none or multiple.
func (s *WorkspacesService) ResolveWorkspaceID(organizationIdentifier, projectIdentifier string) (string, error) {
	ws, err := s.FindByOrganizationAndProject(organizationIdentifier, projectIdentifier)
	if err != nil {
		return "", err
	}
	if len(ws) == 0 {
		return "", fmt.Errorf("no workspace found for organization %q project %q", organizationIdentifier, projectIdentifier)
	}
	if len(ws) > 1 {
		return "", fmt.Errorf("multiple workspaces (%d) found for organization %q project %q", len(ws), organizationIdentifier, projectIdentifier)
	}
	return ws[0].ID, nil
}
