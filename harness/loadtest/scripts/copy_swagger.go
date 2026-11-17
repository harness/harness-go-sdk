//go:build ignore

// Copy the load-test operations this client calls out of the Load Test Manager
// swagger. The gateway prefix stays out of the spec. The provider adds
// /loadTest/manager/api the same way it adds /chaos/manager/api.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

// Operations this SDK slice calls. Anything else in the source spec stays there.
var keepMethods = map[string]map[string]struct{}{
	"/v1/load-tests": {
		"get": {}, "post": {},
	},
	"/v1/load-tests/{identity}": {
		"get": {}, "put": {}, "delete": {},
	},
	"/v1/load-tests/{identity}/variables": {
		"get": {},
	},
	"/v1/load-tests/{identity}/runs": {
		"get": {}, "post": {},
	},
	"/v1/load-tests/{identity}/sync-template": {
		"post": {},
	},
	"/v1/load-tests/from-template": {
		"post": {},
	},
	"/v1/runs/{identity}": {
		"get": {},
	},
	"/v1/runs/{identity}/stop": {
		"post": {},
	},
	"/v1/load-test-templates": {
		"post": {},
	},
	"/v1/load-test-templates/{identity}": {
		"get": {}, "put": {}, "delete": {},
	},
	"/v1/load-test-templates/{identity}/revisions": {
		"post": {},
	},
	"/v1/load-test-templates/{identity}/revisions/{revision}": {
		"get": {}, "delete": {},
	},
	"/v1/load-test-templates/{identity}/variables": {
		"get": {},
	},
	"/v1/load-test-templates/{identity}/referenced-by": {
		"get": {},
	},
	"/v1/composite-load-tests": {
		"post": {},
	},
	"/v1/composite-load-tests/{identifier}": {
		"get": {}, "put": {}, "delete": {},
	},
}

func main() {
	source := os.Getenv("LTM_SWAGGER")
	if source == "" {
		source = findSource()
	}
	if source == "" {
		exitf("loadTestManager/docs/swagger.yaml not found; set LTM_SWAGGER")
	}
	out := os.Getenv("LTM_SWAGGER_OUT")
	if out == "" {
		out = filepath.Join(moduleRoot(), "harness", "loadtest", "api", "swagger.yaml")
	}

	raw, err := os.ReadFile(source)
	if err != nil {
		exitf("read %s: %v", source, err)
	}
	var doc yaml.Node
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		exitf("parse %s: %v", source, err)
	}
	root := mappingRoot(&doc)

	paths := child(root, "paths")
	if paths == nil {
		exitf("source swagger has no paths")
	}
	kept := trimPaths(paths)
	ensureCreateScope(kept)
	flattenSingleAllOf(kept)
	defs := collectDefinitions(root, kept)
	flattenSingleAllOf(defs)
	markUpdateFieldsAlwaysSent(defs)

	trimmed := buildDocument(root, kept, defs)
	encoded, err := yaml.Marshal(trimmed)
	if err != nil {
		exitf("encode swagger: %v", err)
	}
	header := []byte("# Trimmed from the Load Test Manager swagger by harness/loadtest/scripts/copy_swagger.go.\n" +
		"# basePath is / because the source paths already start with /v1. Swagger codegen joins basePath and the path.\n" +
		"# The gateway prefix /loadTest/manager/api is not part of this spec. The provider adds it.\n" +
		"# ApiKeyAuth is the X-API-Key header the chaos auth middleware accepts. The source spec documents the UI bearer token.\n")
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		exitf("mkdir: %v", err)
	}
	if err := os.WriteFile(out, append(header, encoded...), 0o644); err != nil {
		exitf("write %s: %v", out, err)
	}
	fmt.Printf("wrote %s (%d paths, %d definitions) from %s\n", out, len(keepMethods), countKeys(defs), source)
}

func findSource() string {
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

func moduleRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		exitf("wd: %v", err)
	}
	for i := 0; i < 8; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	exitf("go.mod not found")
	return ""
}

