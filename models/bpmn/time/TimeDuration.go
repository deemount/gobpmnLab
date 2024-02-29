package time

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewTimeDuration ...
func NewTimeDuration() TimeDurationRepository {
	return &TimeDuration{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timeduration *TimeDuration) SetTimerDefinitionType() {
	timeduration.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timeduration *TimeDuration) SetTimerDefinition(timerDefinition string) {
	timeduration.TimerDef = timerDefinition
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timeduration TimeDuration) GetTimerDefinitionType() gobpmnTypes.STR_PTR {
	return &timeduration.TimerDefType
}

// GetTimerDefinition ...
func (timeduration TimeDuration) GetTimerDefinition() gobpmnTypes.STR_PTR {
	return &timeduration.TimerDef
}
