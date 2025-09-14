package nextgen

import (
	"encoding/json"
)

func (a *NotificationEventConfigDto) UnmarshalJSON(data []byte) error {

	type Alias NotificationEventConfigDto

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	// Handle empty notification_event_data - this is valid for some notification types
	if len(a.NotificationEventData) == 0 {
		// Empty notification_event_data is valid - no need to process further
		return nil
	}

	// Check if notification_event_data is null
	if len(a.NotificationEventData) == 4 && string(a.NotificationEventData) == "null" {
		// null notification_event_data is valid - no need to process further
		return nil
	}

	// Now, peek into the raw JSON to find the type discriminator.
	probe := NotificationEventParamsDto{}
	if err := json.Unmarshal(a.NotificationEventData, &probe); err != nil {
		// If the JSON is malformed, skip type-specific unmarshaling
		// The NotificationEventData will be preserved as-is
		return nil
	}

	// Check if Type_ is nil to prevent nil pointer dereference
	if probe.Type_ == nil {
		// If type field is missing, skip type-specific unmarshaling
		return nil
	}

	switch *probe.Type_ {
	case DELEGATE_ResourceTypeEnum:
		a.DelegateEventNotificationParamsDto = &DelegateEventNotificationParamsDto{}
		err = json.Unmarshal(aux.NotificationEventData, a.DelegateEventNotificationParamsDto)
	case PIPELINE_ResourceTypeEnum:
		a.PipelineEventNotificationParamsDto = &PipelineEventNotificationParamsDto{}
		err = json.Unmarshal(aux.NotificationEventData, a.PipelineEventNotificationParamsDto)
	case CHAOS_EXPERIMENT_ResourceTypeEnum:
		a.ChaosExperimentEventNotificationParamsDto = &ChaosExperimentEventNotificationParamsDto{}
		err = json.Unmarshal(aux.NotificationEventData, a.ChaosExperimentEventNotificationParamsDto)
	case SERVICE_LEVEL_OBJECTIVE_ResourceTypeEnum:
		a.SloEventNotificationParamsDto = &SloEventNotificationParamsDto{}
		err = json.Unmarshal(aux.NotificationEventData, a.SloEventNotificationParamsDto)
	case STO_EXEMPTION_ResourceTypeEnum:
		a.StoExemptionEventNotificationParamsDto = &StoExemptionEventNotificationParamsDto{}
		err = json.Unmarshal(aux.NotificationEventData, a.StoExemptionEventNotificationParamsDto)
	default:
		// Unknown resource type, skip type-specific unmarshaling
		return nil
	}

	return err
}

func (a *NotificationEventConfigDto) MarshalJSON() ([]byte, error) {
	type Alias NotificationEventConfigDto

	// Always use the original NotificationEventData if it's available
	if len(a.NotificationEventData) > 0 {
		// Validate that the NotificationEventData contains valid JSON
		var validate interface{}
		if err := json.Unmarshal(a.NotificationEventData, &validate); err != nil {
			// If the JSON is malformed, set it to null to avoid API errors
			result := *a
			result.NotificationEventData = json.RawMessage(`null`)
			return json.Marshal((*Alias)(&result))
		}

		// Create a copy to avoid modifying the original
		result := *a
		return json.Marshal((*Alias)(&result))
	}

	// If NotificationEventData is empty, try to reconstruct from specific DTOs
	var notification_event_data []byte
	var err error

	// Check which specific DTO is populated and marshal that
	if a.PipelineEventNotificationParamsDto != nil {
		notification_event_data, err = json.Marshal(a.PipelineEventNotificationParamsDto)
	} else if a.DelegateEventNotificationParamsDto != nil {
		notification_event_data, err = json.Marshal(a.DelegateEventNotificationParamsDto)
	} else if a.ChaosExperimentEventNotificationParamsDto != nil {
		notification_event_data, err = json.Marshal(a.ChaosExperimentEventNotificationParamsDto)
	} else if a.SloEventNotificationParamsDto != nil {
		notification_event_data, err = json.Marshal(a.SloEventNotificationParamsDto)
	} else if a.StoExemptionEventNotificationParamsDto != nil {
		notification_event_data, err = json.Marshal(a.StoExemptionEventNotificationParamsDto)
	} else {
		// If no specific DTO is populated and NotificationEventData is empty,
		// set it to null to avoid API errors
		notification_event_data = json.RawMessage(`null`)
	}

	if err != nil {
		return nil, err
	}

	// Create a copy to avoid modifying the original
	result := *a
	result.NotificationEventData = json.RawMessage(notification_event_data)
	return json.Marshal((*Alias)(&result))
}
