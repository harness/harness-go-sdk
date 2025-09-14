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

	var notification_event_data []byte
	var err error

	// Handle null notification_event_data - this is valid for some notification types
	if len(a.NotificationEventData) == 0 {
		return nil, fmt.Errorf("notification_event_data is empty")
	}

	// Check if notification_event_data is null
	if len(a.NotificationEventData) == 4 && string(a.NotificationEventData) == "null" {
		// null notification_event_data is valid - preserve it as null
		return json.Marshal((*Alias)(a))
	}

	// Now, peek into the raw JSON to find the type discriminator.
	probe := NotificationEventParamsDto{}
	if err = json.Unmarshal(a.NotificationEventData, &probe); err != nil {
		return nil, fmt.Errorf("failed to probe notification_event_data type: %w", err)
	}

	// Check if Type_ is nil to prevent nil pointer dereference
	if probe.Type_ == nil {
		return nil, fmt.Errorf("notification_event_data type field is missing or null")
	}

	switch *probe.Type_ {
	case DELEGATE_ResourceTypeEnum:
		notification_event_data, err = json.Marshal(a.DelegateEventNotificationParamsDto)
	case PIPELINE_ResourceTypeEnum:
		notification_event_data, err = json.Marshal(a.PipelineEventNotificationParamsDto)
	case STO_EXEMPTION_ResourceTypeEnum:
		notification_event_data, err = json.Marshal(a.StoExemptionEventNotificationParamsDto)
	case CHAOS_EXPERIMENT_ResourceTypeEnum:
		notification_event_data, err = json.Marshal(a.ChaosExperimentEventNotificationParamsDto)
	case SERVICE_LEVEL_OBJECTIVE_ResourceTypeEnum:
		notification_event_data, err = json.Marshal(a.SloEventNotificationParamsDto)
	default:
		return nil, fmt.Errorf("unknown resource type %s", *probe.Type_)
	}

	if err != nil {
		return nil, err
	}
	a.NotificationEventData = json.RawMessage(notification_event_data)

	return json.Marshal((*Alias)(a))
}
