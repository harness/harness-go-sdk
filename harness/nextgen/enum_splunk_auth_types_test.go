package nextgen

import (
	"encoding/json"
	"testing"
)

func TestSplunkAuthTypeValues(t *testing.T) {
	tests := []struct {
		name     string
		authType SplunkAuthType
		expected string
	}{
		{
			name:     "UsernamePassword matches backend expectation",
			authType: SplunkAuthTypes.UsernamePassword,
			expected: "UsernamePassword",
		},
		{
			name:     "BearerToken matches backend expectation",
			authType: SplunkAuthTypes.BearerToken,
			expected: "Bearer Token(HTTP Header)",
		},
		{
			name:     "HECToken matches backend expectation",
			authType: SplunkAuthTypes.HECToken,
			expected: "HEC Token",
		},
		{
			name:     "None matches backend expectation (Anonymous)",
			authType: SplunkAuthTypes.None,
			expected: "Anonymous",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.authType.String() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.authType.String())
			}
		})
	}
}

func TestSplunkConnectorJSONSerialization(t *testing.T) {
	tests := []struct {
		name      string
		connector SplunkConnector
		checkFunc func(t *testing.T, jsonData map[string]interface{})
	}{
		{
			name: "Bearer Token serializes with correct field names",
			connector: SplunkConnector{
				SplunkUrl: "https://splunk.example.com",
				AccountId: "test-account",
				AuthType:  SplunkAuthTypes.BearerToken,
				TokenRef:  "account.bearerToken",
			},
			checkFunc: func(t *testing.T, jsonData map[string]interface{}) {
				if jsonData["type"] != "Bearer Token(HTTP Header)" {
					t.Errorf("Expected type field to be 'Bearer Token(HTTP Header)', got %v", jsonData["type"])
				}
				if jsonData["tokenRef"] != "account.bearerToken" {
					t.Errorf("Expected tokenRef field, got %v", jsonData["tokenRef"])
				}
			},
		},
		{
			name: "HEC Token serializes with correct field names",
			connector: SplunkConnector{
				SplunkUrl: "https://splunk.example.com",
				AccountId: "test-account",
				AuthType:  SplunkAuthTypes.HECToken,
				TokenRef:  "account.hecToken",
			},
			checkFunc: func(t *testing.T, jsonData map[string]interface{}) {
				if jsonData["type"] != "HEC Token" {
					t.Errorf("Expected type field to be 'HEC Token', got %v", jsonData["type"])
				}
				if jsonData["tokenRef"] != "account.hecToken" {
					t.Errorf("Expected tokenRef field, got %v", jsonData["tokenRef"])
				}
			},
		},
		{
			name: "Username/Password serializes with correct field names",
			connector: SplunkConnector{
				SplunkUrl:   "https://splunk.example.com",
				AccountId:   "test-account",
				AuthType:    SplunkAuthTypes.UsernamePassword,
				Username:    "testuser",
				PasswordRef: "account.password",
			},
			checkFunc: func(t *testing.T, jsonData map[string]interface{}) {
				if jsonData["type"] != "UsernamePassword" {
					t.Errorf("Expected type field to be 'UsernamePassword', got %v", jsonData["type"])
				}
				if jsonData["username"] != "testuser" {
					t.Errorf("Expected username field, got %v", jsonData["username"])
				}
				if jsonData["passwordRef"] != "account.password" {
					t.Errorf("Expected passwordRef field, got %v", jsonData["passwordRef"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, err := json.Marshal(tt.connector)
			if err != nil {
				t.Fatalf("Failed to marshal connector: %v", err)
			}

			var jsonData map[string]interface{}
			if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
				t.Fatalf("Failed to unmarshal JSON: %v", err)
			}

			tt.checkFunc(t, jsonData)
		})
	}
}
