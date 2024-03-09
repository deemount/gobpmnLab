package factory

import (
	"strings"

	"github.com/deemount/gobpmnLab/models/bpmn/events/elements"
	"github.com/deemount/gobpmnLab/models/bpmn/pool"
	"github.com/deemount/gobpmnLab/utils"
)

type Count struct {
	NumProc       int
	NumPart       int
	NumMsg        int
	NumStartEvent int
	NumEndEvent   int
	NumTask       int
	NumFlow       int
	NumShape      int
	NumEdge       int
}

// inPool ...
func (c *Count) countPool(field, builderField string) {
	if c.isPool(field) {
		if strings.Contains(builderField, fieldProcess) {
			c.NumProc++
		}
		if strings.Contains(builderField, fieldID) {
			c.NumPart++
			c.NumShape++
		}
	}
}

// inMessage ...
func (c *Count) countMessage(field, builderField string) {
	if c.isMessage(field) {
		if strings.Contains(builderField, fieldMessage) {
			c.NumMsg++
			c.NumEdge++
		}
	}
}

// isStartEvent ...
func (c *Count) countStartEvent(builderField string) {
	if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == "" {
		c.NumStartEvent++
		c.NumShape++
	}
}

// isTask ...
func (c *Count) countTask(builderField string) {
	if strings.Contains(builderField, fieldTask) && utils.After(builderField, "From") == "" {
		c.NumTask++
		c.NumShape++
	}
}

// isEndEvent ...
func (c *Count) countEndEvent(builderField string) {
	if strings.Contains(builderField, fieldEndEvent) {
		c.NumEndEvent++
		c.NumShape++
	}
}

// isPool ...
func (c *Count) isPool(field string) bool { return pool.IsPool(field) }

// isMessage ...
func (c *Count) isMessage(field string) bool { return elements.IsMessage(field) }
