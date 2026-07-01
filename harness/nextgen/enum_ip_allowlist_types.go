package nextgen

type AllowedSourceType string

var AllowedSourceTypes = struct {
	UIAllowdSourceType   AllowedSourceType
	APIAllowedSourceType AllowedSourceType
}{
	UIAllowdSourceType:   "UI",
	APIAllowedSourceType: "API",
}

var AllowedSourceTypeValues = []string{
	AllowedSourceTypes.UIAllowdSourceType.String(),
	AllowedSourceTypes.APIAllowedSourceType.String(),
}

func (e AllowedSourceType) String() string {
	return string(e)
}
