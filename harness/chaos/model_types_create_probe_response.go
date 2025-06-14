/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type TypesCreateProbeResponse struct {
	ApmProperties *ProbeApmProbe `json:"apmProperties,omitempty"`
	Description string `json:"description,omitempty"`
	K8sProperties *ProbeK8SProbe `json:"k8sProperties,omitempty"`
	KubernetesCmdProperties *ProbeKubernetesCmdProbe `json:"kubernetesCmdProperties,omitempty"`
	KubernetesDatadogProperties *ProbeDatadogProbe `json:"kubernetesDatadogProperties,omitempty"`
	KubernetesDynatraceProperties *ProbeDynatraceProbe `json:"kubernetesDynatraceProperties,omitempty"`
	KubernetesHttpProperties *ProbeHttpProbe `json:"kubernetesHttpProperties,omitempty"`
	LinuxCmdProperties *ProbeLinuxCmdProbe `json:"linuxCmdProperties,omitempty"`
	LinuxDatadogProperties *ProbeDatadogProbe `json:"linuxDatadogProperties,omitempty"`
	LinuxDynatraceProperties *ProbeDynatraceProbe `json:"linuxDynatraceProperties,omitempty"`
	LinuxHttpProperties *ProbeHttpProbe `json:"linuxHttpProperties,omitempty"`
	Name string `json:"name,omitempty"`
	ProbeId string `json:"probeId,omitempty"`
	PromProperties *ProbePromProbe `json:"promProperties,omitempty"`
	RevisionId string `json:"revisionId,omitempty"`
	SloProperties *ProbeSloProbe `json:"sloProperties,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Variables []TemplateVariable `json:"variables,omitempty"`
	WindowsHttpProperties *ProbeHttpProbe `json:"windowsHttpProperties,omitempty"`
}
