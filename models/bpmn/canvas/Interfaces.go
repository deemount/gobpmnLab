package canvas

import (
	gobpmnDiagram "github.com/deemount/gobpmnDiagram/pkg/canvas"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

type CanvasBoundsElements interface {
	SetBounds()
	GetBounds() *Bounds
}

type CanvasLabelElements interface {
	SetLabel()
	GetLabel() *Label
}

// BoundsRepository ...
type BoundsRepository interface {
	gobpmnDiagram.IFBaseCoordinates
	SetSize(width, height int)
	GetSize() (gobpmnTypes.INT_PTR, gobpmnTypes.INT_PTR)
	SetWidth(width int)
	GetWidth() gobpmnTypes.INT_PTR
	SetHeight(height int)
	GetHeight() gobpmnTypes.INT_PTR
}

// WaypointRepository ...
type WaypointRepository interface {
	gobpmnDiagram.IFBaseCoordinates
}

// Diagram ...
type DiagramRepository interface {
	gobpmnTypes.IFBaseID

	SetPlane()
	GetPlane() *Plane

	GetDescription() string
}

// Edge ...
type EdgeRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseElement
	CanvasLabelElements

	SetWaypoint()
	SetWaypoints(num int)
	GetWaypoint(num int) *Waypoint
}

// Label ...
type LabelRepository interface {
	CanvasBoundsElements
}

// PlaneRepository ...
type PlaneRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseElement

	SetAttrProcessElement(suffix string)
	SetAttrCollaborationElement(suffix string)

	SetShape(num int)
	GetShape(num int) *Shape

	SetEdge(num int)
	GetEdge(num int) *Edge

	GetDescription() string
}

// ShapeRepository ...
type ShapeRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseElement
	CanvasBoundsElements
	CanvasLabelElements

	SetIsHorizontal(isHorizontal bool)
	GetIsHorizontal() gobpmnTypes.BOOL_PTR

	SetIsMarkerVisible(isMarkerVisible bool)
	GetIsMarkerVisible() gobpmnTypes.BOOL_PTR
}