func trimPaths(paths *yaml.Node) *yaml.Node {
	out := &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	for _, path := range sortedKeys(keepMethods) {
		node := child(paths, path)
		if node == nil {
			exitf("source swagger is missing %s", path)
		}
		ops := &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
		for _, method := range sortedKeys(keepMethods[path]) {
			op := child(node, method)
			if op == nil {
				exitf("source swagger is missing %s %s", method, path)
			}
			useAPIKey(op)
			ops.Content = append(ops.Content, scalar(method), op)
		}
		out.Content = append(out.Content, scalar(path), ops)
	}
	return out
}

func useAPIKey(op *yaml.Node) {
	security := child(op, "security")
	if security == nil {
		return
	}
	for _, item := range security.Content {
		if item.Kind != yaml.MappingNode {
			continue
		}
		for i := 0; i < len(item.Content); i += 2 {
			if item.Content[i].Value == "BearerAuth" {
				item.Content[i].Value = "ApiKeyAuth"
			}
		}
	}
}

func collectDefinitions(root, paths *yaml.Node) *yaml.Node {
	all := child(root, "definitions")
	if all == nil {
		exitf("source swagger has no definitions")
	}
	needed := map[string]struct{}{}
	collectRefs(paths, needed)
	seen := map[string]struct{}{}
	for {
		added := false
		for name := range needed {
			if _, ok := seen[name]; ok {
				continue
			}
			def := child(all, name)
			if def == nil {
				exitf("source swagger is missing definition %s", name)
			}
			seen[name] = struct{}{}
			before := len(needed)
			collectRefs(def, needed)
			if len(needed) > before {
				added = true
			}
		}
		if !added && len(seen) == len(needed) {
			break
		}
	}

	out := &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	names := sortedKeys(needed)
	sort.Strings(names)
	for _, name := range names {
		out.Content = append(out.Content, scalar(name), child(all, name))
	}
	return out
}

func ensureCreateScope(paths *yaml.Node) {
	// POST /v1/load-tests binds ScopeQuery and the RBAC middleware rejects a
	// missing accountIdentifier. The source annotation omits those query
	// parameters. Copy them from the get-by-identity operation, which lists
	// the same names.
	create := child(child(paths, "/v1/load-tests"), "post")
	if create == nil {
		exitf("trimmed swagger is missing POST /v1/load-tests")
	}
	if hasQuery(create, "accountIdentifier") {
		return
	}
	source := child(child(paths, "/v1/load-tests/{identity}"), "get")
	params := child(create, "parameters")
	if params == nil {
		params = &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
		add(create, "parameters", params)
	}
	for _, param := range child(source, "parameters").Content {
		if child(param, "in") != nil && child(param, "in").Value == "query" {
			params.Content = append(params.Content, param)
		}
	}
}

func hasQuery(op *yaml.Node, name string) bool {
	params := child(op, "parameters")
	if params == nil {
		return false
	}
	for _, param := range params.Content {
		if child(param, "in") != nil && child(param, "in").Value == "query" && child(param, "name") != nil && child(param, "name").Value == name {
			return true
		}
	}
	return false
}

// flattenSingleAllOf turns a one-ref allOf into a direct $ref. Swagger codegen
// emits an empty struct for allOf around a string enum, so cleanupPolicy would
// be sent as {} instead of "delete".
func flattenSingleAllOf(node *yaml.Node) {
	if node == nil {
		return
	}
	switch node.Kind {
	case yaml.DocumentNode, yaml.SequenceNode:
		for _, item := range node.Content {
			flattenSingleAllOf(item)
		}
	case yaml.MappingNode:
		allOf := child(node, "allOf")
		if allOf != nil && allOf.Kind == yaml.SequenceNode && len(allOf.Content) == 1 && child(allOf.Content[0], "$ref") != nil {
			ref := child(allOf.Content[0], "$ref")
			node.Content = []*yaml.Node{scalar("$ref"), ref}
			return
		}
		for i := 1; i < len(node.Content); i += 2 {
			flattenSingleAllOf(node.Content[i])
		}
	}
}

