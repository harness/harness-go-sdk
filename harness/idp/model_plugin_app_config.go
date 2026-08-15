package idp

// PluginAppConfigRequest is the request body for the save-or-update plugin app config endpoint.
// All fields are wrapped under the "app_config" key when sent to the API.
type PluginAppConfigRequest struct {
	AppConfig PluginAppConfig `json:"app_config"`
}

// PluginAppConfig represents the configuration for an IDP plugin.
type PluginAppConfig struct {
	ConfigId     string                   `json:"config_id"`
	ConfigName   string                   `json:"config_name"`
	Enabled      bool                     `json:"enabled"`
	Configs      string                   `json:"configs"`
	EnvVariables []PluginAppConfigEnvVar  `json:"env_variables"`
	Proxy        []PluginAppConfigProxy   `json:"proxy"`
}

// PluginAppConfigEnvVar represents an environment variable injected into the plugin runtime.
type PluginAppConfigEnvVar struct {
	Identifier              string `json:"identifier,omitempty"`
	EnvName                 string `json:"env_name"`
	Type                    string `json:"type"`
	HarnessSecretIdentifier string `json:"harness_secret_identifier"`
	Created                 int64  `json:"created,omitempty"`
	Updated                 int64  `json:"updated,omitempty"`
	IsDeleted               bool   `json:"is_deleted,omitempty"`
}

// PluginAppConfigProxy represents a proxy configuration for outbound plugin HTTP calls.
type PluginAppConfigProxy struct {
	Host            string   `json:"host"`
	Proxy           bool     `json:"proxy"`
	Selectors       []string `json:"selectors,omitempty"`
	Identifier      string   `json:"identifier,omitempty"`
	PluginId        string   `json:"pluginId,omitempty"`
	HealthCheckPath *string  `json:"healthCheckPath,omitempty"`
}

// PluginAppConfigResponse is the API response from save-or-update and toggle endpoints.
type PluginAppConfigResponse struct {
	AppConfig *PluginAppConfigResponseData `json:"app_config,omitempty"`
}

// PluginAppConfigResponseData contains the response data fields returned by the API.
type PluginAppConfigResponseData struct {
	ConfigId          string                   `json:"config_id"`
	ConfigName        string                   `json:"config_name"`
	Configs           string                   `json:"configs"`
	Enabled           bool                     `json:"enabled"`
	EnabledDisabledAt int64                    `json:"enabled_disabled_at,omitempty"`
	Created           int64                    `json:"created,omitempty"`
	Updated           int64                    `json:"updated,omitempty"`
	EnvVariables      []PluginAppConfigEnvVar  `json:"env_variables"`
	Proxy             []PluginAppConfigProxy   `json:"proxy"`
}

// PluginInfoResponse is the API response from the get-plugin-info endpoint.
type PluginInfoResponse struct {
	Plugin *PluginInfoData `json:"plugin,omitempty"`
}

// PluginInfoData contains plugin details, config, and metadata.
type PluginInfoData struct {
	PluginDetails *PluginDetails           `json:"plugin_details,omitempty"`
	Config        *string                  `json:"config"`
	Exports       *PluginExports           `json:"exports,omitempty"`
	EnvVariables  []PluginAppConfigEnvVar  `json:"env_variables"`
	Proxy         []PluginAppConfigProxy   `json:"proxy"`
	Saved         bool                     `json:"saved"`
}

// PluginDetails contains plugin metadata.
type PluginDetails struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	CreatedBy     string  `json:"created_by,omitempty"`
	IconUrl       string  `json:"icon_url,omitempty"`
	Description   string  `json:"description,omitempty"`
	Category      string  `json:"category,omitempty"`
	Source        string  `json:"source,omitempty"`
	ImageUrl      string  `json:"image_url,omitempty"`
	Documentation string  `json:"documentation,omitempty"`
	Core          bool    `json:"core"`
	Enabled       bool    `json:"enabled"`
	PluginType    string  `json:"plugin_type,omitempty"`
}

// PluginExports contains export counts for the plugin.
type PluginExports struct {
	Pages              int           `json:"pages"`
	TabContents        int           `json:"tab_contents"`
	Cards              int           `json:"cards"`
	DefaultEntityTypes []string      `json:"default_entity_types"`
	ExportDetails      []any `json:"export_details"`
}

// PluginListItem represents a single plugin in the list-all-plugins response.
type PluginListItem struct {
	Plugin *PluginListDetails `json:"plugin,omitempty"`
}

// PluginListDetails contains the plugin details returned in list responses.
type PluginListDetails struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	CreatedBy     string   `json:"created_by,omitempty"`
	IconUrl       string   `json:"icon_url,omitempty"`
	Description   string   `json:"description,omitempty"`
	Category      string   `json:"category,omitempty"`
	Source        string   `json:"source,omitempty"`
	ImageUrl      string   `json:"image_url,omitempty"`
	Images        []string `json:"images"`
	Documentation string   `json:"documentation,omitempty"`
	Core          bool     `json:"core"`
	Enabled       bool     `json:"enabled"`
	PluginType    string   `json:"plugin_type,omitempty"`
}
