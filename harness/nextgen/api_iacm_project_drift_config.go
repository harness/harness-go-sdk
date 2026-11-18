/*
 * Infrastructure as Code Management (IACM)
 *
 * Project-level drift detection configuration for Harness IaCM.
 * Manages whether scheduled drift detection is enabled across all workspaces
 * in a project for a given provisioner (terraform, opentofu, terragrunt, awscdk).
 *
 * API endpoint: /iacm/api/orgs/{org}/projects/{project}/settings/drift-config
 *
 * API docs:
 *   GET    https://apidocs.harness.io/settings/settings_get-project-drift
 *   PUT    https://apidocs.harness.io/settings/settings_enable-project-drift
 *   DELETE https://apidocs.harness.io/settings/settings_disable-project-drift
 */
package nextgen

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type IacmProjectDriftConfigApiService service

// UpsertIacmProjectDriftConfig enables or updates project-level drift automation for all workspaces
// under the given provisioner. Uses PUT upsert semantics — create and update are the same call.
// Use DisableIacmProjectDriftConfig to remove the configuration.
func (a *IacmProjectDriftConfigApiService) UpsertIacmProjectDriftConfig(
	ctx context.Context,
	harnessAccount string,
	org string,
	project string,
	body IacmProjectDriftConfigRequestBody,
) (*http.Response, error) {
	return a.sendIacmProjectDriftConfigRequest(ctx, "PUT", harnessAccount, org, project, &body, "")
}

// DisableIacmProjectDriftConfig disables drift automation for all workspaces in a project
// for the given provisioner. Enqueues a background job to remove all drift triggers.
func (a *IacmProjectDriftConfigApiService) DisableIacmProjectDriftConfig(
	ctx context.Context,
	harnessAccount string,
	org string,
	project string,
	provisioner string,
) (*http.Response, error) {
	return a.sendIacmProjectDriftConfigRequest(ctx, "DELETE", harnessAccount, org, project, nil, provisioner)
}

// GetIacmProjectDriftConfig returns whether drift automation is currently enabled at the
// project level for the given provisioner, along with its pipeline_id and schedule.
func (a *IacmProjectDriftConfigApiService) GetIacmProjectDriftConfig(
	ctx context.Context,
	harnessAccount string,
	org string,
	project string,
	provisioner string,
) (IacmProjectDriftConfigResponse, *http.Response, error) {
	var result IacmProjectDriftConfigResponse

	resp, err := a.sendIacmProjectDriftConfigRequest(ctx, "GET", harnessAccount, org, project, nil, provisioner)
	if err != nil {
		return result, resp, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, resp, err
	}

	if err := a.client.decode(&result, body, resp.Header.Get("Content-Type")); err != nil {
		return result, resp, err
	}

	return result, resp, nil
}

// sendIacmProjectDriftConfigRequest sends the REST request to the drift config endpoint.
// PUT:    body=IacmProjectDriftConfigRequestBody (provisioner field inside body), provisioner="".
// GET:    body=nil, provisioner=query param to filter the config by provisioner type.
// DELETE: body=nil, provisioner=query param to identify which provisioner config to remove.
func (a *IacmProjectDriftConfigApiService) sendIacmProjectDriftConfigRequest(
	ctx context.Context,
	method string,
	harnessAccount string,
	org string,
	project string,
	body interface{},
	provisioner string,
) (*http.Response, error) {
	urlPath := a.client.cfg.BasePath + iacmProjectDriftConfigPath
	urlPath = strings.Replace(urlPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	urlPath = strings.Replace(urlPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}

	if provisioner != "" {
		queryParams.Add("provisioner", provisioner)
	}

	contentTypes := []string{}
	if body != nil {
		contentTypes = []string{"application/json"}
	}
	if ct := selectHeaderContentType(contentTypes); ct != "" {
		headerParams["Content-Type"] = ct
	}

	accepts := []string{"application/json", "application/vnd.goa.error"}
	if accept := selectHeaderAccept(accepts); accept != "" {
		headerParams["Accept"] = accept
	}

	headerParams["Harness-Account"] = parameterToString(harnessAccount, "")
	if ctx != nil {
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			key := auth.Key
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			}
			headerParams["x-api-key"] = key
		}
	}

	r, err := a.client.prepareRequest(ctx, urlPath, strings.ToUpper(method), body,
		headerParams, queryParams, url.Values{}, "", nil)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.callAPI(r)
	if err != nil {
		return resp, err
	}
	if resp == nil {
		return nil, fmt.Errorf("IaCM drift config API returned nil response")
	}

	if resp.StatusCode >= http.StatusMultipleChoices {
		respBody, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			return resp, fmt.Errorf("reading IaCM API error response: %w", readErr)
		}
		return resp, decodeIacmError(a.client, resp, respBody)
	}

	return resp, nil
}

// decodeIacmError wraps an error response body in a GenericSwaggerError.
// It attempts to decode the body as IacmError; if that fails (e.g. a gateway
// or proxy returned a different format), the raw status line is used instead
// and the raw body is still available via GenericSwaggerError.Body().
func decodeIacmError(client *APIClient, resp *http.Response, body []byte) error {
	newErr := GenericSwaggerError{
		body:  body,
		error: resp.Status,
	}
	var v IacmError
	if err := client.decode(&v, body, resp.Header.Get("Content-Type")); err == nil {
		newErr.model = v
	}
	return newErr
}
