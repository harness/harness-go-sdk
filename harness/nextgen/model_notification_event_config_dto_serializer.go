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

	switch *probe.Type_ {
	case DELEGATE_ResourceTypeEnum:
		var delegateEventData DelegateEventNotificationParamsDto
		if err := json.Unmarshal(aux.NotificationEventData, &delegateEventData); err != nil {
			return err
		}
		a.NotificationEventDataDecoded = &delegateEventData

	case PIPELINE_ResourceTypeEnum:
		var pipelineEventData PipelineEventNotificationParamsDto
		if err := json.Unmarshal(aux.NotificationEventData, &pipelineEventData); err != nil {
			return err
		}
		a.NotificationEventDataDecoded = &pipelineEventData
	case CHAOS_EXPERIMENT_ResourceTypeEnum:
		var chaosExperimentEventData ChaosExperimentEventNotificationParamsDto
		if err := json.Unmarshal(aux.NotificationEventData, &chaosExperimentEventData); err != nil {
			return err
		}
		a.NotificationEventDataDecoded = &chaosExperimentEventData
	case SERVICE_LEVEL_OBJECTIVE_ResourceTypeEnum:
		var sloEventData SloEventNotificationParamsDto
		if err := json.Unmarshal(aux.NotificationEventData, &sloEventData); err != nil {
			return err
		}
		a.NotificationEventDataDecoded = &sloEventData
	case STO_EXEMPTION_ResourceTypeEnum:
		var stoExemptionEventData StoExemptionEventNotificationParamsDto
		if err := json.Unmarshal(aux.NotificationEventData, &stoExemptionEventData); err != nil {
			return err
		}
		a.NotificationEventDataDecoded = &stoExemptionEventData
	default:
		a.NotificationEventDataDecoded = &probe
	}

	return err
}

func (a *NotificationEventConfigDto) MarshalJSON() ([]byte, error) {
	type Alias NotificationEventConfigDto

	// If there is a decoded object, marshal it and put the raw JSON
	// into the NotificationEventData field before marshaling the parent struct.
	if a.NotificationEventDataDecoded != nil {
		var err error
		a.NotificationEventData, err = json.Marshal(a.NotificationEventDataDecoded)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal NotificationEventDataDecoded: %w", err)
		}
	}

	return json.Marshal((*Alias)(a))
}
