package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCacCreateGcpCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.GcpCloudProvider).(*cac.GcpCloudProvider)
	cpInput.Name = expectedName
	cpInput.DelegateSelectors = []string{"Primary"}
	cpInput.SkipValidation = true
	cpInput.UseDelegateSelectors = true

	c := getClient()
	cp := &cac.GcpCloudProvider{}
	err := c.ConfigAsCode().UpsertObject(cpInput, cac.GetCloudProviderYamlPath(expectedName), cp)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)
}

func TestCacCreatePhysicalDataCenterCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.PhysicalDataCenterCloudProvider).(*cac.PhysicalDatacenterCloudProvider)
	cpInput.Name = expectedName

	c := getClient()
	cp := &cac.PhysicalDatacenterCloudProvider{}
	err := c.ConfigAsCode().UpsertObject(cpInput, cac.GetCloudProviderYamlPath(expectedName), cp)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)
}
