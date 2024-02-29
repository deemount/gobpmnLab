package marker

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

/*
 * Base
 */

// MarkerBaseReference ...
type MarkerBaseReferences interface {
	SetSourceRef(typ string, sourceRef interface{})
	GetSourceRef() *string
	SetTargetRef(typ string, targetRef interface{})
	GetTargetRef() *string
}

// MarkerFlow ...
type MarkerFlow interface {
	SetFlow(suffix interface{})
	GetFlow() *string
}

// MarkerIncoming ...
type MarkerIncoming interface {
	SetIncoming(num int)
	GetIncoming(num int) *Incoming
}

// MarkerOutgoing ...
type MarkerOutgoing interface {
	SetOutgoing(num int)
	GetOutgoing(num int) *Outgoing
}

// MarkerIncomingOutgoing
type MarkerIncomingOutgoing interface {
	MarkerIncoming
	MarkerOutgoing
}

// MarkerBase ...
type MarkerBase interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseName
}

/*
 * Repositories
 */

// CategoryRepository ...
type CategoryRepository interface {
	gobpmnTypes.IFBaseID
	attributes.AttributesBaseElements

	SetCategoryValue()
	GetCategoryValue() CATEGORY_VALUE_PTR
}

// CategoryValueRepository ...
type CategoryValueRepository interface {
	gobpmnTypes.IFBaseID

	SetValue(value string)
	GetValue() *string
}

// IncomingRepository ...
type IncomingRepository interface{ MarkerFlow }

// OutgoingRepository ...
type OutgoingRepository interface{ MarkerFlow }

// GroupRepository ...
type GroupRepository interface {
	gobpmnTypes.IFBaseID

	SetCategoryValueRef(suffix string)
	GetCategoryValueRef() *string

	attributes.AttributesBaseElements
}
