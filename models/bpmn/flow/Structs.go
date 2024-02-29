package flow

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/canvas"
	"github.com/deemount/gobpmnLab/models/bpmn/conditional"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// flow
type DelegateParameter struct {
	SF    *SequenceFlow
	MF    *MessageFlow
	ED    *canvas.Edge
	WP    []canvas.Waypoint
	BS    canvas.Bounds  // bounds ref by value for ... // TODO: describe for what (look at collaborative_process -> fromCheckIncomingClaim)
	BSPTR *canvas.Bounds // bounds pointer to get bounds of last element
	ST    string         // source type
	TT    string         // target type
	T     string         // typ
	N     string         // name
	H     []string       // hash
}

type SourceTargetRef struct {
	SourceRef string `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef string `xml:"targetRef,attr" json:"targetRef,omitempty"`
}

// Association ...
type Association struct {
	ID string `xml:"id,attr"`
	SourceTargetRef
	attributes.Attributes
}

// TAssociation ...
type TAssociation struct {
	ID string `xml:"id,attr"`
	SourceTargetRef
	attributes.TAttributes
}

// DataInputAssociation ...
type DataInputAssociation struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	attributes.Attributes
}

// TDataInputAssociation ...
type TDataInputAssociation struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	attributes.TAttributes
}

// MessageFlow ...
type MessageFlow struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	SourceTargetRef
}

// TMessageFlow ...
type TMessageFlow struct {
	gobpmnTypes.BaseAttributes
	attributes.TAttributes
	SourceTargetRef
}

// SequenceFlow ...
type SequenceFlow struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	SourceTargetRef
	ConditionExpression []conditional.ConditionExpression `xml:"bpmn:conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}

// TSequenceFlow ...
type TSequenceFlow struct {
	gobpmnTypes.BaseAttributes
	attributes.TAttributes
	SourceTargetRef
	ConditionExpression []conditional.ConditionExpression `xml:"conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}
