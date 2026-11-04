package nextgen

import (
	"encoding/json"
	"testing"
)

func TestCeAwsConnector_CredentialOidcAuthenticationRoundTrip(t *testing.T) {
	raw := `{
		"type": "CEAws",
		"spec": {
			"awsAccountId": "123456789012",
			"crossAccountAccess": {"crossAccountRoleArn": "arn:aws:iam::123456789012:role/example"},
			"featuresEnabled": ["BILLING"],
			"credential": {
				"type": "OidcAuthentication",
				"region": "us-east-1",
				"spec": {"iamRoleArn": "arn:aws:iam::123456789012:role/oidc-example"}
			}
		}
	}`

	var connector ConnectorInfo
	if err := json.Unmarshal([]byte(raw), &connector); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if connector.AwsCC == nil {
		t.Fatal("expected AwsCC connector to be populated")
	}
	if connector.AwsCC.Credential == nil {
		t.Fatal("expected credential to be populated")
	}
	if connector.AwsCC.Credential.Type_ != "OidcAuthentication" {
		t.Fatalf("expected type OidcAuthentication, got %q", connector.AwsCC.Credential.Type_)
	}
	if connector.AwsCC.Credential.Spec == nil || connector.AwsCC.Credential.Spec.IamRoleArn != "arn:aws:iam::123456789012:role/oidc-example" {
		t.Fatalf("unexpected spec: %+v", connector.AwsCC.Credential.Spec)
	}

	data, err := json.Marshal(connector)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var roundTripped ConnectorInfo
	if err := json.Unmarshal(data, &roundTripped); err != nil {
		t.Fatalf("round-trip Unmarshal failed: %v", err)
	}
	if roundTripped.AwsCC.Credential.Spec.IamRoleArn != connector.AwsCC.Credential.Spec.IamRoleArn {
		t.Fatalf("round-trip mismatch: %+v", roundTripped.AwsCC.Credential.Spec)
	}
}

func TestCeAwsConnector_CredentialDefaultHasNilSpec(t *testing.T) {
	raw := `{
		"type": "CEAws",
		"spec": {
			"awsAccountId": "123456789012",
			"featuresEnabled": ["BILLING"],
			"credential": {"type": "Default", "spec": null, "region": null}
		}
	}`

	var connector ConnectorInfo
	if err := json.Unmarshal([]byte(raw), &connector); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if connector.AwsCC.Credential == nil {
		t.Fatal("expected credential to be populated")
	}
	if connector.AwsCC.Credential.Type_ != "Default" {
		t.Fatalf("expected type Default, got %q", connector.AwsCC.Credential.Type_)
	}
	if connector.AwsCC.Credential.Spec != nil {
		t.Fatalf("expected nil spec for Default credential, got %+v", connector.AwsCC.Credential.Spec)
	}

	data, err := json.Marshal(connector)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var roundTripped ConnectorInfo
	if err := json.Unmarshal(data, &roundTripped); err != nil {
		t.Fatalf("round-trip Unmarshal failed: %v", err)
	}
	if roundTripped.AwsCC.Credential.Spec != nil {
		t.Fatalf("round-trip expected nil spec, got %+v", roundTripped.AwsCC.Credential.Spec)
	}
}
