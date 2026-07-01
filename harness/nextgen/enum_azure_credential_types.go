package nextgen

type AzureCredentialType string

var AzureCredentialTypes = struct {
	InheritFromDelegate AzureCredentialType
	ManualConfig        AzureCredentialType
	OidcAuthentication  AzureCredentialType
}{
	InheritFromDelegate: "InheritFromDelegate",
	ManualConfig:        "ManualConfig",
	OidcAuthentication:  "OidcAuthentication",
}

var AzureCredentialTypeValues = []string{
	AzureCredentialTypes.InheritFromDelegate.String(),
	AzureCredentialTypes.ManualConfig.String(),
	AzureCredentialTypes.OidcAuthentication.String(),
}

func (e AzureCredentialType) String() string {
	return string(e)
}
