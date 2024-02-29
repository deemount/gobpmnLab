package definitions

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewMessageEventDefinition ...
func NewMessageEventDefinition() MessageEventDefinitionRepository {
	return &MessageEventDefinition{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (med *MessageEventDefinition) SetID(typ string, suffix interface{}) {
	med.ID = SetID(typ, suffix)
}

// SetMsgRef ...
func (med *MessageEventDefinition) SetMsgRef(suffix string) {
	med.MsgRef = fmt.Sprintf("Message_%s", suffix)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (med MessageEventDefinition) GetID() gobpmnTypes.STR_PTR {
	return &med.ID
}

// GetMsgRef ...
func (med MessageEventDefinition) GetMsgRef() gobpmnTypes.STR_PTR {
	return &med.MsgRef
}
