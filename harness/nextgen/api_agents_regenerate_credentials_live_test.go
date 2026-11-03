package nextgen_test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

/*
CDS-136012 live matrix: RegenerateCredentials must reach agents at account, org and project scope.

Creates a throwaway org, project and one agent per scope, runs the matrix, and removes everything.
Like the other tests in this package it only runs when live credentials are passed:

	go test ./harness/nextgen \
	  -run '^TestAgentServiceForServerRegenerateCredentials_LiveScopeMatrix$' \
	  -v -count=1 \
	  -args \
	  -account <ACCOUNT_ID> \
	  -api-key <API_KEY> \
	  -endpoint https://qa.harness.io/gateway

The offline wire-level tests live in harness/nextgen/clienttests.
*/
func TestAgentServiceForServerRegenerateCredentials_LiveScopeMatrix(t *testing.T) {
	client, ctx := setupClient()
	acct := *accountID
	suffix := strings.ToLower(utils.RandStringBytes(6))
	orgId := "sdkrc" + suffix
	projId := "sdkrcp" + suffix

	_, _, err := client.OrganizationApi.PostOrganization(ctx, nextgen.OrganizationRequest{
		Organization: &nextgen.Organization{Identifier: orgId, Name: orgId},
	}, acct)
	require.NoError(t, err)
	t.Cleanup(func() {
		_, _, _ = client.OrganizationApi.DeleteOrganization(ctx, orgId, acct, &nextgen.OrganizationApiDeleteOrganizationOpts{})
	})

	_, _, err = client.ProjectApi.PostProject(ctx, nextgen.ProjectRequest{
		Project: &nextgen.Project{OrgIdentifier: orgId, Identifier: projId, Name: projId},
	}, acct, &nextgen.ProjectApiPostProjectOpts{OrgIdentifier: optional.NewString(orgId)})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, _, _ = client.ProjectApi.DeleteProject(ctx, projId, acct, &nextgen.ProjectApiDeleteProjectOpts{OrgIdentifier: optional.NewString(orgId)})
	})

	type agent struct {
		id, org, project string
		lastKey          string
	}
	create := func(kind, org, project string) *agent {
		a := &agent{id: "sdkrc" + kind + suffix, org: org, project: project}
		typ := nextgen.MANAGED_ARGO_PROVIDER_V1AgentType
		op := nextgen.ARGO_V1AgentOperator
		created, _, err := client.AgentApi.AgentServiceForServerCreate(ctx, nextgen.V1Agent{
			AccountIdentifier: acct,
			OrgIdentifier:     org,
			ProjectIdentifier: project,
			Identifier:        a.id,
			Name:              a.id,
			Type_:             &typ,
			Operator:          &op,
			Metadata:          &nextgen.V1AgentMetadata{Namespace: "sdk-live-test"},
		})
		require.NoError(t, err, "creating %s agent", kind)
		if created.Credentials != nil {
			a.lastKey = created.Credentials.PrivateKey
		}
		agentRef := a
		t.Cleanup(func() {
			_, _, _ = client.AgentApi.AgentServiceForServerDelete(ctx, agentRef.id, &nextgen.AgentsApiAgentServiceForServerDeleteOpts{
				AccountIdentifier: optional.NewString(acct),
				OrgIdentifier:     optional.NewString(agentRef.org),
				ProjectIdentifier: optional.NewString(agentRef.project),
			})
		})
		return a
	}
	accountAgent := create("acct", "", "")
	orgAgent := create("org", orgId, "")
	projectAgent := create("proj", orgId, projId)

	withScope := func(c context.Context, id, org, project string) (nextgen.V1Agent, int, error) {
		opts := &nextgen.AgentsApiAgentServiceForServerRegenerateCredentialsWithScopeOpts{}
		if org != "" {
			opts.OrgIdentifier = optional.NewString(org)
		}
		if project != "" {
			opts.ProjectIdentifier = optional.NewString(project)
		}
		got, resp, err := client.AgentApi.AgentServiceForServerRegenerateCredentialsWithScope(c, id, acct, opts)
		return got, statusOf(resp), err
	}
	legacy := func(id string) (nextgen.V1Agent, int, error) {
		got, resp, err := client.AgentApi.AgentServiceForServerRegenerateCredentials(ctx, id)
		return got, statusOf(resp), err
	}

	type call func() (nextgen.V1Agent, int, error)
	cases := []struct {
		name   string
		target *agent // non-nil when a successful call must rotate this agent's key
		do     call
		want   int
	}{
		{"account agent: WithScope(account) rotates", accountAgent, func() (nextgen.V1Agent, int, error) { return withScope(ctx, accountAgent.id, "", "") }, http.StatusOK},
		{"account agent: legacy method still works", accountAgent, func() (nextgen.V1Agent, int, error) { return legacy(accountAgent.id) }, http.StatusOK},
		{"org agent: WithScope(org) rotates", orgAgent, func() (nextgen.V1Agent, int, error) { return withScope(ctx, orgAgent.id, orgId, "") }, http.StatusOK},
		{"org agent: WithScope(account only) is 404", nil, func() (nextgen.V1Agent, int, error) { return withScope(ctx, orgAgent.id, "", "") }, http.StatusNotFound},
		{"org agent: legacy method is 404 (the CDS-136012 bug)", nil, func() (nextgen.V1Agent, int, error) { return legacy(orgAgent.id) }, http.StatusNotFound},
		{"project agent: WithScope(org,project) rotates", projectAgent, func() (nextgen.V1Agent, int, error) { return withScope(ctx, projectAgent.id, orgId, projId) }, http.StatusOK},
		{"project agent: org only is 404", nil, func() (nextgen.V1Agent, int, error) { return withScope(ctx, projectAgent.id, orgId, "") }, http.StatusNotFound},
		{"project agent: account only is 404", nil, func() (nextgen.V1Agent, int, error) { return withScope(ctx, projectAgent.id, "", "") }, http.StatusNotFound},
		{"project agent: wrong project is 404", nil, func() (nextgen.V1Agent, int, error) { return withScope(ctx, projectAgent.id, orgId, projId+"x") }, http.StatusNotFound},
		{"project agent: legacy method is 404 (the CDS-136012 bug)", nil, func() (nextgen.V1Agent, int, error) { return legacy(projectAgent.id) }, http.StatusNotFound},
		{"project agent: second rotation rotates again", projectAgent, func() (nextgen.V1Agent, int, error) { return withScope(ctx, projectAgent.id, orgId, projId) }, http.StatusOK},
		{"nonexistent agent is 404", nil, func() (nextgen.V1Agent, int, error) { return withScope(ctx, "sdkrcnone"+suffix, "", "") }, http.StatusNotFound},
		{"invalid API key is 401", nil, func() (nextgen.V1Agent, int, error) {
			bad := context.WithValue(context.Background(), nextgen.ContextAPIKey, nextgen.APIKey{Key: "invalid-key"})
			return withScope(bad, accountAgent.id, "", "")
		}, http.StatusUnauthorized},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%02d %s", i+1, tc.name), func(t *testing.T) {
			got, status, err := tc.do()
			require.Equal(t, tc.want, status, "unexpected HTTP status (err=%v)", err)
			if tc.want != http.StatusOK {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, got.Credentials, "response carries no credentials")
			key := got.Credentials.PrivateKey
			require.NotEmpty(t, key)
			if tc.target.lastKey != "" {
				require.NotEqual(t, tc.target.lastKey, key, "key was not rotated")
			}
			tc.target.lastKey = key
		})
	}
}

func statusOf(resp *http.Response) int {
	if resp == nil {
		return 0
	}
	return resp.StatusCode
}
