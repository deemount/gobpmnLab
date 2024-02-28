package pool

import (
	"github.com/deemount/gobpmnLab/models/bpmn/impl"
)

// FlowNodeRefRepository ...
type FlowNodeRefRepository interface {
	impl.IFBaseID
}

// LaneRepository ...
type LaneRepository interface {
	impl.IFBaseID

	SetFlowNodeRef(num int)
	GetFlowNodeRef(num int) *FlowNodeRef
}

// LaneSetRepository ...
type LaneSetRepository interface {
	impl.IFBaseID
	SetLane(num int)
	GetLane(num int) *Lane
}
