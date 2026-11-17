package loadtest

// ModelMap is a free-form JSON object. Swagger codegen names that schema "map",
// which is a Go keyword, and then does not emit the renamed type.
type ModelMap map[string]interface{}
