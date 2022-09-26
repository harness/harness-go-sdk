package nextgen

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetGitopsCluster(t *testing.T) {
	c, ctx := getClientWithContext()

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	// resp, _, err := c.EnvironmentGroupApi.GetEnvironmentGroup(ctx, id, id, id, id, &EnvironmentGroupApiGetEnvironmentGroupOpts{})

	// resp, _, err := c.AgentClusterApi.AgentClusterServiceGet(ctx, id, id, &AgentClusterServiceApiAgentClusterServiceGetOpts{
	// 	AccountIdentifier: optional.NewString(c.AccountId),
	// 	OrgIdentifier:     optional.NewString(id),
	// 	ProjectIdentifier: optional.NewString(id),
	// 	QueryServer:       optional.NewString(id),
	// 	QueryName:         optional.NewString(id),
	// })

	var createRequest = &ClusterClusterCreateRequest{
		Upsert: false,
		Cluster: &Applicationv1alpha1Cluster{
			Server: "TestServer123",
			Name:   "ServerNameTest",
		},
	}

	resp, _, err := c.AgentClusterApi.AgentClusterServiceCreate(ctx, *createRequest, id, &AgentClusterServiceApiAgentClusterServiceCreateOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(id),
		ProjectIdentifier: optional.NewString(id),
		Identifier:        optional.NewString(id),
	})
	fmt.Println("RESP: ", resp)
	require.NoError(t, err)

}
