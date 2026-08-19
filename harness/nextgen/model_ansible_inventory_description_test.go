package nextgen

import (
	"encoding/json"
	"strings"
	"testing"
)

// IAC-8123: every inventory request and response body on the server carries a
// description, but none of the SDK models declared the field, so it was silently
// dropped on create and update and never surfaced on read. Update is a PUT and the
// server writes `description = :description` unconditionally, so the field has to be
// serialized even when empty - see UpdateInventoryRequest for why it has no
// omitempty.

func TestCreateInventoryRequest_DescriptionSerialized(t *testing.T) {
	testCases := map[string]struct {
		request  CreateInventoryRequest
		expected string
		absent   string
	}{
		"Given a description - then it is sent": {
			request:  CreateInventoryRequest{Identifier: "inventory1", Name: "inventory1", Description: "prod web hosts"},
			expected: `"description":"prod web hosts"`,
		},
		"Given no description - then the field is omitted": {
			request: CreateInventoryRequest{Identifier: "inventory1", Name: "inventory1"},
			absent:  `"description"`,
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			data, err := json.Marshal(tc.request)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			if tc.expected != "" && !strings.Contains(string(data), tc.expected) {
				t.Errorf("expected serialized request to contain %s, got %s", tc.expected, string(data))
			}
			if tc.absent != "" && strings.Contains(string(data), tc.absent) {
				t.Errorf("expected serialized request not to contain %s, got %s", tc.absent, string(data))
			}
		})
	}
}

// The whole point of dropping omitempty: clearing the description has to be
// expressible on the wire, so an empty description is still sent.
func TestUpdateInventoryRequest_EmptyDescriptionStillSerialized(t *testing.T) {
	testCases := map[string]struct {
		request  UpdateInventoryRequest
		expected string
	}{
		"Given a description - then it is sent": {
			request:  UpdateInventoryRequest{Name: "inventory1", Description: "prod web hosts"},
			expected: `"description":"prod web hosts"`,
		},
		"Given a cleared description - then an empty description is still sent": {
			request:  UpdateInventoryRequest{Name: "inventory1"},
			expected: `"description":""`,
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			data, err := json.Marshal(tc.request)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			if !strings.Contains(string(data), tc.expected) {
				t.Errorf("expected serialized request to contain %s, got %s", tc.expected, string(data))
			}
		})
	}
}

func TestInventoryResponses_DescriptionDeserialized(t *testing.T) {
	body := []byte(`{"identifier":"inventory1","name":"inventory1","description":"prod web hosts","data":{}}`)

	testCases := map[string]struct {
		decode func([]byte) (string, error)
	}{
		"Given a create response - then the description is read": {
			decode: func(b []byte) (string, error) {
				var resp CreateInventoryResponse
				err := json.Unmarshal(b, &resp)
				return resp.Description, err
			},
		},
		"Given a show response - then the description is read": {
			decode: func(b []byte) (string, error) {
				var resp ShowInventoryResponse
				err := json.Unmarshal(b, &resp)
				return resp.Description, err
			},
		},
		"Given a list detail - then the description is read": {
			decode: func(b []byte) (string, error) {
				var detail HarnessIacmInventoryDetail
				err := json.Unmarshal(b, &detail)
				return detail.Description, err
			},
		},
	}

	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			got, err := tc.decode(body)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if got != "prod web hosts" {
				t.Errorf("expected description to be %q, got %q", "prod web hosts", got)
			}
		})
	}
}

// The server omits description entirely when it is unset, so an absent field has to
// decode to an empty string rather than failing.
func TestInventoryResponses_AbsentDescriptionDecodesEmpty(t *testing.T) {
	var resp ShowInventoryResponse
	if err := json.Unmarshal([]byte(`{"identifier":"inventory1","name":"inventory1","data":{}}`), &resp); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if resp.Description != "" {
		t.Errorf("expected absent description to decode as empty, got %q", resp.Description)
	}
}
