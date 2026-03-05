package nextgen

type SplunkAuthType string

var SplunkAuthTypes = struct {
	UsernamePassword SplunkAuthType
	BearerToken      SplunkAuthType
	HECToken         SplunkAuthType
	None             SplunkAuthType
}{
	UsernamePassword: "UsernamePassword",
	BearerToken:      "BearerToken",
	HECToken:         "HECToken",
	None:             "None",
}

var SplunkAuthTypeValues = []string{
	SplunkAuthTypes.UsernamePassword.String(),
	SplunkAuthTypes.BearerToken.String(),
	SplunkAuthTypes.HECToken.String(),
	SplunkAuthTypes.None.String(),
}

func (s SplunkAuthType) String() string {
	return string(s)
}
