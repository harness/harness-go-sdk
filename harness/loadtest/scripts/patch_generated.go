//go:build ignore

// Reapply the provider fields after swagger-codegen. The generated client has no
// account id or API key. Service discovery's client carries both, and the
// provider sets them the same way.
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	patch("configuration.go",
		"type Configuration struct {\n",
		"type Configuration struct {\n\tAccountId     string            `json:\"accountId,omitempty\"`\n\tApiKey        string            `json:\"apiKey,omitempty\"`\n",
		"AccountId     string")
	patch("client.go",
		"\tcommon service // Reuse a single struct instead of allocating one for each service on the heap.\n\n",
		"\tcommon service // Reuse a single struct instead of allocating one for each service on the heap.\n\n\tAccountId string\n\tApiKey    string\n\n",
		"\tAccountId string\n")
	patch("client.go",
		"\tc.cfg = cfg\n\tc.common.client = c\n\n",
		"\tc.cfg = cfg\n\tc.common.client = c\n\tc.AccountId = cfg.AccountId\n\tc.ApiKey = cfg.ApiKey\n\n",
		"c.AccountId = cfg.AccountId")
	patchUpdatePointers()
}

// A plain int32 or bool cannot represent "leave the stored value". Null does.
// 0 still clears the duration cap and false still turns managed infrastructure off.
func patchUpdatePointers() {
	path := "model_internal_api_update_load_test_request.go"
	raw, err := os.ReadFile(path)
	if err != nil {
		exitf("read %s: %v", path, err)
	}
	text := string(raw)
	text = replaceField(text, "ManagedInfra", "managedInfra", "\tManagedInfra *bool `json:\"managedInfra\"`")
	text = replaceField(text, "MaxDurationSec", "maxDurationSec", "\tMaxDurationSec *int32 `json:\"maxDurationSec\"`")
	if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
		exitf("write %s: %v", path, err)
	}
	fmt.Printf("patched %s pointers\n", path)
}

func replaceField(text, field, jsonName, replacement string) string {
	pattern := regexp.MustCompile("(?m)^[ \\t]*" + field + "\\s+\\*?[A-Za-z0-9_.]+\\s+`json:\"" + jsonName + "(?:,omitempty)?\"`")
	if !pattern.MatchString(text) {
		exitf("update model is missing %s", field)
	}
	return pattern.ReplaceAllString(text, replacement)
}

func patch(path, old, new, already string) {
	raw, err := os.ReadFile(path)
	if err != nil {
		exitf("read %s: %v", path, err)
	}
	text := string(raw)
	if strings.Contains(text, already) {
		return
	}
	if !strings.Contains(text, old) {
		exitf("%s is missing the generated text this patch expects", path)
	}
	if err := os.WriteFile(path, []byte(strings.Replace(text, old, new, 1)), 0o644); err != nil {
		exitf("write %s: %v", path, err)
	}
	fmt.Printf("patched %s\n", path)
}

func exitf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
