/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * API version: 3.0
 * Contact: contact@harness.io
 */
package nextgen

type UserGroupGovernanceMetadata struct {
	Id        string                           `json:"id,omitempty"`
	Deny      bool                             `json:"deny,omitempty"`
	Details   []UserGroupPolicySetMetadata     `json:"details,omitempty"`
	Message   string                           `json:"message,omitempty"`
	Timestamp string                           `json:"timestamp,omitempty"`
	Status    string                           `json:"status,omitempty"`
	AccountId string                           `json:"accountId,omitempty"`
	OrgId     string                           `json:"orgId,omitempty"`
	ProjectId string                           `json:"projectId,omitempty"`
	Entity    string                           `json:"entity,omitempty"`
	Type      string                           `json:"type,omitempty"`
	Action    string                           `json:"action,omitempty"`
	Created   string                           `json:"created,omitempty"`
}

type UserGroupPolicySetMetadata struct {
	PolicySetId   string                       `json:"policySetId,omitempty"`
	Deny          bool                         `json:"deny,omitempty"`
	PolicyMetadata []UserGroupPolicyMetadata   `json:"policyMetadata,omitempty"`
	PolicySetName string                       `json:"policySetName,omitempty"`
	Status        string                       `json:"status,omitempty"`
	Identifier    string                       `json:"identifier,omitempty"`
	Created       string                       `json:"created,omitempty"`
	AccountId     string                       `json:"accountId,omitempty"`
	OrgId         string                       `json:"orgId,omitempty"`
	ProjectId     string                       `json:"projectId,omitempty"`
	Description   string                       `json:"description,omitempty"`
}

type UserGroupPolicyMetadata struct {
	PolicyId     string   `json:"policyId,omitempty"`
	PolicyName   string   `json:"policyName,omitempty"`
	Severity     string   `json:"severity,omitempty"`
	DenyMessages []string `json:"denyMessages,omitempty"`
	Status       string   `json:"status,omitempty"`
	Identifier   string   `json:"identifier,omitempty"`
	AccountId    string   `json:"accountId,omitempty"`
	OrgId        string   `json:"orgId,omitempty"`
	ProjectId    string   `json:"projectId,omitempty"`
	Created      string   `json:"created,omitempty"`
	Updated      string   `json:"updated,omitempty"`
	Error        string   `json:"error,omitempty"`
}
