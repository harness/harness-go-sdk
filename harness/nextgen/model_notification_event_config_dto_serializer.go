package nextgen

import (
	"encoding/json"
	"fmt"
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

	// Handle null notification_event_data - this is valid for some notification types
	if len(a.NotificationEventData) == 0 {
		return fmt.Errorf("notification_event_data is empty")
	}

	// Check if notification_event_data is null
	if len(a.NotificationEventData) == 4 && string(a.NotificationEventData) == "null" {
		// null notification_event_data is valid - no need to process further
		return nil
	}

	// Now, peek into the raw JSON to find the type discriminator.
	probe := NotificationEventParamsDto{}
	if err := json.Unmarshal(a.NotificationEventData, &probe); err != nil {
		return fmt.Errorf("failed to probe notification_event_data type: %w", err)
	}

	// Check if Type_ is nil to prevent nil pointer dereference
	if probe.Type_ == nil {
		return fmt.Errorf("notification_event_data type field is missing or null")
	}

	switch *probe.Type_ {
	case DELEGATE_ResourceTypeEnum:
		err = json.Unmarshal(aux.NotificationEventData, &a.DelegateEventNotificationParamsDto)
	case PIPELINE_ResourceTypeEnum:
		err = json.Unmarshal(aux.NotificationEventData, &a.PipelineEventNotificationParamsDto)
	case CHAOS_EXPERIMENT_ResourceTypeEnum:
		err = json.Unmarshal(aux.NotificationEventData, &a.ChaosExperimentEventNotificationParamsDto)
	case SERVICE_LEVEL_OBJECTIVE_ResourceTypeEnum:
		err = json.Unmarshal(aux.NotificationEventData, &a.SloEventNotificationParamsDto)
	case STO_EXEMPTION_ResourceTypeEnum:
		err = json.Unmarshal(aux.NotificationEventData, &a.StoExemptionEventNotificationParamsDto)
	default:
		return fmt.Errorf("unknown resource type %s", *probe.Type_)
	}

	return err
}

func (a *NotificationEventConfigDto) MarshalJSON() ([]byte, error) {
	type Alias NotificationEventConfigDto

	// If NotificationEventData is already populated, use it directly
	if len(a.NotificationEventData) > 0 {
		// Check if it's null
		if len(a.NotificationEventData) == 4 && string(a.NotificationEventData) == "null" {
			// Preserve null as-is
			return json.Marshal((*Alias)(a))
		}

		// For non-null data, try to marshal the specific DTO if it's populated
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
			// No specific DTO is populated, use the original NotificationEventData
			notification_event_data = a.NotificationEventData
		}

		if err != nil {
			return nil, err
		}

		// Create a copy to avoid modifying the original
		result := *a
		result.NotificationEventData = json.RawMessage(notification_event_data)
		return json.Marshal((*Alias)(&result))
	}

	// If NotificationEventData is empty, return error
	return nil, fmt.Errorf("notification_event_data is empty")
}
