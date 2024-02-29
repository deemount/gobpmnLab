package definitions

import (
	"github.com/deemount/gobpmnLab/models/bpmn/conditional"
	"github.com/deemount/gobpmnLab/models/bpmn/time"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// @EndEvent only
type DefinitionsGetTerminateBase interface {
	GetTerminateEventDefinition() *TerminateEventDefinition
}

type DefinitionsGetElementsBase interface {
	GetMessageEventDefinition() *MessageEventDefinition
	GetEscalationEventDefinition() *EscalationEventDefinition
	GetErrorEventDefinition() *ErrorEventDefinition
	GetSignalEventDefinition() *SignalEventDefinition
	GetCompensateEventDefinition() *CompensateEventDefinition
}

type DefinitionsGetElements interface {
	DefinitionsGetElementsBase
	// @BoundaryEvent
	GetCancelEventDefinition() *CancelEventDefinition
	GetTimerEventDefinition() *TimerEventDefinition
	GetConditionalEventDefinition() *ConditionalEventDefinition
}

type DefinitionsBase interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseName
}

// CancelEventDefinitionRepository ...
type CancelEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID
}

// CompensateEventDefinitionRepository ...
type CompensateEventDefinitionRepository interface{}

// ConditionalEventDefinitionRepository ...
type ConditionalEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID

	SetCamundaVariableName(variableName string)
	GetCamundaVariableName() *string

	SetCondition()
	GetCondition() *conditional.Condition
}

// ErrorEventDefinitionRepository ...
type ErrorEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID
}

// EscalationEventDefinitionRepository ...
type EscalationEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID
}

// LinkEventDefinitionRepository ...
type LinkEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID
}

// MessageEventDefinitionRepository ...
type MessageEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID

	SetMsgRef(suffix string)
	GetMsgRef() gobpmnTypes.STR_PTR
}

// SignalEventDefinitionRepository ...
type SignalEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID

	SetSignalRef(suffix string)
	GetSignalRef() *string
}

// TerminateEventDefinitionRepository ...
type TerminateEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID
}

// TimerEventDefinitionRepository ...
type TimerEventDefinitionRepository interface {
	gobpmnTypes.IFBaseID

	SetTimeDate()
	GetTimeDate() *time.TimeDate
	SetTimeDuration()
	GetTimeDuration() *time.TimeDuration
}
