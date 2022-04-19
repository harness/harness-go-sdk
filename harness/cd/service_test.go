package cd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListServices(t *testing.T) {
	c := getClient()
	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		services, pagination, err := c.ServiceClient.ListServices(limit, offset)
		require.NoError(t, err, "Failed to list services: %s", err)
		require.NotEmpty(t, services, "No services found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(services) == limit
		offset += limit
	}
}

func TestListServicesByApplicationId(t *testing.T) {
	c := getClient()

	app, err := c.ApplicationClient.GetApplicationByName("Harness Demo Application")
	require.NoError(t, err, "Failed to get application: %s", err)

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		services, pagination, err := c.ServiceClient.ListServicesByApplicationId(app.Id, 100, 0)
		require.NoError(t, err, "Failed to list services: %s", err)
		require.NotEmpty(t, services, "No services found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(services) == limit
		offset += limit
	}
}
