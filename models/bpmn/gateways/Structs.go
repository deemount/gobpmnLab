package gateways

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/camunda"
	"github.com/deemount/gobpmnLab/models/bpmn/marker"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// DelegateParameter ...
type DelegateParameter struct {
	CG  *ComplexGateway
	EBG *EventBasedGateway
	EG  *ExclusiveGateway
	IG  *InclusiveGateway
	PG  *ParallelGateway
}

/*
 * Elementary
 */

// Gateways ...
type Gateways struct {
	ExclusiveGateway  EXCLUSIVE_GATEWAY_SLC    `xml:"bpmn:exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty" csv:"-"`
	InclusiveGateway  INCLUSIVE_GATEWAY_SLC    `xml:"bpmn:inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty" csv:"-"`
	ParallelGateway   PARALLEL_GATEWAY_SLC     `xml:"bpmn:parallelGateway,omitempty" json:"parallelGateway,omitempty" csv:"-"`
	ComplexGateway    COMPLEX_GATEWAY_SLC      `xml:"bpmn:complexGateway,omitempty" json:"complexGateway,omitempty" csv:"-"`
	EventBasedGateway EVENT_BASED_GATEWAYS_SLC `xml:"bpmn:eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty" csv:"-"`
}

// TGateways ...
type TGateways struct {
	ExclusiveGateway  TEXCLUSIVE_GATEWAY_SLC    `xml:"exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty" csv:"-"`
	InclusiveGateway  TINCLUSIVE_GATEWAY_SLC    `xml:"inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty" csv:"-"`
	ParallelGateway   TPARALLEL_GATEWAY_SLC     `xml:"parallelGateway,omitempty" json:"parallelGateway,omitempty" csv:"-"`
	ComplexGateway    TCOMPLEX_GATEWAY_SLC      `xml:"complexGateway,omitempty" json:"complexGateway,omitempty" csv:"-"`
	EventBasedGateway TEVENT_BASED_GATEWAYS_SLC `xml:"eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty" csv:"-"`
}

// ComplexGateway ...
type ComplexGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TComplexGateway ...
type TComplexGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.TCoreAttributes
	attributes.Attributes
	marker.TIncomingOutgoing
}

// EventBasedGateway ...
type EventBasedGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TEventBasedGateway ...
type TEventBasedGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}

// ExclusiveGateway ...
type ExclusiveGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TExclusiveGateway ...
type TExclusiveGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}

// InclusiveGateway ...
type InclusiveGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TInclusiveGateway ...
type TInclusiveGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}

// ParallelGateway ...
type ParallelGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TParallelGateway ...
type TParallelGateway struct {
	gobpmnTypes.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}
