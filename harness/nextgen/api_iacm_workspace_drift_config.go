/*
 * Infrastructure as Code Management (IACM)
 *
 * Workspace-level drift detection configuration for Harness IaCM.
 * Manages whether scheduled drift detection is enabled for a specific workspace
 * for a given provisioner (terraform, opentofu, terragrunt, awscdk).
 *
 * API endpoint: /iacm/api/orgs/{org}/projects/{project}/workspaces/{workspace}/drift-config
 *
 * API docs:
 *   GET    https://apidocs.harness.io/tag/Workspaces#operation/workspace_get_drift_config
 *   PUT    https://apidocs.harness.io/tag/Workspaces#operation/workspace_enable_drift
 *   DELETE https://apidocs.harness.io/tag/Workspaces#operation/workspace_disable_drift
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

type IacmWorkspaceDriftConfigApiService service

// UpsertIacmWorkspaceDriftConfig enables or updates workspace-level drift automation.
// Uses PUT upsert semantics — create and update are the same call.
// Use DisableIacmWorkspaceDriftConfig to remove the configuration.
func (a *IacmWorkspaceDriftConfigApiService) UpsertIacmWorkspaceDriftConfig(
	ctx context.Context,
	harnessAccount string,
	org string,
	project string,
	workspace string,
	body IacmWorkspaceDriftConfigRequestBody,
) (*http.Response, error) {
	return a.sendIacmWorkspaceDriftConfigRequest(ctx, "PUT", harnessAccount, org, project, workspace, &body, "")
}

// DisableIacmWorkspaceDriftConfig disables drift automation for the given workspace and provisioner.
func (a *IacmWorkspaceDriftConfigApiService) DisableIacmWorkspaceDriftConfig(
	ctx context.Context,
	harnessAccount string,
	org string,
	project string,
	workspace string,
	provisioner string,
) (*http.Response, error) {
	return a.sendIacmWorkspaceDriftConfigRequest(ctx, "DELETE", harnessAccount, org, project, workspace, nil, provisioner)
}

// GetIacmWorkspaceDriftConfig returns the current drift automation config for the workspace,
// including whether it is enabled, pipeline_id, schedule, and trigger_id.
func (a *IacmWorkspaceDriftConfigApiService) GetIacmWorkspaceDriftConfig(
	ctx context.Context,
	harnessAccount string,
	org string,
	project string,
	workspace string,
	provisioner string,
) (IacmWorkspaceDriftConfigResponse, *http.Response, error) {
	var result IacmWorkspaceDriftConfigResponse

	resp, err := a.sendIacmWorkspaceDriftConfigRequest(ctx, "GET", harnessAccount, org, project, workspace, nil, provisioner)
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

// sendIacmWorkspaceDriftConfigRequest sends the REST request to the workspace drift config endpoint.
// PUT:    body=IacmWorkspaceDriftConfigRequestBody (provisioner field inside body), provisioner="".
// GET:    body=nil, provisioner=query param.
// DELETE: body=nil, provisioner=query param.
func (a *IacmWorkspaceDriftConfigApiService) sendIacmWorkspaceDriftConfigRequest(
	ctx context.Context,
	method string,
	harnessAccount string,
	org string,
	project string,
	workspace string,
	body interface{},
	provisioner string,
) (*http.Response, error) {
	urlPath := a.client.cfg.BasePath + iacmWorkspaceDriftConfigPath
	urlPath = strings.Replace(urlPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	urlPath = strings.Replace(urlPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)
	urlPath = strings.Replace(urlPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)

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
		return nil, fmt.Errorf("IaCM workspace drift config API returned nil response")
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
