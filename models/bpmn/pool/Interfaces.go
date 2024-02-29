package pool

import (
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// FlowNodeRefRepository ...
type FlowNodeRefRepository interface {
	gobpmnTypes.IFBaseID
}

// LaneRepository ...
type LaneRepository interface {
	gobpmnTypes.IFBaseID

	SetFlowNodeRef(num int)
	GetFlowNodeRef(num int) *FlowNodeRef
}

// LaneSetRepository ...
type LaneSetRepository interface {
	gobpmnTypes.IFBaseID
	SetLane(num int)
	GetLane(num int) *Lane
}
