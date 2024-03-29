package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/marker"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

func NewUserTask() UserTaskRepository {
	return &UserTask{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (utask *UserTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		utask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		utask.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (utask *UserTask) SetName(name string) {
	utask.Name = name
}

/** Camunda **/

// SetCamundaFormKey ...
func (utask *UserTask) SetCamundaFormKey(formKey string) {
	utask.CamundaFormKey = formKey
}

// SetCamundaAsyncBefore ...
func (utask *UserTask) SetCamundaAsyncBefore(asyncBefore bool) {
	utask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (utask *UserTask) SetCamundaAsyncAfter(asyncAfter bool) {
	utask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (utask *UserTask) SetCamundaJobPriority(priority int) {
	utask.CamundaJobPriority = priority
}

// SetCamundaAssignee ...
func (utask *UserTask) SetCamundaAssignee(assignee string) {
	utask.CamundaAssignee = assignee
}

// SetCamundaCandidUsers ...
func (ut *UserTask) SetCamundaCandidateUsers(users string) {
	ut.CamundaCandidateUsers = users
}

// SetCamundaCandidGroups ...
func (ut *UserTask) SetCamundaCandidateGroups(groups string) {
	ut.CamundaCandidateGroups = groups
}

// SetCamundaDueDate ...
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask *UserTask) SetCamundaDueDate(due string) {
	utask.CamundaDueDate = due
}

// SetCamundaFollowUpDate
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask *UserTask) SetCamundaFollowUpDate(followUp string) {
	utask.CamundaFollowUpDate = followUp
}

// SetCamundaPriority ...
func (utask *UserTask) SetCamundaPriority(priority int) {
	utask.CamundaPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (utask *UserTask) SetDocumentation() {
	utask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (utask *UserTask) SetExtensionElements() {
	utask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (utask *UserTask) SetIncoming(num int) {
	utask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (utask *UserTask) SetOutgoing(num int) {
	utask.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (utask UserTask) GetID() gobpmnTypes.STR_PTR {
	return &utask.ID
}

// GetName ...
func (utask UserTask) GetName() gobpmnTypes.STR_PTR {
	return &utask.Name
}

/** Camunda **/

// GetCamundaFormKey ...
func (utask UserTask) GetCamundaFormKey() gobpmnTypes.STR_PTR {
	return &utask.CamundaFormKey
}

// GetCamundaAsyncBefore ...
func (utask UserTask) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &utask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (utask UserTask) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &utask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (utask UserTask) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &utask.CamundaJobPriority
}

// GetCamundaAssignee ...
func (utask UserTask) GetCamundaAssignee() gobpmnTypes.STR_PTR {
	return &utask.CamundaAssignee
}

// GetCamundaCandidUsers ...
func (ut UserTask) GetCamundaCandidateUsers() gobpmnTypes.STR_PTR {
	return &ut.CamundaCandidateUsers
}

// GetCamundaCandidGroups ...
func (ut UserTask) GetCamundaCandidateGroups() gobpmnTypes.STR_PTR {
	return &ut.CamundaCandidateGroups
}

// GetCamundaDueDate ...
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask UserTask) GetCamundaDueDate() gobpmnTypes.STR_PTR {
	return &utask.CamundaDueDate
}

// GetCamundaFollowUpDate
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask UserTask) GetCamundaFollowUpDate() gobpmnTypes.STR_PTR {
	return &utask.CamundaFollowUpDate
}

// GetCamundaPriority ...
func (utask UserTask) GetCamundaPriority() gobpmnTypes.INT_PTR {
	return &utask.CamundaPriority
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (utask UserTask) GetDocumentation() *attributes.Documentation {
	return &utask.Documentation[0]
}

// GetExtensionElements ...
func (utask UserTask) GetExtensionElements() *attributes.ExtensionElements {
	return &utask.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (utask UserTask) GetIncoming(num int) *marker.Incoming {
	return &utask.Incoming[num]
}

// GetOutgoing ...
func (utask UserTask) GetOutgoing(num int) *marker.Outgoing {
	return &utask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (utask UserTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", utask.ID, utask.Name)
}
