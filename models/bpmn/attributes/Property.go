package attributes

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewProperty ...
func NewProperty() PropertyRepository {
	return &Property{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (prop *Property) SetID(typ string, suffix interface{}) {
	switch typ {
	case "property":
		prop.ID = fmt.Sprintf("Property_%v", suffix)
		break
	case "id":
		prop.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
// Notice: maybe put a placeholder like name="__targetRef_placeholder" in it
func (prop *Property) SetName(name string) {
	prop.Name = name
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (prop Property) GetID() gobpmnTypes.STR_PTR {
	return &prop.ID

}

// GetName ...
func (prop Property) GetName() gobpmnTypes.STR_PTR {
	return &prop.Name
}
