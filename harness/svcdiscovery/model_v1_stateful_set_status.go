/*
 * Service Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package svcdiscovery

type V1StatefulSetStatus struct {
	// Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset. +optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`
	// collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision. +optional
	CollisionCount int32 `json:"collisionCount,omitempty"`
	// Represents the latest available observations of a statefulset's current state. +optional +patchMergeKey=type +patchStrategy=merge
	Conditions []V1StatefulSetCondition `json:"conditions,omitempty"`
	// currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
	CurrentReplicas int32 `json:"currentReplicas,omitempty"`
	// currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
	CurrentRevision string `json:"currentRevision,omitempty"`
	// observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server. +optional
	ObservedGeneration int32 `json:"observedGeneration,omitempty"`
	// readyReplicas is the number of pods created for this StatefulSet with a Ready Condition.
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
	// replicas is the number of Pods created by the StatefulSet controller.
	Replicas int32 `json:"replicas,omitempty"`
	// updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
	UpdateRevision string `json:"updateRevision,omitempty"`
	// updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`
}
