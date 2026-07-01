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
	var ipAddress = "0.0.0.0/0"

	ipAllowlistrequest := IpAllowlistConfigRequest{
		IpAllowlistConfig: &IpAllowlistConfig{
			Name:              name,
			Identifier:        name,
			Description:       name,
			Enabled:           true,
			AllowedSourceType: []AllowedSourceType{"UI"},
			IpAddress:         ipAddress,
		},
	}

	opts := &IPAllowlistApiCreateIpAllowlistConfigOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	}
	// Create test
	resp, _, err := c.IPAllowlistApiService.CreateIpAllowlistConfig(ctx, ipAllowlistrequest, opts)
	require.NoError(t, err)
	require.NotNil(t, resp.IpAllowlistConfig)

	id := resp.IpAllowlistConfig.Identifier

	// Validating IP Address Before Creating
	allowodOps := &IPAllowlistApiAllowedIpAddressOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	}

	allowedResp, _, err := c.IPAllowlistApiService.AllowedIpAddress(ctx, ipAddress, allowodOps)
	require.NoError(t, err)
	require.NotNil(t, allowedResp)

	// Update
	updateOps := &IPAllowlistApiUpdateIpAllowlistConfigOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	}

	ipAllowlistrequest.IpAllowlistConfig.Description = "updated description"

	resp, _, err = c.IPAllowlistApiService.UpdateIpAllowlistConfig(ctx, id, ipAllowlistrequest, updateOps)
	require.NoError(t, err)
	require.NotNil(t, resp.IpAllowlistConfig)

	// list IP Allowlist Configs
	listOpts := &IPAllowlistApiGetIpAllowlistConfigsOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	}

	listResp, _, err := c.IPAllowlistApiService.GetIpAllowlistConfigs(ctx, listOpts)
	for _, ipConfig := range listResp {
		fmt.Printf("IP Allowlist Config: %+v\n", ipConfig.IpAllowlistConfig)
	}
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(listResp), 1)

	// Get IP Allowlist Config by ID
	getOpts := &IPAllowlistApiGetIpAllowlistConfigOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	}

	getResp, _, err := c.IPAllowlistApiService.GetIpAllowlistConfig(ctx, id, getOpts)
	require.NoError(t, err)
	require.Equal(t, getResp.IpAllowlistConfig.Identifier, id)

	// Cleaning up using the list
	allowedConfigsIdes := []string{}
	for _, ipConfig := range listResp {
		allowedConfigsIdes = append(allowedConfigsIdes, ipConfig.IpAllowlistConfig.Identifier)
	}

	// Delete the created IP Allowlist Config
	dopts := &IPAllowlistApiDeleteIpAllowlistConfigOpts{
		HarnessAccount: optional.NewString(c.AccountId),
	}
	for _, id := range allowedConfigsIdes {
		_, err = c.IPAllowlistApiService.DeleteIpAllowlistConfig(ctx, id, dopts)
		require.NoError(t, err)
	}

	getResp, _, err = c.IPAllowlistApiService.GetIpAllowlistConfig(ctx, id, getOpts)
	require.Nil(t, getResp.IpAllowlistConfig)
}
