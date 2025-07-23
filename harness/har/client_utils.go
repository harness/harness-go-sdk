package har

import "context"

func (c *APIClient) WithAuthContext(ctx context.Context) (*APIClient, context.Context) {
	authCtx := context.WithValue(ctx, ContextAPIKeys, APIKey{Key: c.ApiKey})
	return c, authCtx
}
