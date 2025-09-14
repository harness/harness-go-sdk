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

	// Now, peek into the raw JSON to find the type discriminator.
	var probe NotificationEventParamsDto
	if len(a.NotificationEventData) > 0 && a.NotificationEventData[0] != 'n' { // ignore if null
		if err := json.Unmarshal(a.NotificationEventData, &probe); err != nil {
			return fmt.Errorf("failed to probe notification_event_data type: %w", err)
		}
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

	// Now, peek into the raw JSON to find the type discriminator.
	var probe NotificationEventParamsDto
	if len(a.NotificationEventData) > 0 && a.NotificationEventData[0] != 'n' { // ignore if null
		if err = json.Unmarshal(a.NotificationEventData, &probe); err != nil {
			return nil, fmt.Errorf("failed to probe notification_event_data type: %w", err)
		}
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
