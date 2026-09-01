package nextgen

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAwsSmCredentialSpecOidcConfig_SerializesOidcSessionTagKeys(t *testing.T) {
	spec := AwsSmCredentialSpecOidcConfig{
		IamRoleArn: "arn:aws:iam::123456789012:role/example",
		OidcSessionTagKeys: []string{
			"account_id",
			"organization_id",
			"project_id",
		},
	}

	data, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if !strings.Contains(string(data), `"oidcSessionTagKeys":["account_id","organization_id","project_id"]`) {
		t.Fatalf("unexpected serialized spec: %s", string(data))
	}

	var decoded AwsSmCredentialSpecOidcConfig
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.IamRoleArn != spec.IamRoleArn {
		t.Fatalf("expected iamRoleArn %q, got %q", spec.IamRoleArn, decoded.IamRoleArn)
	}

	if len(decoded.OidcSessionTagKeys) != len(spec.OidcSessionTagKeys) {
		t.Fatalf("expected %d oidcSessionTagKeys, got %d", len(spec.OidcSessionTagKeys), len(decoded.OidcSessionTagKeys))
	}
}

func TestAwsSmCredentialSpecOidcConfig_OmitsEmptyOidcSessionTagKeys(t *testing.T) {
	spec := AwsSmCredentialSpecOidcConfig{
		IamRoleArn: "arn:aws:iam::123456789012:role/example",
	}

	data, err := json.Marshal(spec)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if strings.Contains(string(data), "oidcSessionTagKeys") {
		t.Fatalf("expected oidcSessionTagKeys to be omitted when empty, got: %s", string(data))
	}
}
