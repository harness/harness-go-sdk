/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Notification Management Service
 *
 * API version: 1.0
 * Contact: contact.harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Response DTO for DefaultNotificationTemplateSet
type DefaultNotificationTemplateSetResponse struct {
	DefaultNotificationTemplateSet *DefaultNotificationTemplateSetResponseDto `json:"default_notification_template_set"`
	Created                        int64                                      `json:"created,omitempty"`
	Updated                        int64                                      `json:"updated,omitempty"`
}
