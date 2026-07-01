package nextgen

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestInfrastructureRequest_DescriptionEmptyStringIsSerialized(t *testing.T) {
	// IAC-7067: description must be sent on the wire even when empty so the
	// API can clear a previously-set description on update. With omitempty
	// the field would be dropped and the update would be a silent no-op.
	req := InfrastructureRequest{
		Identifier:        "infra1",
		OrgIdentifier:     "org",
		ProjectIdentifier: "proj",
		EnvironmentRef:    "env",
		Name:              "infra1",
		Description:       "",
		Type_:             "KubernetesDirect",
		Yaml:              "infrastructureDefinition:\n  identifier: infra1\n",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if !strings.Contains(string(data), `"description":""`) {
		t.Errorf("expected serialized request to contain \"description\":\"\", got %s", string(data))
	}
}

func TestInfrastructureRequest_DescriptionNonEmptyIsSerialized(t *testing.T) {
	req := InfrastructureRequest{
		Identifier: "infra1",
		Name:       "infra1",
		Type_:      "KubernetesDirect",
		Description: "my infra",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if !strings.Contains(string(data), `"description":"my infra"`) {
		t.Errorf("expected serialized request to contain \"description\":\"my infra\", got %s", string(data))
	}
}
