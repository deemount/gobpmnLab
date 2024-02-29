package pool

import (
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// pool
type DelegateParameter struct {
	T string   // typ
	N string   // name
	H []string // hash
}

/*
 * Elementary
 */

// FlowNodeRef ...
type FlowNodeRef struct {
	gobpmnTypes.CoreInnerID
}

// Lane ...
type Lane struct {
	gobpmnTypes.CoreID
	FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// TLane ...
type TLane struct {
	gobpmnTypes.CoreID
	FlowNodeRef []FlowNodeRef `xml:"flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// LaneSet ...
type LaneSet struct {
	gobpmnTypes.CoreID
	Lane []Lane `xml:"bpmn:lane,omitempty" json:"lane,omitempty"`
}

// TLaneSet ...
type TLaneSet struct {
	gobpmnTypes.CoreID
	Lane []TLane `xml:"lane,omitempty" json:"lane,omitempty"`
}
