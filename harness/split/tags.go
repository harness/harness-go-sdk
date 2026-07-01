/*
 * Split API client - Tags service (associate tags with objects).
 * https://docs.split.io/reference/associate-tags
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

const tagsPath = "/internal/api/v2/tags"

// ObjectTypeSplit is the object type for feature flags (splits).
const ObjectTypeSplit = "Split"

// TagsService provides tag association for Split objects (e.g. feature flags).
type TagsService struct {
	client *APIClient
}

// AssociateTags associates tags with an existing object (e.g. a feature flag).
// workspaceID is the Split workspace ID; objectName is the object name (e.g. split name);
// objectType should be ObjectTypeSplit for feature flags; tags is the list of tag names.
func (s *TagsService) AssociateTags(workspaceID, objectName, objectType string, tags []string) error {
	if objectType == "" {
		objectType = ObjectTypeSplit
	}
	u := s.client.BasePath + tagsPath + "/ws/" + url.PathEscape(workspaceID) + "/object/" + url.PathEscape(objectName) + "/objecttype/" + url.PathEscape(objectType)
	body, err := json.Marshal(tags)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("associate tags: %d %s: %s", resp.StatusCode, resp.Status, string(b))
	}
	return nil
}
