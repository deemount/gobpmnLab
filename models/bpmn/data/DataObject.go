package data

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewDataObject ...
func NewDataObject() DataObjectRepository {
	return &DataObject{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (do *DataObject) SetID(typ string, suffix interface{}) {
	switch typ {
	case "dataobject":
		do.ID = fmt.Sprintf("DataObject_%v", suffix)
		break
	case "id":
		do.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (do DataObject) GetID() gobpmnTypes.STR_PTR {
	return &do.ID
}
