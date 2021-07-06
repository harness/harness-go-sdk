package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/httphelpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"gopkg.in/yaml.v3"
)

func (c *Client) ConfigAsCode() *ConfigAsCodeClient {
	return &ConfigAsCodeClient{
		ApiClient: c,
	}
}

func FindConfigAsCodeItemByPath(rootItem *cac.ConfigAsCodeItem, path string) *cac.ConfigAsCodeItem {

	if rootItem.DirectoryPath != nil && rootItem.DirectoryPath.Path == path {
		return rootItem
	}

	for _, item := range rootItem.Children {
		if matchingItem := FindConfigAsCodeItemByPath(item, path); matchingItem != nil {
			return matchingItem
		}
	}

	return nil
}

func (c *ConfigAsCodeClient) GetDirectoryItemContent(restName string, uuid string, applicationId string) (*cac.ConfigAsCodeItem, error) {
	path := fmt.Sprintf("/gateway/api/setup-as-code/yaml/%s/%s", restName, uuid)

	req, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, path), nil)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(httphelpers.HeaderAccept, httphelpers.HeaderApplicationJson)
	req.Header.Set(httphelpers.HeaderAuthorization, fmt.Sprintf("Bearer %s", c.ApiClient.BearerToken))

	// Set query parameters
	q := req.URL.Query()
	q.Add(QueryParamAccountId, c.ApiClient.AccountId)

	if applicationId != "" {
		q.Add(QueryParamApplicationId, applicationId)
	}
	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) GetDirectoryTree(applicationId string) (*cac.ConfigAsCodeItem, error) {
	path := "/gateway/api/setup-as-code/yaml/directory"

	req, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, path), nil)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(httphelpers.HeaderApiKey, c.ApiClient.APIKey)
	req.Header.Set(httphelpers.HeaderContentType, httphelpers.HeaderApplicationJson)
	req.Header.Set(httphelpers.HeaderAccept, httphelpers.HeaderApplicationJson)

	// Set query parameters
	q := req.URL.Query()
	q.Add(QueryParamAccountId, c.ApiClient.AccountId)

	if applicationId != "" {
		q.Add(QueryParamApplicationId, applicationId)
	}

	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) UpsertYamlEntity(filePath string, entity interface{}) (*cac.ConfigAsCodeItem, error) {

	payload, err := yaml.Marshal(&entity)
	if err != nil {
		return nil, err
	}

	// Setup form fields
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormField("yamlContent")
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(fw, strings.NewReader(string(payload))); err != nil {
		return nil, err
	}

	w.Close()

	log.Printf("[DEBUG] HTTP Request Body: %s", string(payload))

	req, err := retryablehttp.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, "/gateway/api/setup-as-code/yaml/upsert-entity"), &b)
	if err != nil {
		return nil, err
	}

	// Configure additional headers
	req.Header.Set(httphelpers.HeaderApiKey, c.ApiClient.APIKey)
	req.Header.Set(httphelpers.HeaderContentType, w.FormDataContentType())
	req.Header.Set(httphelpers.HeaderAccept, httphelpers.HeaderApplicationJson)

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(QueryParamAccountId, c.ApiClient.AccountId)
	q.Add(QueryParamYamlFilePath, filePath)
	req.URL.RawQuery = q.Encode()

	item, err := c.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *ConfigAsCodeClient) ExecuteRequest(request *retryablehttp.Request) (*cac.ConfigAsCodeItem, error) {

	log.Printf("[DEBUG] Request url: %s", request.URL)

	res, err := c.ApiClient.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	if ok, err := checkStatusCode(res.StatusCode); !ok {
		return nil, err
	}

	defer res.Body.Close()

	// Make sure we can parse the body properly
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}

	// Check for request throttling
	responseString := buf.String()
	log.Printf("[DEBUG] HTTP response: %d - %s", res.StatusCode, responseString)

	responseObj := &cac.Response{}

	// Unmarshal into our response object
	if err := json.NewDecoder(&buf).Decode(&responseObj); err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	if responseObj.IsEmpty() {
		return nil, errors.New("received an empty response")
	}

	if len(responseObj.ResponseMessages) > 0 {
		return nil, responseObj.ResponseMessages[0].ToError()
	}

	if responseObj.Resource.Status != "" && responseObj.Resource.Status != statusSuccess {
		return nil, fmt.Errorf("%s: %s", responseObj.Resource.Status, responseObj.Resource.ErrorMessage)
	}

	return &responseObj.Resource, nil
}

const (
	statusSuccess = "SUCCESS"
)

func checkStatusCode(statusCode int) (bool, error) {
	if statusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("received http status code '%d'", statusCode)
}

type ConfigAsCodeClient struct {
	ApiClient *Client
}

func (c *ConfigAsCodeClient) DeleteEntities(filePaths []string) error {

	req, err := retryablehttp.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.ApiClient.Endpoint, "/gateway/api/setup-as-code/yaml/delete-entities"), nil)
	if err != nil {
		return err
	}

	// Configure additional headers
	req.Header.Set(httphelpers.HeaderApiKey, c.ApiClient.APIKey)
	req.Header.Set(httphelpers.HeaderAuthorization, fmt.Sprintf("Bearer %s", c.ApiClient.BearerToken))
	req.Header.Set(httphelpers.HeaderAccept, httphelpers.HeaderApplicationJson)

	// Add the account ID to the query string
	q := req.URL.Query()
	q.Add(QueryParamAccountId, c.ApiClient.AccountId)
	q.Add(QueryParamFilePaths, strings.Join(filePaths, ","))
	req.URL.RawQuery = q.Encode()

	_, err = c.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigAsCodeClient) UpsertObject(input interface{}, filePath string, responseObj interface{}) error {
	if input == nil {
		return errors.New("object to upsert is nil")
	}

	// Check required fields are implemented
	requiredFields := []string{"Name", "Id"}
	for _, f := range requiredFields {
		if !utils.HasField(input, f) {
			return fmt.Errorf("obj must have field `%s`", f)
		}
	}

	// If the object implements the Validation interface then check it
	if v, ok := input.(cac.Validation); ok {
		if ok, err := v.Validate(); !ok {
			return err
		}
	}

	// Upsert the yaml document
	_, err := c.UpsertYamlEntity(filePath, input)
	if err != nil {
		return err
	}

	appId, ok := utils.TryGetFieldValue(input, "ApplicationId")
	if !ok {
		appId = ""
	}

	err = c.FindObject(appId.(string), filePath, responseObj)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigAsCodeClient) FindObject(applicationId string, filePath string, obj interface{}) error {
	rootItem, err := c.GetDirectoryTree(applicationId)
	if err != nil {
		return err
	}

	item := FindConfigAsCodeItemByPath(rootItem, filePath)
	if item == nil {
		return fmt.Errorf("unable to item at `%s`", filePath)
	}

	itemContent, err := c.GetDirectoryItemContent(item.RestName, item.UUID, applicationId)
	if err != nil {
		return err
	}

	// Parse the new entity
	err = itemContent.ParseYamlContent(obj)
	if err != nil {
		return err
	}

	// Parse the new entity
	err = itemContent.ParseYamlContent(obj)
	if err != nil {
		return err
	}

	// Set the required fields for all entities
	utils.TrySetField(obj, "Name", cac.GetEntityNameFromPath(filePath))
	utils.TrySetField(obj, "Id", item.UUID)

	if applicationId != "" && utils.HasField(obj, "ApplicationId") {
		utils.TrySetField(obj, "ApplicationId", applicationId)
	}

	return nil
}
