package definitions

import (
	"fmt"

	"github.com/deemount/gobpmnLab/models/bpmn/impl"
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
func (sed SignalEventDefinition) GetID() impl.STR_PTR {
	return &sed.ID
}

// GetSignalRef ...
func (sed SignalEventDefinition) GetSignalRef() *string {
	return &sed.SignalRef
}
