package definitions

import (
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewLinkEventDefinition ...
func NewLinkEventDefinition() LinkEventDefinitionRepository {
	return &LinkEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (led *LinkEventDefinition) SetID(typ string, suffix interface{}) {
	led.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (led LinkEventDefinition) GetID() gobpmnTypes.STR_PTR {
	return &led.ID
}
