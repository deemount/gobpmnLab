package canvas

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewWaypoint ...
func NewWaypoint() WaypointRepository {
	return &Waypoint{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetX ...
func (wp *Waypoint) SetX(x int) {
	wp.X = x
}

// SetY ...
func (wp *Waypoint) SetY(y int) {
	wp.Y = y
}

// SetCoordinates ...
func (wp *Waypoint) SetCoordinates(x, y int) {
	wp.X = x
	wp.Y = y
}

/*
 * Default Getters
 */

/* Attributes */

// GetX ...
func (wp Waypoint) GetX() gobpmnTypes.INT_PTR {
	return &wp.X
}

// GetY ...
func (wp Waypoint) GetY() gobpmnTypes.INT_PTR {
	return &wp.Y
}

// GetCoordinates ...
func (wp Waypoint) GetCoordinates() (gobpmnTypes.INT_PTR, gobpmnTypes.INT_PTR) {
	return &wp.X, &wp.Y
}
