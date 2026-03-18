/*
 * Split API client - API Keys service.
 * https://docs.split.io/reference/api-keys-overview
 */
package split

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiKeysPath = "/internal/api/v2/apiKeys"

// ApiKeysService provides access to API key create/delete.
// Note: When using harness_token (x-api-key), some key types are deprecated.
type ApiKeysService struct {
	client *APIClient
}

// KeyRequest is the body for creating an API key.
type KeyRequest struct {
	Name         string                 `json:"name"`
	KeyType      string                 `json:"apiKeyType"` // e.g. "admin", "server_side", "client_side"
	Roles        []string               `json:"roles"`
	Environments []KeyEnvironmentRequest `json:"environments"`
	Workspace    *KeyWorkspaceRequest   `json:"workspace"`
}

// KeyEnvironmentRequest references an environment.
type KeyEnvironmentRequest struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// KeyWorkspaceRequest references a workspace.
type KeyWorkspaceRequest struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// KeyResponse is the created key (includes the secret key value).
type KeyResponse struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Roles     []string `json:"roles"`
	Type      string   `json:"type"`
	ApiKeyType string  `json:"apiKeyType"`
	Key       string   `json:"key"` // the actual API key value
}

// Create creates an API key.
func (s *ApiKeysService) Create(req KeyRequest) (*KeyResponse, error) {
	body, _ := json.Marshal(req)
	u := s.client.BasePath + apiKeysPath
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
		return nil, fmt.Errorf("api key create: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	var out KeyResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete deletes an API key by key ID.
func (s *ApiKeysService) Delete(keyID string) error {
	u := s.client.BasePath + apiKeysPath + "/" + keyID
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
		return fmt.Errorf("api key delete: %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}
