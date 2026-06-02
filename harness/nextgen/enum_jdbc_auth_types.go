package nextgen

type JDBCAuthType string

var JDBCAuthTypes = struct {
	UsernamePassword    JDBCAuthType
	ServiceAccount      JDBCAuthType
	KeyPair             JDBCAuthType
	InheritFromDelegate JDBCAuthType
	Oidc                JDBCAuthType
}{
	UsernamePassword:    "UsernamePassword",
	ServiceAccount:      "ServiceAccount",
	KeyPair:             "KeyPair",
	InheritFromDelegate: "InheritFromDelegate",
	Oidc:                "Oidc",
}

func (e JDBCAuthType) String() string {
	return string(e)
}
