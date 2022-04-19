package cd

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
)

type ServiceClient struct {
	ApiClient *ApiClient
}

func (ac *ServiceClient) ListServices(limit int, offset int) ([]*graphql.Service, *graphql.PageInfo, error) {
	return ac.listServicesInternal(limit, offset, "[]")
}

func (ac *ServiceClient) listServicesInternal(limit int, offset int, filters string) ([]*graphql.Service, *graphql.PageInfo, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query {
			services(limit: %[3]d, offset: %[4]d, filters: %[5]s) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, standardServiceFields, paginationFields, limit, offset, filters),
	}

	res := struct {
		Services struct {
			Nodes    []*graphql.Service
			PageInfo *graphql.PageInfo
		}
	}{}

	err := ac.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.Services.Nodes, res.Services.PageInfo, nil
}

func (ac *ServiceClient) ListServicesByApplicationId(appId string, limit int, offset int) ([]*graphql.Service, *graphql.PageInfo, error) {

	filterString := fmt.Sprintf(`[{
		application: {
			operator: %[1]s,
			values: ["%[2]s"]
		}
	}]`, graphql.IdOperatorTypes.Equals, appId)

	return ac.listServicesInternal(limit, offset, filterString)
}

const (
	standardServiceFields = `
	name
	artifactType
	createdAt
	createdBy {
		email
		externalUserId
		id
		name
	}
	deploymentType
	description
	id
	name
	tags {
		name
		value
	}
`
)
