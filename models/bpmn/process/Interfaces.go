package process

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/camunda"
	"github.com/deemount/gobpmnLab/models/bpmn/data"
	"github.com/deemount/gobpmnLab/models/bpmn/events"
	"github.com/deemount/gobpmnLab/models/bpmn/flow"
	"github.com/deemount/gobpmnLab/models/bpmn/gateways"
	"github.com/deemount/gobpmnLab/models/bpmn/pool"
	"github.com/deemount/gobpmnLab/models/bpmn/subprocesses"
	"github.com/deemount/gobpmnLab/models/bpmn/tasks"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// ProcessRepository ...
type ProcessRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseName
	attributes.AttributesBaseElements
	events.ProcessEventsElementsRepository
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository
	subprocesses.SubprocessesElementsRepository
	camunda.CamundaProcessAttributes
	flow.FlowSequenceFlow

	SetIsExecutable(isExec bool)
	GetIsExecutable() *bool

	SetLaneSet()
	GetLaneSet() *pool.LaneSet

	SetDataObject(num int)
	GetDataObject(num int) *data.DataObject
}
