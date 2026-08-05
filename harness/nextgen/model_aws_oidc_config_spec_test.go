package nextgen

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAwsOidcConfigSpec_SerializesOidcSessionTagKeys(t *testing.T) {
	spec := AwsOidcConfigSpec{
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

	var decoded AwsOidcConfigSpec
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

func TestAwsCredential_OidcAuthenticationSerializesOidcSessionTagKeys(t *testing.T) {
	credential := AwsCredential{
		Type_:  AwsAuthTypes.OidcAuthentication,
		Region: "me-central-1",
		OidcConfig: &AwsOidcConfigSpec{
			IamRoleArn: "arn:aws:iam::123456789012:role/example",
			OidcSessionTagKeys: []string{
				"account_id",
				"environment_id",
			},
		},
	}

	data, err := json.Marshal(credential)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded AwsCredential
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.OidcConfig == nil {
		t.Fatal("expected oidc config to be populated")
	}

	if len(decoded.OidcConfig.OidcSessionTagKeys) != 2 {
		t.Fatalf("expected 2 oidcSessionTagKeys, got %d", len(decoded.OidcConfig.OidcSessionTagKeys))
	}
}
