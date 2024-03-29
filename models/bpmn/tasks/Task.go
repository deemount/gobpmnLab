package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/flow"
	"github.com/deemount/gobpmnLab/models/bpmn/marker"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewTask ...
func NewTask() TaskRepository {
	return &Task{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (task *Task) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		task.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		task.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (task *Task) SetName(name string) {
	task.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (task *Task) SetCamundaAsyncBefore(asyncBefore bool) {
	task.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (task *Task) SetCamundaAsyncAfter(asyncAfter bool) {
	task.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (task *Task) SetCamundaJobPriority(priority int) {
	task.CamundaJobPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

// SetProperty ...
func (task *Task) SetProperty() {
	task.Property = make([]attributes.Property, 1)
}

// SetDocumentation ...
func (task *Task) SetDocumentation() {
	task.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (task *Task) SetExtensionElements() {
	task.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetDataInputAssociation ...
func (task *Task) SetDataInputAssociation(num int) {
	task.DataInputAssociation = make([]flow.DataInputAssociation, num)
}

// SetIncoming ...
func (task *Task) SetIncoming(num int) {
	task.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (task *Task) SetOutgoing(num int) {
	task.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (task Task) GetID() gobpmnTypes.STR_PTR {
	return &task.ID
}

// GetName ...
func (task Task) GetName() gobpmnTypes.STR_PTR {
	return &task.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (task Task) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &task.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (task Task) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &task.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (task Task) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &task.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetProperty ...
func (task Task) GetProperty() *attributes.Property {
	return &task.Property[0]
}

// GetDocumentation ...
func (task Task) GetDocumentation() *attributes.Documentation {
	return &task.Documentation[0]
}

// GetExtensionElements ...
func (task Task) GetExtensionElements() *attributes.ExtensionElements {
	return &task.ExtensionElements[0]
}

// GetDataInputAssociation ...
func (task Task) GetDataInputAssociation(num int) *flow.DataInputAssociation {
	return &task.DataInputAssociation[num]
}

// GetIncoming ...
func (task Task) GetIncoming(num int) *marker.Incoming {
	return &task.Incoming[num]
}

// GetOutgoing ...
func (task Task) GetOutgoing(num int) *marker.Outgoing {
	return &task.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (task Task) String() string {
	return fmt.Sprintf("id=%v, name=%v", task.ID, task.Name)
}
