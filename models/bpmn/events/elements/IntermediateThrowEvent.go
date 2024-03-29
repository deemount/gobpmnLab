package elements

import (
	"context"
	"fmt"
	"log"

	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/events/definitions"
	"github.com/deemount/gobpmnLab/models/bpmn/marker"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewIntermediateThrowEvent ...
func NewIntermediateThrowEvent() IntermediateThrowEventRepository {
	return &IntermediateThrowEvent{}
}

func (intermediateThrowEvent TIntermediateThrowEvent) Handle(ctx context.Context) error {
	log.Printf("after: %+v\n", intermediateThrowEvent)
	return nil
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetID(typ string, suffix interface{}) {
	intermediateThrowEvent.ID = SetID(typ, suffix)
}

// SetName ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetName(name string) {
	intermediateThrowEvent.Name = name
}

/*** Make Elements +++*/

/** BPMN **/

// SetDocumentation ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetDocumentation() {
	intermediateThrowEvent.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetExtensionElements() {
	intermediateThrowEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetIncoming(num int) {
	intermediateThrowEvent.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetOutgoing(num int) {
	intermediateThrowEvent.Outgoing = make([]marker.Outgoing, num)
}

// SetCompensateEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetCompensateEventDefinition() {
	intermediateThrowEvent.CompensateEventDefinition = make([]definitions.CompensateEventDefinition, 1)
}

// SetLinkEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetLinkEventDefinition() {
	intermediateThrowEvent.LinkEventDefinition = make([]definitions.LinkEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetEscalationEventDefinition() {
	intermediateThrowEvent.EscalationEventDefinition = make([]definitions.EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetMessageEventDefinition() {
	intermediateThrowEvent.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetSignalEventDefinition() {
	intermediateThrowEvent.SignalEventDefinition = make([]definitions.SignalEventDefinition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (intermediateThrowEvent IntermediateThrowEvent) GetID() gobpmnTypes.STR_PTR {
	return &intermediateThrowEvent.ID
}

// GetName ...
func (intermediateThrowEvent IntermediateThrowEvent) GetName() gobpmnTypes.STR_PTR {
	return &intermediateThrowEvent.Name
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (intermediateThrowEvent IntermediateThrowEvent) GetDocumentation() *attributes.Documentation {
	return &intermediateThrowEvent.Documentation[0]
}

// GetExtensionElements ...
func (intermediateThrowEvent IntermediateThrowEvent) GetExtensionElements() *attributes.ExtensionElements {
	return &intermediateThrowEvent.ExtensionElements[0]
}

// GetIncoming ...
func (intermediateThrowEvent IntermediateThrowEvent) GetIncoming(num int) *marker.Incoming {
	return &intermediateThrowEvent.Incoming[num]
}

// GetOutgoing ...
func (intermediateThrowEvent IntermediateThrowEvent) GetOutgoing(num int) *marker.Outgoing {
	return &intermediateThrowEvent.Outgoing[num]
}

// GetCompensateEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetCompensateEventDefinition() *definitions.CompensateEventDefinition {
	return &intermediateThrowEvent.CompensateEventDefinition[0]
}

// GetLinkEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetLinkEventDefinition() *definitions.LinkEventDefinition {
	return &intermediateThrowEvent.LinkEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetEscalationEventDefinition() *definitions.EscalationEventDefinition {
	return &intermediateThrowEvent.EscalationEventDefinition[0]
}

// GetMessageEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &intermediateThrowEvent.MessageEventDefinition[0]
}

// GetSignalEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetSignalEventDefinition() *definitions.SignalEventDefinition {
	return &intermediateThrowEvent.SignalEventDefinition[0]
}

/*
 * Default String
 */

// String ...
func (intermediateThrowEvent IntermediateThrowEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", intermediateThrowEvent.ID, intermediateThrowEvent.Name)
}

// String ...
func (intermediateThrowEvent TIntermediateThrowEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", intermediateThrowEvent.ID, intermediateThrowEvent.Name)
}
