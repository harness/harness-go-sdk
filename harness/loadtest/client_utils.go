package loadtest

import "context"

func (c *APIClient) WithAuthContext(ctx context.Context) (*APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	return c, context.WithValue(ctx, ContextAPIKey, APIKey{Key: c.ApiKey})
}
