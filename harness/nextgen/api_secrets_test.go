package nextgen

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateSecret(t *testing.T) {
	c, ctx := getClientWithContext()

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	secret := &Secret{
		Type_:      SecretTypes.SecretText,
		Name:       id,
		Identifier: id,
		Text: &SecretTextSpec{
			Type_:                   SecretSpecTypes.Text,
			ValueType:               SecretTextValueTypes.Inline,
			Value:                   "test",
			SecretManagerIdentifier: "harnessSecretManager",
		},
	}

	resp, _, err := c.SecretsApi.PostSecret(ctx, SecretRequestWrapper{Secret: secret}, c.AccountId, &SecretsApiPostSecretOpts{})
	require.NoError(t, err)
	require.NotNil(t, resp.Data.Secret)
	require.Equal(t, secret.Name, resp.Data.Secret.Name)
}

func TestForceDeleteSecret(t *testing.T) {
	c, ctx := getClientWithContext()

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	secret := &Secret{
		Type_:      SecretTypes.SecretText,
		Name:       id,
		Identifier: id,
		Text: &SecretTextSpec{
			Type_:                   SecretSpecTypes.Text,
			ValueType:               SecretTextValueTypes.Inline,
			Value:                   "test",
			SecretManagerIdentifier: "harnessSecretManager",
		},
	}

	resp, _, err := c.SecretsApi.PostSecret(ctx, SecretRequestWrapper{Secret: secret}, c.AccountId, &SecretsApiPostSecretOpts{})
	require.NoError(t, err)
	require.NotNil(t, resp.Data.Secret)
	require.Equal(t, secret.Name, resp.Data.Secret.Name)

	r, _, err := c.SecretsApi.DeleteSecretV2(ctx, secret.Identifier, c.AccountId, &SecretsApiDeleteSecretV2Opts{ForceDelete: optional.NewBool(true)})
	require.NoError(t, err)
	require.NotNil(t, r)
	require.Equal(t, r.Data, true)
}
