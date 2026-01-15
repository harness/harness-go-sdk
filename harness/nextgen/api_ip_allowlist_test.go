package nextgen

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetIpConfig(t *testing.T) {
	c, ctx := getClientWithContext()

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	ipAllowlistrequest := IpAllowlistConfigRequest{
		IpAllowlistConfig: &IpAllowlistConfig{
			Name:              name,
			Identifier:        name,
			Description:       name,
			Enabled:           true,
			AllowedSourceType: []AllowedSourceType{"UI"},
			IpAddress:         "0.0.0.0/0",
		},
	}
	fmt.Print(ipAllowlistrequest)
	opts := &IPAllowlistApiCreateIpAllowlistConfigOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	}
	resp, _, err := c.IPAllowlistApiService.CreateIpAllowlistConfig(ctx, ipAllowlistrequest, opts)

	require.NoError(t, err)
	require.NotNil(t, resp.IpAllowlistConfig)
}
