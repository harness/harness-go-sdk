package mesh

// Service IDs matching Java AuthorizationServiceHeader.getServiceId().
// Last segment of a SPIFFE path must equal one of these for principal mapping.
const (
	ServiceAnalyzerService              = "AnalyzerService"
	ServiceAnnotations                  = "Annotations"
	ServiceBearer                       = "Bearer"
	ServiceManager                      = "Manager"
	ServiceNextGenManager               = "NextGenManager"
	ServiceBatchProcessing              = "BatchProcessing"
	ServiceAuditEventStreaming          = "AuditEventStreaming"
	ServiceCIManager                    = "CIManager"
	ServiceSTOManager                   = "STOManager"
	ServiceCVNextGen                    = "CVNextGen"
	ServiceCENextGen                    = "CENextGen"
	ServiceCustomDashboards             = "CustomDashboards"
	ServiceDelegateService              = "DelegateService"
	ServiceIdentityService              = "IdentityService"
	ServiceAdminPortal                  = "AdminPortal"
	ServiceNotificationService          = "NotificationService"
	ServiceAuditService                 = "AuditService"
	ServiceAlertService                 = "AlertService"
	ServiceStreamingService             = "StreamingService"
	ServicePipelineService              = "PipelineService"
	ServiceTemplateService              = "TemplateService"
	ServiceAccessControlService         = "accessControlService"
	ServiceResourceGroupService         = "ResourceGroupService"
	ServicePlatformService              = "PlatformService"
	ServiceGitSyncService               = "GitSyncService"
	ServiceDefault                      = "Default"
	ServiceDashboardAggregationService  = "DashboardAggregationService"
	ServiceDelegateManagementService    = "DelegateManagementService"
	ServiceDebeziumService              = "DebeziumService"
	ServiceSubscriptionService          = "SubscriptionService"
	ServiceChaosService                 = "ChaosService"
	ServiceLoadTestService              = "LoadTestService"
	ServiceServiceDiscoveryService      = "ServiceDiscoveryService"
	ServiceMonitoringManagerService     = "MonitoringManagerService"
	ServiceCode                         = "Code"
	ServiceIACMManager                  = "IACMManager"
	ServiceMigratorService              = "MigratorService"
	ServiceIDPService                   = "IDPService"
	ServiceSSCAService                  = "SSCAService"
	ServiceComponentService             = "ComponentService"
	ServiceComponentAnalysisService     = "ComponentAnalysisService"
	ServiceIDPUi                        = "IDPUi"
	ServiceDBDevopsService              = "DBDevopsService"
	ServiceIROManagerService            = "IROManagerService"
	ServiceRelicx                       = "Relicx"
	ServiceLicenseManager               = "LicenseManager"
	ServiceCDEManager                   = "CDEManager"
	ServiceGitopsService                = "GitopsService"
	ServiceAppsecService                = "AppsecService"
	ServiceResourceHierarchyService     = "ResourceHierarchyService"
	ServicePlatformConfigService        = "PlatformConfigService"
	ServiceSecretConnectorService       = "SecretConnectorService"
	ServiceVersionManagementService     = "VersionManagementService"
	ServiceHarnessStatuspage            = "HarnessStatuspage"
)

// KnownServiceIDs is the allow-list used by principal mapping and mint-side name derivation.
var KnownServiceIDs = map[string]struct{}{
	ServiceAnalyzerService:             {},
	ServiceAnnotations:                 {},
	ServiceBearer:                      {},
	ServiceManager:                     {},
	ServiceNextGenManager:              {},
	ServiceBatchProcessing:             {},
	ServiceAuditEventStreaming:         {},
	ServiceCIManager:                   {},
	ServiceSTOManager:                  {},
	ServiceCVNextGen:                   {},
	ServiceCENextGen:                   {},
	ServiceCustomDashboards:            {},
	ServiceDelegateService:             {},
	ServiceIdentityService:             {},
	ServiceAdminPortal:                 {},
	ServiceNotificationService:         {},
	ServiceAuditService:                {},
	ServiceAlertService:                {},
	ServiceStreamingService:            {},
	ServicePipelineService:             {},
	ServiceTemplateService:             {},
	ServiceAccessControlService:        {},
	ServiceResourceGroupService:        {},
	ServicePlatformService:             {},
	ServiceGitSyncService:              {},
	ServiceDefault:                     {},
	ServiceDashboardAggregationService: {},
	ServiceDelegateManagementService:   {},
	ServiceDebeziumService:             {},
	ServiceSubscriptionService:         {},
	ServiceChaosService:                {},
	ServiceLoadTestService:             {},
	ServiceServiceDiscoveryService:     {},
	ServiceMonitoringManagerService:    {},
	ServiceCode:                        {},
	ServiceIACMManager:                 {},
	ServiceMigratorService:             {},
	ServiceIDPService:                  {},
	ServiceSSCAService:                 {},
	ServiceComponentService:            {},
	ServiceComponentAnalysisService:    {},
	ServiceIDPUi:                       {},
	ServiceDBDevopsService:             {},
	ServiceIROManagerService:           {},
	ServiceRelicx:                      {},
	ServiceLicenseManager:              {},
	ServiceCDEManager:                  {},
	ServiceGitopsService:               {},
	ServiceAppsecService:               {},
	ServiceResourceHierarchyService:    {},
	ServicePlatformConfigService:       {},
	ServiceSecretConnectorService:      {},
	ServiceVersionManagementService:    {},
	ServiceHarnessStatuspage:           {},
}

// IsKnownServiceID reports whether id is a registered AuthorizationServiceHeader value.
func IsKnownServiceID(id string) bool {
	_, ok := KnownServiceIDs[id]
	return ok
}
