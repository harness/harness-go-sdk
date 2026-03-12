package nextgen

type SplunkAuthType string

var SplunkAuthTypes = struct {
	UsernamePassword SplunkAuthType
	BearerToken      SplunkAuthType
	HECToken         SplunkAuthType
	None             SplunkAuthType
}{
	UsernamePassword: "UsernamePassword",
	BearerToken:      "Bearer Token(HTTP Header)",
	HECToken:         "HEC Token",
	None:             "Anonymous",
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
