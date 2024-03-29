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

// Subprocesses ...
type Subprocesses struct {
	CallActivity    CALL_ACTIVITY_SLC    `xml:"bpmn:callActivity,omitempty" json:"callActivity,omitempty" csv:"-"`
	SubProcess      SUBPROCESS_SLC       `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty" csv:"-"`
	Transaction     TRANSACTION_SLC      `xml:"bpmn:transaction,omitempty" json:"transaction,omitempty" csv:"-"`
	AdHocSubProcess ADHOC_SUBPROCESS_SLC `xml:"bpmn:adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty" csv:"-"`
}

// TSubprocesses ...
type TSubprocesses struct {
	CallActivity    TCALL_ACTIVITY_SLC    `xml:"callActivity,omitempty" json:"callActivity,omitempty"`
	SubProcess      TSUBPROCESS_SLC       `xml:"subProcess,omitempty" json:"subProcess,omitempty"`
	Transaction     TTRANSACTION_SLC      `xml:"transaction,omitempty" json:"transaction,omitempty"`
	AdHocSubProcess TADHOC_SUBPROCESS_SLC `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty" csv:"-"`
}

// AdHocSubProcess ...
type AdHocSubProcess struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	gateways.Gateways
	tasks.Tasks
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	StartEvent                       []elements.StartEvent                   `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty"`
	EndEvent                         []elements.EndEvent                     `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	SubProcess                       SUBPROCESS_SLC                          `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  ADHOC_SUBPROCESS_SLC                    `xml:"bpmn:adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	SequenceFlow                     []flow.SequenceFlow                     `xml:"bpmn:sequenceFlow,omitempty" json:"equenceFlow,omitempty"`
}

// TSubProcess ...
type TAdHocSubProcess struct {
	gobpmnTypes.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	gateways.TGateways
	tasks.TTasks
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	StartEvent                       []elements.TStartEvent                  `xml:"startEvent,omitemnpty" json:"startEvent,omitempty"`
	EndEvent                         []elements.TEndEvent                    `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	SubProcess                       TSUBPROCESS_SLC                         `xml:"subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  TADHOC_SUBPROCESS_SLC                   `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	SequenceFlow                     []flow.SequenceFlow                     `xml:"sequenceFlow,omitempty" json:"equenceFlow,omitempty"`
}

// CallActivity ...
type CallActivity struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	CalledElement                    string                                  `xml:"calledElement,attr,omitempty" json:"calledElement,omitempty"`
	CamundaCalledElementTenantID     string                                  `xml:"camunda:calledElementTenantId,attr,omitempty" json:"calledElementTenantId,omitempty"`
	CamundaVariableMappingClass      string                                  `xml:"camunda:variableMappingClass,attr,omitempty" json:"variableMappingClass,omitempty"`
	StandardLoopCharacteristics      []loop.StandardLoopCharacteristics      `xml:"bpmn:standardLoopCharacteristics,omitempty" json:"standardLoopCharacteristics,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics,omitempty"`
}

// TCallActivity ...
type TCallActivity struct {
	gobpmnTypes.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	CalledElement                    string                                  `xml:"calledElement,attr,omitempty" json:"calledElement,omitempty"`
	CamundaCalledElementTenantID     string                                  `xml:"calledElementTenantId,attr,omitempty" json:"calledElementTenantId,omitempty"`
	CamundaVariableMappingClass      string                                  `xml:"variableMappingClass,attr,omitempty" json:"variableMappingClass,omitempty"`
	StandardLoopCharacteristics      []loop.StandardLoopCharacteristics      `xml:"standardLoopCharacteristics,omitempty" json:"standardLoopCharacteristics,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics,omitempty"`
}

// SubProcess ...
type SubProcess struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	events.ProcessEvents
	tasks.Tasks
	gateways.Gateways
	Subprocesses
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	SequenceFlow                     []flow.SequenceFlow                     `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// TSubProcess ...
type TSubProcess struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	camunda.TCoreAttributes
	marker.IncomingOutgoing
	events.TProcessEvents
	tasks.TTasks
	gateways.TGateways
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	SubProcess                       TSUBPROCESS_SLC                         `xml:"subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  TADHOC_SUBPROCESS_SLC                   `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	SequenceFlow                     []flow.SequenceFlow                     `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// Transaction ...
type Transaction struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	SequenceFlow []flow.SequenceFlow `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// TTransaction ...
type TTransaction struct {
	gobpmnTypes.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	SequenceFlow []flow.SequenceFlow `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}
