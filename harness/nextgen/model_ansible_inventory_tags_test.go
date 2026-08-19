package nextgen

import (
	"encoding/json"
	"strings"
	"testing"
)

// IAC-8112: the IaCM API accepts and returns Ansible inventory tags as an array
// of strings. The generated models typed Tags as a plain string, so callers such
// as the Terraform provider could not send tags at all - the request either
// failed to compile or serialized `"tags":"..."` and was rejected by the server.

func TestCreateInventoryRequest_TagsSerializedAsArray(t *testing.T) {
	testCases := map[string]struct {
		tags     []string
		expected string
	}{
		"Given multiple tags - then they serialize as a JSON array": {
			tags:     []string{"env:prod", "team:platform"},
			expected: `"tags":["env:prod","team:platform"]`,
		},
		"Given a single tag - then it still serializes as a JSON array": {
			tags:     []string{"env:prod"},
			expected: `"tags":["env:prod"]`,
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			data, err := json.Marshal(CreateInventoryRequest{
				Identifier: "inventory1",
				Name:       "inventory1",
				Type_:      "manual",
				Tags:       tc.tags,
			})
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			if !strings.Contains(string(data), tc.expected) {
				t.Errorf("expected serialized request to contain %s, got %s", tc.expected, string(data))
			}
		})
	}
}

func TestCreateInventoryRequest_NoTagsAreOmitted(t *testing.T) {
	data, err := json.Marshal(CreateInventoryRequest{
		Identifier: "inventory1",
		Name:       "inventory1",
		Type_:      "manual",
	})
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if strings.Contains(string(data), `"tags"`) {
		t.Errorf("expected tags to be omitted when unset, got %s", string(data))
	}
}

// Update is a PUT, so the tags in the request replace the stored set outright.
// Tags must therefore be serialized even when empty - otherwise a caller that
// removed its last tag would send no tags field and the stored tags would be
// left behind as orphans.
func TestUpdateInventoryRequest_TagsAlwaysSerialized(t *testing.T) {
	testCases := map[string]struct {
		tags     []string
		expected string
	}{
		"Given tags - then they serialize as a JSON array": {
			tags:     []string{"env:prod"},
			expected: `"tags":["env:prod"]`,
		},
		"Given an empty slice - then tags are still sent so the stored tags are cleared": {
			tags:     []string{},
			expected: `"tags":[]`,
		},
		"Given a nil slice - then tags are still sent so the stored tags are cleared": {
			tags:     nil,
			expected: `"tags":null`,
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			data, err := json.Marshal(UpdateInventoryRequest{
				Name: "inventory1",
				Tags: tc.tags,
			})
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			if !strings.Contains(string(data), tc.expected) {
				t.Errorf("expected serialized request to contain %s, got %s", tc.expected, string(data))
			}
		})
	}
}

func TestInventoryResponses_TagsDeserializedFromArray(t *testing.T) {
	body := []byte(`{"identifier":"inventory1","name":"inventory1","tags":["env:prod","team:platform"]}`)
	expected := []string{"env:prod", "team:platform"}

	testCases := map[string]struct {
		decode func([]byte) ([]string, error)
	}{
		"Given a create response - then tags decode into a slice": {
			decode: func(b []byte) ([]string, error) {
				var resp CreateInventoryResponse
				err := json.Unmarshal(b, &resp)
				return resp.Tags, err
			},
		},
		"Given a show response - then tags decode into a slice": {
			decode: func(b []byte) ([]string, error) {
				var resp ShowInventoryResponse
				err := json.Unmarshal(b, &resp)
				return resp.Tags, err
			},
		},
		"Given a list detail - then tags decode into a slice": {
			decode: func(b []byte) ([]string, error) {
				var detail HarnessIacmInventoryDetail
				err := json.Unmarshal(b, &detail)
				return detail.Tags, err
			},
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			tags, err := tc.decode(body)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if len(tags) != len(expected) {
				t.Fatalf("expected %d tags, got %d (%v)", len(expected), len(tags), tags)
			}
			for i, tag := range expected {
				if tags[i] != tag {
					t.Errorf("expected tag %d to be %q, got %q", i, tag, tags[i])
				}
			}
		})
	}
}
