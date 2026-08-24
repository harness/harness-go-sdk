package idp

import (
	"context"
	"net/url"

	"github.com/antihax/optional"
)

func (c *APIClient) WithAuthContext(ctx context.Context) (*APIClient, context.Context) {
	authCtx := context.WithValue(ctx, ContextAPIKey, APIKey{Key: c.ApiKey})
	return c, authCtx
}

func applyHarnessAccount(headerParams map[string]string, queryParams url.Values, harnessAccount optional.String) {
	if harnessAccount.IsSet() {
		headerParams["Harness-Account"] = parameterToString(harnessAccount.Value(), "")
		queryParams.Add("accountIdentifier", parameterToString(harnessAccount.Value(), ""))
	}
}
