package canvas

import (
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewBounds ...
func NewBounds() BoundsRepository {
	return &Bounds{}
}

/*
 * Default Setters
 */

// SetCoordinates ...
func (bnds *Bounds) SetCoordinates(x, y int) {
	bnds.X = x
	bnds.Y = y
}

// SetX ...
func (bnds *Bounds) SetX(x int) {
	bnds.X = x
}

// SetY ...
func (bnds *Bounds) SetY(y int) {
	bnds.Y = y
}

// SetSize ...
func (bnds *Bounds) SetSize(width, height int) {
	bnds.Width = width
	bnds.Height = height
}

// SetWidth ...
func (bnds *Bounds) SetWidth(width int) {
	bnds.Width = width
}

// SetHeight ...
func (bnds *Bounds) SetHeight(height int) {
	bnds.Height = height
}

/*
 * Default Getters
 */

// GetCoordinates ...
func (bnds Bounds) GetCoordinates() (gobpmnTypes.INT_PTR, gobpmnTypes.INT_PTR) {
	return &bnds.X, &bnds.Y
}

// GetX ...
func (bnds Bounds) GetX() gobpmnTypes.INT_PTR {
	return &bnds.X
}

// GetY ...
func (bnds Bounds) GetY() gobpmnTypes.INT_PTR {
	return &bnds.Y
}

// GetSize ...
func (bnds Bounds) GetSize() (gobpmnTypes.INT_PTR, gobpmnTypes.INT_PTR) {
	return &bnds.Width, &bnds.Height
}

// GetWidth ...
func (bnds Bounds) GetWidth() gobpmnTypes.INT_PTR {
	return &bnds.Width
}

// GetHeight ...
func (bnds Bounds) GetHeight() gobpmnTypes.INT_PTR {
	return &bnds.Height
}
