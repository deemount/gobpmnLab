package flow

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/conditional"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

/*
 * Base
 */

// FlowBaseReference ...
type FlowBaseReferences interface {
	SetSourceRef(typ string, sourceRef interface{})
	GetSourceRef() gobpmnTypes.STR_PTR
	SetTargetRef(typ string, targetRef interface{})
	GetTargetRef() gobpmnTypes.STR_PTR
}

// FlowSequenceFlow ...
type FlowSequenceFlow interface {
	SetSequenceFlow(num int)
	GetSequenceFlow(num int) *SequenceFlow
}

// FlowBase ...
type FlowBase interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseName
}

/*
 * Repositories
 */

// AssociationRepository ...
type AssociationRepository interface {
	gobpmnTypes.IFBaseID
	attributes.AttributesBaseElements

	FlowBaseReferences
}

// DataInputAssociationRepository ...
type DataInputAssociationRepository interface {
	gobpmnTypes.IFBaseID
	attributes.AttributesBaseElements
}

// MessageFlowRepository ...
type MessageFlowRepository interface {
	FlowBase
	attributes.AttributesBaseElements
	FlowBaseReferences
}

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	FlowBase
	FlowBaseReferences
	attributes.AttributesBaseElements

	SetConditionExpression()
	GetConditionExpression() *conditional.ConditionExpression
}
