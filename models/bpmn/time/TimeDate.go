package time

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewTimeDate ...
func NewTimeDate() TimeDateRepository {
	return &TimeDate{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timedate *TimeDate) SetTimerDefinitionType() {
	timedate.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timedate *TimeDate) SetTimerDefinition(timerDefinition string) {
	timedate.TimerDef = timerDefinition
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timedate TimeDate) GetTimerDefinitionType() gobpmnTypes.STR_PTR {
	return &timedate.TimerDefType
}

// GetTimerDefinition ...
func (timedate TimeDate) GetTimerDefinition() gobpmnTypes.STR_PTR {
	return &timedate.TimerDef
}
