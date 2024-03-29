package definitions

import (
	"github.com/deemount/gobpmnLab/models/bpmn/conditional"
	"github.com/deemount/gobpmnLab/models/bpmn/time"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

type DelegateParameter struct {
	CED *CancelEventDefinition
	CPD *CompensateEventDefinition
	COD *ConditionalEventDefinition
	EED *EscalationEventDefinition
	ERD *ErrorEventDefinition
	MED *MessageEventDefinition
	SED *SignalEventDefinition
	TED *TimerEventDefinition
	T   string   // Typ
	TD  string   // TimerDefinition
	H   []string // Hash
}

// Boundaries ...
type Boundaries struct {
	CancelEventDefinition      CANCEL_EVENT_DEF_SLC      `xml:"bpmn:cancelEventDefinition,omitempty" json:"cancelEventDefinition,omitempty"`
	MessageEventDefinition     MESSAGE_EVENT_DEF_SLC     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       TIMER_EVENT_DEF_SLC       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  ESCALATION_EVENT_DEF_SLC  `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition CONDITIONAL_EVENT_DEF_SLC `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       ERROR_EVENT_DEF_SLC       `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      SIGNAL_EVENT_DEF_SLC      `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  COMPENSATE_EVENT_DEF_SLC  `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
}

type TBoundaries struct {
	CancelEventDefinition      CANCEL_EVENT_DEF_SLC          `xml:"cancelEventDefinition,omitempty" json:"cancelEventDefinition,omitempty"`
	MessageEventDefinition     MESSAGE_EVENT_DEF_SLC         `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       []TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  ESCALATION_EVENT_DEF_SLC      `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition []TConditionalEventDefinition `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       ERROR_EVENT_DEF_SLC           `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      SIGNAL_EVENT_DEF_SLC          `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  COMPENSATE_EVENT_DEF_SLC      `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
}

type EndEvent struct {
	CompensateEventDefinition COMPENSATE_EVENT_DEF_SLC `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition ESCALATION_EVENT_DEF_SLC `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    MESSAGE_EVENT_DEF_SLC    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      ERROR_EVENT_DEF_SLC      `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     SIGNAL_EVENT_DEF_SLC     `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  TERMINATE_EVENT_DEF_SLC  `xml:"bpmn:terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
}

type TEndEvent struct {
	CompensateEventDefinition []TCompensateEventDefinition `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition ESCALATION_EVENT_DEF_SLC     `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    MESSAGE_EVENT_DEF_SLC        `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      ERROR_EVENT_DEF_SLC          `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     SIGNAL_EVENT_DEF_SLC         `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  TERMINATE_EVENT_DEF_SLC      `xml:"terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
}

type IntermediateCatchEvent struct {
	MessageEventDefinition     MESSAGE_EVENT_DEF_SLC     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	LinkEventDefinition        LINK_EVENT_DEF_SLC        `xml:"bpmn:linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	ConditionalEventDefinition CONDITIONAL_EVENT_DEF_SLC `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	TimerEventDefinition       TIMER_EVENT_DEF_SLC       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
}

type TIntermediateCatchEvent struct {
	ConditionalEventDefinition []TConditionalEventDefinition `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	LinkEventDefinition        LINK_EVENT_DEF_SLC            `xml:"linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	TimerEventDefinition       []TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	MessageEventDefinition     MESSAGE_EVENT_DEF_SLC         `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

type IntermediateThrowEvent struct {
	MessageEventDefinition    MESSAGE_EVENT_DEF_SLC    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	LinkEventDefinition       LINK_EVENT_DEF_SLC       `xml:"bpmn:linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	CompensateEventDefinition COMPENSATE_EVENT_DEF_SLC `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition ESCALATION_EVENT_DEF_SLC `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	SignalEventDefinition     SIGNAL_EVENT_DEF_SLC     `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
}

type TIntermediateThrowEvent struct {
	MessageEventDefinition    MESSAGE_EVENT_DEF_SLC        `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	LinkEventDefinition       LINK_EVENT_DEF_SLC           `xml:"linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	CompensateEventDefinition []TCompensateEventDefinition `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition ESCALATION_EVENT_DEF_SLC     `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	SignalEventDefinition     SIGNAL_EVENT_DEF_SLC         `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
}

type StartEvent struct {
	ConditionalEventDef    CONDITIONAL_EVENT_DEF_SLC `xml:"bpmn:conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MessageEventDefinition MESSAGE_EVENT_DEF_SLC     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef          TIMER_EVENT_DEF_SLC       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
}

type TStartEvent struct {
	ConditionalEventDef    []TConditionalEventDefinition `xml:"conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MessageEventDefinition MESSAGE_EVENT_DEF_SLC         `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef          []TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
}

// CancelEventDefinition ...
type CancelEventDefinition struct {
	gobpmnTypes.CoreID
}

// CompensateEventDefinition ...
type CompensateEventDefinition struct{}

// TCompensateEventDefinition ...
type TCompensateEventDefinition struct{}

// ConditionalEventDefinition ...
type ConditionalEventDefinition struct {
	gobpmnTypes.CoreID
	CamundaVariableName string                  `xml:"camunda:variableName" json:"variableName,omitempty"`
	Condition           []conditional.Condition `xml:"bpmn:condition,omitempty" json:"condition,omitempty"`
}

// TConditionalEventDefinition ...
type TConditionalEventDefinition struct {
	gobpmnTypes.CoreID
	CamundaVariableName string                  `xml:"variableName" json:"variableName,omitempty"`
	Condition           []conditional.Condition `xml:"condition,omitempty" json:"condition,omitempty"`
}

// ErrorEventDefinition ...
type ErrorEventDefinition struct {
	gobpmnTypes.CoreID
}

// EscalationEventDefinition ...
type EscalationEventDefinition struct {
	gobpmnTypes.CoreID
}

// LinkEventDefinition ...
type LinkEventDefinition struct {
	gobpmnTypes.CoreID
}

// MessageEventDefinition ...
type MessageEventDefinition struct {
	gobpmnTypes.CoreID
	MsgRef string `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
}

// SignalEventDefinition ...
type SignalEventDefinition struct {
	gobpmnTypes.CoreID
	SignalRef string `xml:"signalRef,attr,omitempty" json:"signalRef,omitempty"`
}

// TerminateEventDefinition ...
type TerminateEventDefinition struct {
	gobpmnTypes.CoreID
}

// TimerEventDefinition ...
type TimerEventDefinition struct {
	gobpmnTypes.CoreID
	TimeDate     []time.TimeDate     `xml:"bpmn:timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration []time.TimeDuration `xml:"bpmn:timeDuration,omitempty" json:"timeDuration,omitempty"`
}

// TTimerEventDefinition ...
type TTimerEventDefinition struct {
	gobpmnTypes.CoreID
	TimeDate     []time.TimeDate     `xml:"timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration []time.TimeDuration `xml:"timeDuration,omitempty" json:"timeDuration,omitempty"`
}
