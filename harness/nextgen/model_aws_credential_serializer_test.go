package nextgen

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAwsCredential_ManualConfigRoundTrip(t *testing.T) {
	credential := AwsCredential{
		Type_:  AwsAuthTypes.ManualConfig,
		Region: "us-east-1",
		ManualConfig: &AwsManualConfigSpec{
			AccessKey:    "AKIAEXAMPLE",
			SecretKeyRef: "account.aws-secret",
		},
	}

	data, err := json.Marshal(credential)
	require.NoError(t, err)
	assert.Contains(t, string(data), `"spec":{"accessKey":"AKIAEXAMPLE","secretKeyRef":"account.aws-secret"}`)

	var decoded AwsCredential
	require.NoError(t, json.Unmarshal(data, &decoded))
	require.NotNil(t, decoded.ManualConfig)
	assert.Equal(t, "AKIAEXAMPLE", decoded.ManualConfig.AccessKey)
	assert.Equal(t, "account.aws-secret", decoded.ManualConfig.SecretKeyRef)
	assert.Equal(t, "us-east-1", decoded.Region)
}

func TestAwsCredential_InheritFromDelegateOmitsSpec(t *testing.T) {
	data, err := json.Marshal(AwsCredential{Type_: AwsAuthTypes.InheritFromDelegate})
	require.NoError(t, err)
	assert.NotContains(t, string(data), `"spec"`)

	var decoded AwsCredential
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Nil(t, decoded.ManualConfig)
	assert.Nil(t, decoded.OidcConfig)
}

func TestAwsCredential_UnknownTypePanics(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = json.Marshal(AwsCredential{Type_: "Bogus"})
	})

	assert.Panics(t, func() {
		_ = json.Unmarshal([]byte(`{"type":"Bogus","spec":{}}`), &AwsCredential{})
	})
}
