/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type EnumWebhookTrigger string

// List of EnumWebhookTrigger
const (
	BRANCH_CREATED_EnumWebhookTrigger          EnumWebhookTrigger = "branch_created"
	BRANCH_DELETED_EnumWebhookTrigger          EnumWebhookTrigger = "branch_deleted"
	BRANCH_UPDATED_EnumWebhookTrigger          EnumWebhookTrigger = "branch_updated"
	PULLREQ_BRANCH_UPDATED_EnumWebhookTrigger  EnumWebhookTrigger = "pullreq_branch_updated"
	PULLREQ_CLOSED_EnumWebhookTrigger          EnumWebhookTrigger = "pullreq_closed"
	PULLREQ_COMMENT_CREATED_EnumWebhookTrigger EnumWebhookTrigger = "pullreq_comment_created"
	PULLREQ_CREATED_EnumWebhookTrigger         EnumWebhookTrigger = "pullreq_created"
	PULLREQ_MERGED_EnumWebhookTrigger          EnumWebhookTrigger = "pullreq_merged"
	PULLREQ_REOPENED_EnumWebhookTrigger        EnumWebhookTrigger = "pullreq_reopened"
	TAG_CREATED_EnumWebhookTrigger             EnumWebhookTrigger = "tag_created"
	TAG_DELETED_EnumWebhookTrigger             EnumWebhookTrigger = "tag_deleted"
	TAG_UPDATED_EnumWebhookTrigger             EnumWebhookTrigger = "tag_updated"
)
