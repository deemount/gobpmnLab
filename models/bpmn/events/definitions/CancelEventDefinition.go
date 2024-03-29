package definitions

import (
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewCancelEventDefinition ...
func NewCancelEventDefinition() CancelEventDefinitionRepository {
	return &CancelEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ced *CancelEventDefinition) SetID(typ string, suffix interface{}) {
	ced.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ced CancelEventDefinition) GetID() gobpmnTypes.STR_PTR {
	return &ced.ID
}
