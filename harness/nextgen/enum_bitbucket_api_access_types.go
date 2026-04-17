package nextgen

type BitBucketApiAccessType string

var BitBucketApiAccessTypes = struct {
	UsernameToken    BitBucketApiAccessType
	AccessToken      BitBucketApiAccessType
	EmailAndApiToken BitBucketApiAccessType
}{
	UsernameToken:    "UsernameToken",
	AccessToken:      "AccessToken",
	EmailAndApiToken: "EmailAndApiToken",
}

var BitBucketApiAccessTypeValues = []string{
	BitBucketApiAccessTypes.UsernameToken.String(),
	BitBucketApiAccessTypes.AccessToken.String(),
	BitBucketApiAccessTypes.EmailAndApiToken.String(),
}

func (e BitBucketApiAccessType) String() string {
	return string(e)
}