func markUpdateFieldsAlwaysSent(defs *yaml.Node) {
	// Swagger codegen adds omitempty to optional fields, which drops an empty
	// list. An empty tags or variables list is how those fields are cleared,
	// and resources {} clears CPU and memory. maxDurationSec and managedInfra
	// stay optional: a missing value must not become 0 or false. patch_generated.go
	// turns those two into pointers with no omitempty, so null leaves the stored value.
	def := child(defs, "internal_api.UpdateLoadTestRequest")
	if def == nil {
		exitf("trimmed swagger is missing internal_api.UpdateLoadTestRequest")
	}
	ensureRequired(def, "tags", "variables", "resources")
	template := child(defs, "internal_api.UpdateLoadTestTemplateRequest")
	if template == nil {
		exitf("trimmed swagger is missing internal_api.UpdateLoadTestTemplateRequest")
	}
	ensureRequired(template, "tags", "variables")
}

func ensureRequired(def *yaml.Node, names ...string) {
	required := child(def, "required")
	if required == nil {
		required = &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
		def.Content = append(def.Content, scalar("required"), required)
	}
	have := map[string]struct{}{}
	for _, item := range required.Content {
		have[item.Value] = struct{}{}
	}
	for _, name := range names {
		if _, ok := have[name]; ok {
			continue
		}
		required.Content = append(required.Content, scalar(name))
	}
}

func buildDocument(root, paths, defs *yaml.Node) *yaml.Node {
	info := child(root, "info")
	if info == nil {
		exitf("source swagger has no info")
	}
	doc := &yaml.Node{Kind: yaml.DocumentNode}
	body := &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	doc.Content = []*yaml.Node{body}
	add(body, "swagger", scalar("2.0"))
	add(body, "info", info)
	add(body, "basePath", scalar("/"))
	schemes := &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
	schemes.Content = []*yaml.Node{scalar("https"), scalar("http")}
	add(body, "schemes", schemes)
	add(body, "consumes", stringSeq("application/json"))
	add(body, "produces", stringSeq("application/json"))
	add(body, "securityDefinitions", apiKeySecurity())
	add(body, "security", securityRequirement())
	add(body, "paths", paths)
	add(body, "definitions", defs)
	return doc
}

func apiKeySecurity() *yaml.Node {
	def := &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	add(def, "type", scalar("apiKey"))
	add(def, "in", scalar("header"))
	add(def, "name", scalar("X-API-Key"))
	add(def, "description", scalar("Harness API key. Load Test Manager accepts the chaos auth middleware header."))
	out := &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	add(out, "ApiKeyAuth", def)
	return out
}

func securityRequirement() *yaml.Node {
	item := &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	add(item, "ApiKeyAuth", &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"})
	out := &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
	out.Content = []*yaml.Node{item}
	return out
}

func collectRefs(node *yaml.Node, refs map[string]struct{}) {
	if node == nil {
		return
	}
	switch node.Kind {
	case yaml.DocumentNode, yaml.SequenceNode:
		for _, child := range node.Content {
			collectRefs(child, refs)
		}
	case yaml.MappingNode:
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			val := node.Content[i+1]
			if key.Value == "$ref" && val.Value != "" {
				const prefix = "#/definitions/"
				if len(val.Value) > len(prefix) && val.Value[:len(prefix)] == prefix {
					refs[val.Value[len(prefix):]] = struct{}{}
				}
				continue
			}
			collectRefs(val, refs)
		}
	}
}

func mappingRoot(doc *yaml.Node) *yaml.Node {
	if doc.Kind == yaml.DocumentNode && len(doc.Content) == 1 {
		return doc.Content[0]
	}
	return doc
}

func child(mapping *yaml.Node, key string) *yaml.Node {
	if mapping == nil || mapping.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i < len(mapping.Content); i += 2 {
		if mapping.Content[i].Value == key {
			return mapping.Content[i+1]
		}
	}
	return nil
}

func add(mapping *yaml.Node, key string, val *yaml.Node) {
	mapping.Content = append(mapping.Content, scalar(key), val)
}

func scalar(value string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: value}
}

func stringSeq(values ...string) *yaml.Node {
	out := &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
	for _, value := range values {
		out.Content = append(out.Content, scalar(value))
	}
	return out
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func countKeys(mapping *yaml.Node) int {
	if mapping == nil {
		return 0
	}
	return len(mapping.Content) / 2
}

func exitf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
