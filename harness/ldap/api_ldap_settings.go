package ldap

import (
	"context"
	"net/http"
)

// LdapSettingsApiService wraps the /v1/ldap-settings endpoints.
//
// The backend stores exactly one LDAP configuration per account; every
// endpoint resolves the configuration from the Harness-Account header. The
// {identifier} path parameter accepted by some endpoints is required by the
// spec but ignored server-side (see NGLdapSettingsApiImpl in harness-core) —
// pass the configuration's identifier for forward compatibility.
type LdapSettingsApiService struct {
	client *APIClient
}

// CreateLdapSettings creates the account's LDAP configuration.
// Returns 400 "Ldap settings already exist for this account." when one exists.
func (s *LdapSettingsApiService) CreateLdapSettings(ctx context.Context, settings LdapSettingsDTO) (LdapSettingsResponse, *http.Response, error) {
	var out LdapSettingsResponse
	resp, err := s.client.do(ctx, http.MethodPost, "/v1/ldap-settings", LdapSettingsRequest{LdapSettings: &settings}, &out)
	return out, resp, err
}

// GetLdapSettings returns the account's LDAP configuration, or 404 when none exists.
func (s *LdapSettingsApiService) GetLdapSettings(ctx context.Context) (LdapSettingsResponse, *http.Response, error) {
	var out LdapSettingsResponse
	resp, err := s.client.do(ctx, http.MethodGet, "/v1/ldap-settings", nil, &out)
	return out, resp, err
}

// GetLdapSettingsByIdentifier returns the account's LDAP configuration via the
// {identifier} path variant (identifier ignored server-side).
func (s *LdapSettingsApiService) GetLdapSettingsByIdentifier(ctx context.Context, identifier string) (LdapSettingsResponse, *http.Response, error) {
	var out LdapSettingsResponse
	resp, err := s.client.do(ctx, http.MethodGet, "/v1/ldap-settings/"+identifier, nil, &out)
	return out, resp, err
}

// UpdateLdapSettings replaces the account's LDAP configuration (full-body PUT).
func (s *LdapSettingsApiService) UpdateLdapSettings(ctx context.Context, identifier string, settings LdapSettingsDTO) (LdapSettingsResponse, *http.Response, error) {
	var out LdapSettingsResponse
	resp, err := s.client.do(ctx, http.MethodPut, "/v1/ldap-settings/"+identifier, LdapSettingsRequest{LdapSettings: &settings}, &out)
	return out, resp, err
}

// DeleteLdapSettings deletes the account's LDAP configuration. Returns 404 when
// none exists and 400 when user groups are still linked to the configuration.
func (s *LdapSettingsApiService) DeleteLdapSettings(ctx context.Context, identifier string) (*http.Response, error) {
	return s.client.do(ctx, http.MethodDelete, "/v1/ldap-settings/"+identifier, nil, nil)
}

// GetIterations previews the next run times for a cron expression.
func (s *LdapSettingsApiService) GetIterations(ctx context.Context, cron string) (LdapSettingsIterations, *http.Response, error) {
	var out LdapSettingsIterations
	resp, err := s.client.do(ctx, http.MethodPost, "/v1/ldap-settings/iterations", CronExpressionRequestDTO{Cron: cron}, &out)
	return out, resp, err
}
