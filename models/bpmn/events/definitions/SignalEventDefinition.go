package definitions

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewSignalEventDefinition ...
func NewSignalEventDefinition() SignalEventDefinitionRepository {
	return &SignalEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (sed *SignalEventDefinition) SetID(typ string, suffix interface{}) {
	sed.ID = SetID(typ, suffix)
}

// SetSignalRef ...
func (sed *SignalEventDefinition) SetSignalRef(suffix string) {
	sed.SignalRef = fmt.Sprintf("Signal_%s", suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (sed SignalEventDefinition) GetID() gobpmnTypes.STR_PTR {
	return &sed.ID
}

// GetSignalRef ...
func (sed SignalEventDefinition) GetSignalRef() *string {
	return &sed.SignalRef
}
