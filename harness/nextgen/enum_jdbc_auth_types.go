package nextgen

type JDBCAuthType string

var JDBCAuthTypes = struct {
	UsernamePassword JDBCAuthType
	ServiceAccount   JDBCAuthType
	KeyPair          JDBCAuthType
	Oidc             JDBCAuthType
}{
	UsernamePassword: "UsernamePassword",
	ServiceAccount:   "ServiceAccount",
	KeyPair:          "KeyPair",
	Oidc:             "Oidc",
}

func (e JDBCAuthType) String() string {
	return string(e)
}
