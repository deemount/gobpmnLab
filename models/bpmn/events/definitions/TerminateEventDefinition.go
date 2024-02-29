package definitions

import (
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewTerminateEventDefinition ...
func NewTerminateEventDefinition() TerminateEventDefinitionRepository {
	return &TerminateEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ted *TerminateEventDefinition) SetID(typ string, suffix interface{}) {
	ted.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ted TerminateEventDefinition) GetID() gobpmnTypes.STR_PTR {
	return &ted.ID
}
