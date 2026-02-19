package nextgen

type JDBCAuthType string

var JDBCAuthTypes = struct {
	UsernamePassword JDBCAuthType
	ServiceAccount   JDBCAuthType
	KeyPair          JDBCAuthType
}{
	UsernamePassword: "UsernamePassword",
	ServiceAccount:   "ServiceAccount",
	KeyPair:          "KeyPair",
}

func (e JDBCAuthType) String() string {
	return string(e)
}
