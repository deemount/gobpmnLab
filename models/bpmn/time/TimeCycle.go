package time

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewTimeCycle ...
func NewTimeCycle() TimeCycleRepository {
	return &TimeCycle{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timecycle *TimeCycle) SetTimerDefinitionType() {
	timecycle.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timecycle *TimeCycle) SetTimerDefinition(timerDefinition string) {
	timecycle.TimerDef = timerDefinition
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timecycle TimeCycle) GetTimerDefinitionType() gobpmnTypes.STR_PTR {
	return &timecycle.TimerDefType
}

// GetTimerDefinition ...
func (timecycle TimeCycle) GetTimerDefinition() gobpmnTypes.STR_PTR {
	return &timecycle.TimerDef
}
