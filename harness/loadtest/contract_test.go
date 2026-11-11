package loadtest

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLoadTestContractMatchesLTMSwagger(t *testing.T) {
	path := ltmSwaggerPath()
	if path == "" {
		t.Skip("loadTestManager swagger.yaml not found; set LTM_SWAGGER or check out loadTestManager next to harness-go-sdk")
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var doc swaggerDoc
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}

	assertRequiredCovered(t, doc, "internal_api.CreateLoadTestRequest", LoadTestCreateRequest{})
	assertRequiredCovered(t, doc, "internal_api.CreateLoadTestRunRequest", LoadTestRunCreateRequest{})
	// Request bodies must be able to send every field the spec defines. A new
	// response field is ignored by the decoder, so a newer Load Test Manager
	// spec does not fail this test until a path, query, or request field changes.
	for _, model := range []struct {
		definition string
		value      interface{}
		request    bool
	}{
		{"internal_api.CreateLoadTestRequest", LoadTestCreateRequest{}, true},
		{"internal_api.UpdateLoadTestRequest", LoadTestUpdateRequest{}, true},
		{"internal_api.CreateLoadTestRunRequest", LoadTestRunCreateRequest{}, true},
		{"internal_api.LoadTestResponse", LoadTest{}, false},
		{"internal_api.PaginatedLoadTestListResponse", LoadTestListResponse{}, false},
		{"internal_api.SuccessResponse", LoadTestSuccessResponse{}, false},
		{"internal_api.LoadTestRunResponse", LoadTestRun{}, false},
	} {
		assertJSONFields(t, doc, model.definition, model.value)
		if model.request {
			assertSwaggerPropertiesPresent(t, doc, model.definition, model.value)
		}
	}

	assertEnumCovers(t, doc, "internal_api.LoadTestToolType",
		LoadTestToolLocust, LoadTestToolJMeter, LoadTestToolK6, LoadTestToolGatling, LoadTestToolCustom)
}

// Operations this client generates. POST /v1/load-tests is missing its scope
// queries in the source annotation; the vendored copy adds the same names the
// other operations already list.
var loadTestOperations = map[string][]string{
	"/v1/load-tests":                      {"get", "post"},
	"/v1/load-tests/{identity}":           {"delete", "get", "put"},
	"/v1/load-tests/{identity}/runs":      {"get", "post"},
	"/v1/load-tests/{identity}/variables": {"get"},
	"/v1/runs/{identity}":                 {"get"},
	"/v1/runs/{identity}/stop":            {"post"},
}

var scopeQueryNames = map[string]struct{}{
	"accountIdentifier":      {},
	"organizationIdentifier": {},
	"projectIdentifier":      {},
	"parentUniqueId":         {},
}

func TestVendoredOperationsMatchLTMSwagger(t *testing.T) {
	sourcePath := ltmSwaggerPath()
	if sourcePath == "" {
		t.Skip("loadTestManager swagger.yaml not found; set LTM_SWAGGER or check out loadTestManager next to harness-go-sdk")
	}
	source := readSwagger(t, sourcePath)
	vendored := readSwagger(t, filepath.Join("api", "swagger.yaml"))
	if vendored.BasePath != "/" {
		t.Errorf("vendored basePath = %q, want /", vendored.BasePath)
	}
	if len(vendored.Paths) != len(loadTestOperations) {
		t.Errorf("vendored paths = %v", pathKeys(vendored.Paths))
	}
	for path, methods := range loadTestOperations {
		sourceOps, ok := source.Paths[path]
		if !ok {
			t.Errorf("source swagger is missing %s", path)
			continue
		}
		vendoredOps, ok := vendored.Paths[path]
		if !ok {
			t.Errorf("vendored swagger is missing %s", path)
			continue
		}
		for _, method := range methods {
			sourceOp, ok := sourceOps[method]
			if !ok {
				t.Errorf("source swagger is missing %s %s", strings.ToUpper(method), path)
				continue
			}
			vendoredOp, ok := vendoredOps[method]
			if !ok {
				t.Errorf("vendored swagger is missing %s %s", strings.ToUpper(method), path)
				continue
			}
			assertQueryNames(t, path, method, queryNames(sourceOp), queryNames(vendoredOp))
		}
		for method := range vendoredOps {
			if !containsMethod(methods, method) {
				t.Errorf("vendored swagger has extra %s %s", strings.ToUpper(method), path)
			}
		}
	}
	for path := range vendored.Paths {
		if _, ok := loadTestOperations[path]; !ok {
			t.Errorf("vendored swagger has extra path %s", path)
		}
	}
}

func readSwagger(t *testing.T, path string) swaggerDoc {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var doc swaggerDoc
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	return doc
}

