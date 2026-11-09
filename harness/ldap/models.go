package ldap

// Models for the spec-server v1 LDAP settings API. Field names mirror the
// snake_case JSON of io.harness.spec.server.ng.v1.model.LdapSettings*DTO in
// harness-core (120-ng-manager openapi v1 spec) — note in particular the
// bind_d_n / base_d_n spellings, which come straight from the spec.

type LdapSettingsDTO struct {
	AccountIdentifier      string                     `json:"account_identifier,omitempty"`
	Identifier             string                     `json:"identifier,omitempty"`
	Name                   string                     `json:"name,omitempty"`
	Url                    string                     `json:"url,omitempty"`
	Disabled               bool                       `json:"disabled,omitempty"`
	CronExpression         string                     `json:"cron_expression,omitempty"`
	SsoType                string                     `json:"sso_type,omitempty"`
	LdapConnectionSettings *LdapConnectionSettingsDTO `json:"ldap_connection_settings,omitempty"`
	LdapUserSettings       []LdapUserSettingsDTO      `json:"ldap_user_settings,omitempty"`
	LdapGroupSettings      []LdapGroupSettingsDTO     `json:"ldap_group_settings,omitempty"`
}

type LdapConnectionSettingsDTO struct {
	Host                              string   `json:"host,omitempty"`
	Port                              int32    `json:"port,omitempty"`
	SslEnabled                        bool     `json:"ssl_enabled,omitempty"`
	ReferralsEnabled                  *bool    `json:"referrals_enabled,omitempty"`
	MaxReferralHops                   *int32   `json:"max_referral_hops,omitempty"`
	ConnectionTimeout                 int32    `json:"connection_timeout,omitempty"`
	ResponseTimeout                   int32    `json:"response_timeout,omitempty"`
	UseRecursiveGroupMembershipSearch *bool    `json:"use_recursive_group_membership_search,omitempty"`
	BindDN                            string   `json:"bind_d_n,omitempty"`
	SecretRefPath                     string   `json:"secret_ref_path,omitempty"`
	DelegateSelectors                 []string `json:"delegate_selectors,omitempty"`
}

type LdapUserSettingsDTO struct {
	BaseDN              string `json:"base_d_n,omitempty"`
	SearchFilter        string `json:"search_filter,omitempty"`
	UidAttr             string `json:"uid_attr,omitempty"`
	SamAccountNameAttr  string `json:"sam_account_name_attr,omitempty"`
	EmailAttr           string `json:"email_attr,omitempty"`
	DisplayNameAttr     string `json:"display_name_attr,omitempty"`
	GroupMembershipAttr string `json:"group_membership_attr,omitempty"`
}

type LdapGroupSettingsDTO struct {
	BaseDN             string `json:"base_d_n,omitempty"`
	SearchFilter       string `json:"search_filter,omitempty"`
	NameAttr           string `json:"name_attr,omitempty"`
	DescriptionAttr    string `json:"description_attr,omitempty"`
	UserMembershipAttr string `json:"user_membership_attr,omitempty"`
	ReferencedUserAttr string `json:"referenced_user_attr,omitempty"`
}

type LdapSettingsRequest struct {
	LdapSettings *LdapSettingsDTO `json:"ldap_settings,omitempty"`
}

type LdapSettingsResponse struct {
	LdapSettings *LdapSettingsDTO `json:"ldap_settings,omitempty"`
	Created      *int64           `json:"created,omitempty"`
	LastModified *int64           `json:"last_modified,omitempty"`
}

type CronExpressionRequestDTO struct {
	Cron string `json:"cron,omitempty"`
}

type LdapSettingsIterations struct {
	Iterations []int64 `json:"iterations,omitempty"`
}

// ResponseMessage mirrors io.harness.eraro.ResponseMessage, the error body
// returned by the v1 API impl on 400/404 responses.
type ResponseMessage struct {
	Message string `json:"message,omitempty"`
}
