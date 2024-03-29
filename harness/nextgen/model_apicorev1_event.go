/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type Apicorev1Event struct {
	Metadata           *V1ObjectMeta      `json:"metadata,omitempty"`
	InvolvedObject     *V1ObjectReference `json:"involvedObject,omitempty"`
	Reason             string             `json:"reason,omitempty"`
	Message            string             `json:"message,omitempty"`
	Source             *V1EventSource     `json:"source,omitempty"`
	FirstTimestamp     *V1Time            `json:"firstTimestamp,omitempty"`
	LastTimestamp      *V1Time            `json:"lastTimestamp,omitempty"`
	Count              int32              `json:"count,omitempty"`
	Type_              string             `json:"type,omitempty"`
	EventTime          *V1MicroTime       `json:"eventTime,omitempty"`
	Series             *V1EventSeries     `json:"series,omitempty"`
	Action             string             `json:"action,omitempty"`
	Related            *V1ObjectReference `json:"related,omitempty"`
	ReportingComponent string             `json:"reportingComponent,omitempty"`
	ReportingInstance  string             `json:"reportingInstance,omitempty"`
}