func assertQueryNames(t *testing.T, path, method string, source, vendored map[string]struct{}) {
	t.Helper()
	for name := range source {
		if _, ok := vendored[name]; !ok {
			t.Errorf("%s %s query %q is in the Load Test Manager spec and missing from the vendored spec", strings.ToUpper(method), path, name)
		}
	}
	for name := range vendored {
		if _, ok := source[name]; ok {
			continue
		}
		if _, ok := scopeQueryNames[name]; ok {
			continue
		}
		t.Errorf("%s %s query %q is in the vendored spec and not in the Load Test Manager spec", strings.ToUpper(method), path, name)
	}
}

func queryNames(op swaggerOperation) map[string]struct{} {
	names := map[string]struct{}{}
	for _, param := range op.Parameters {
		if param.In == "query" && param.Name != "" {
			names[param.Name] = struct{}{}
		}
	}
	return names
}

func containsMethod(methods []string, method string) bool {
	for _, item := range methods {
		if item == method {
			return true
		}
	}
	return false
}

func pathKeys(paths map[string]map[string]swaggerOperation) []string {
	keys := make([]string, 0, len(paths))
	for key := range paths {
		keys = append(keys, key)
	}
	return keys
}

func ltmSwaggerPath() string {
	if path := os.Getenv("LTM_SWAGGER"); path != "" {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for i := 0; i < 8; i++ {
		path := filepath.Join(dir, "loadTestManager", "docs", "swagger.yaml")
		if _, err := os.Stat(path); err == nil {
			return path
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

type swaggerDoc struct {
	BasePath    string                                 `yaml:"basePath"`
	Paths       map[string]map[string]swaggerOperation `yaml:"paths"`
	Definitions map[string]swaggerDefinition           `yaml:"definitions"`
}

type swaggerOperation struct {
	Parameters []swaggerParameter `yaml:"parameters"`
}

type swaggerParameter struct {
	Name string `yaml:"name"`
	In   string `yaml:"in"`
}

type swaggerDefinition struct {
	Required   []string                   `yaml:"required"`
	Properties map[string]swaggerProperty `yaml:"properties"`
	Enum       []string                   `yaml:"enum"`
}

type swaggerProperty struct {
	Ref string `yaml:"$ref"`
}

func assertJSONFields(t *testing.T, doc swaggerDoc, definition string, value interface{}) {
	t.Helper()
	def, ok := doc.Definitions[definition]
	if !ok {
		t.Errorf("swagger definition %s is missing", definition)
		return
	}
	for _, name := range jsonFieldNames(value) {
		if _, ok := def.Properties[name]; !ok {
			t.Errorf("SDK field %q is not a property of swagger %s", name, definition)
		}
	}
}

func assertSwaggerPropertiesPresent(t *testing.T, doc swaggerDoc, definition string, value interface{}) {
	t.Helper()
	def, ok := doc.Definitions[definition]
	if !ok {
		t.Errorf("swagger definition %s is missing", definition)
		return
	}
	have := map[string]struct{}{}
	for _, name := range jsonFieldNames(value) {
		have[name] = struct{}{}
	}
	for name := range def.Properties {
		if _, ok := have[name]; !ok {
			t.Errorf("swagger %s property %q is missing from the SDK type", definition, name)
		}
	}
}

func assertRequiredCovered(t *testing.T, doc swaggerDoc, definition string, value interface{}) {
	t.Helper()
	def, ok := doc.Definitions[definition]
	if !ok {
		t.Errorf("swagger definition %s is missing", definition)
		return
	}
	have := map[string]struct{}{}
	for _, name := range jsonFieldNames(value) {
		have[name] = struct{}{}
	}
	for _, name := range def.Required {
		if _, ok := have[name]; !ok {
			t.Errorf("swagger %s requires %q and the SDK body cannot send it", definition, name)
		}
	}
}

func assertEnumCovers(t *testing.T, doc swaggerDoc, definition string, values ...string) {
	t.Helper()
	def, ok := doc.Definitions[definition]
	if !ok {
		t.Errorf("swagger definition %s is missing", definition)
		return
	}
	have := map[string]struct{}{}
	for _, value := range def.Enum {
		have[value] = struct{}{}
	}
	for _, value := range values {
		if _, ok := have[value]; !ok {
			t.Errorf("SDK value %q is not in swagger %s enum %v", value, definition, def.Enum)
		}
	}
}

func jsonFieldNames(value interface{}) []string {
	fieldType := reflect.TypeOf(value)
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}
	names := make([]string, 0, fieldType.NumField())
	for i := 0; i < fieldType.NumField(); i++ {
		tag := fieldType.Field(i).Tag.Get("json")
		name := strings.Split(tag, ",")[0]
		if name == "" || name == "-" {
			continue
		}
		names = append(names, name)
	}
	return names
}
