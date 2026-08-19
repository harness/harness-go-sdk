package nextgen

import (
	"encoding/json"
	"strings"
	"testing"
)

// IAC-8122: the IaCM API carries Ansible inventory data as a JSON object, but the
// published spec types it as a string (the design declares the attribute without a
// type argument, so Goa falls back to String while the struct:field:type meta binds
// the server field to encoding/json.RawMessage). ShowInventoryResponse had already
// been corrected by hand; UpdateInventoryRequest and HarnessIacmInventoryDetail had
// not, so a caller could not carry a document read from a show response into an
// update, and sent `"data":""` instead - rejected by the server as invalid
// inventory data.

const inventoryDataJSON = `{"hosts":{"host1.example.com":null,"host2.example.com":null}}`

func TestUpdateInventoryRequest_DataSerializedAsJsonObject(t *testing.T) {
	data, err := json.Marshal(UpdateInventoryRequest{
		Name: "inventory1",
		Data: json.RawMessage(inventoryDataJSON),
	})
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	expected := `"data":` + inventoryDataJSON
	if !strings.Contains(string(data), expected) {
		t.Errorf("expected serialized request to contain %s, got %s", expected, string(data))
	}

	// A string-typed field would have produced a quoted, escaped document.
	if strings.Contains(string(data), `"data":"`) {
		t.Errorf("expected data to serialize as a JSON object, not a quoted string, got %s", string(data))
	}
}

// Data is required by the API, so an unset value has to surface as an explicit
// null rather than silently disappearing from the body.
func TestUpdateInventoryRequest_UnsetDataSerializedAsNull(t *testing.T) {
	data, err := json.Marshal(UpdateInventoryRequest{Name: "inventory1"})
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if !strings.Contains(string(data), `"data":null`) {
		t.Errorf("expected unset data to serialize as null, got %s", string(data))
	}
}

func TestInventoryResponses_DataDeserializedFromJsonObject(t *testing.T) {
	body := []byte(`{"identifier":"inventory1","name":"inventory1","type":"manual","data":` + inventoryDataJSON + `}`)

	testCases := map[string]struct {
		decode func([]byte) (json.RawMessage, error)
	}{
		"Given a show response - then data decodes as a raw JSON object": {
			decode: func(b []byte) (json.RawMessage, error) {
				var resp ShowInventoryResponse
				err := json.Unmarshal(b, &resp)
				return resp.Data, err
			},
		},
		"Given a list detail - then data decodes as a raw JSON object": {
			decode: func(b []byte) (json.RawMessage, error) {
				var detail HarnessIacmInventoryDetail
				err := json.Unmarshal(b, &detail)
				return detail.Data, err
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

			if string(got) != inventoryDataJSON {
				t.Errorf("expected data to be %s, got %s", inventoryDataJSON, string(got))
			}
		})
	}
}

// The point of the change: a document read from a show response has to be usable
// verbatim as the data on the following update, so the provider can update name or
// tags without having to rebuild the inventory document.
func TestInventoryData_ShowResponseRoundTripsIntoUpdateRequest(t *testing.T) {
	var show ShowInventoryResponse
	if err := json.Unmarshal([]byte(`{"identifier":"inventory1","name":"inventory1","data":`+inventoryDataJSON+`}`), &show); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	data, err := json.Marshal(UpdateInventoryRequest{
		Name: show.Name,
		Data: show.Data,
		Tags: []string{"env:prod"},
	})
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	expected := `"data":` + inventoryDataJSON
	if !strings.Contains(string(data), expected) {
		t.Errorf("expected round-tripped request to contain %s, got %s", expected, string(data))
	}
}
