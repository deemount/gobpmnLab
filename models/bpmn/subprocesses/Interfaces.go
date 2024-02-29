package subprocesses

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/camunda"
	"github.com/deemount/gobpmnLab/models/bpmn/events"
	"github.com/deemount/gobpmnLab/models/bpmn/events/elements"
	"github.com/deemount/gobpmnLab/models/bpmn/flow"
	"github.com/deemount/gobpmnLab/models/bpmn/gateways"
	"github.com/deemount/gobpmnLab/models/bpmn/loop"
	"github.com/deemount/gobpmnLab/models/bpmn/marker"
	"github.com/deemount/gobpmnLab/models/bpmn/tasks"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

type SubprocessesBase interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseName
	attributes.AttributesBaseElements
	marker.MarkerIncomingOutgoing
	camunda.CamundaDefaultAttributes
}

// SubprocessesElementsRepository ...
type SubprocessesElementsRepository interface {
	SetCallActivity(num int)
	GetCallActivity(num int) *CallActivity
	SetSubProcess(num int)
	GetSubProcess(num int) *SubProcess
	SetTransaction(num int)
	GetTransaction(num int) *Transaction
	SetAdHocSubProcess(num int)
	GetAdHocSubProcess(num int) *AdHocSubProcess
}

// AdHocSubProcessRepository ...
type AdHocSubProcessRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
	gateways.GatewaysElementsRepository
	tasks.TasksElementsRepository

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() *bool

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics

	SetStartEvent(num int)
	GetStartEvent(num int) *elements.StartEvent
	SetEndEvent(num int)
	GetEndEvent(num int) *elements.EndEvent

	SetSubProcess(num int)
	GetSubProcess(num int) *SubProcess
	SetAdHocSubProcess(num int)
	GetAdHocSubProcess(num int) *AdHocSubProcess
}

// CallActivityRepository ...
type CallActivityRepository interface {
	SubprocessesBase

	SetCalledElement(element string)
	GetCalledElement() *string

	SetCamundaCalledElementTenantID(tenantID string)
	GetCamundaCalledElementTenantID() *string

	SetCamundaVariableMappingClass(class string)
	GetCamundaVariableMappingClass() *string

	SetStandardLoopCharacteristics()
	GetStandardLoopCharacteristics() *loop.StandardLoopCharacteristics

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics
}

// SubProcessRepository ...
type SubProcessRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository
	events.ProcessEventsElementsRepository
	SubprocessesElementsRepository

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() *bool

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics
}

// TransactionRepository ...
type TransactionRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
}
