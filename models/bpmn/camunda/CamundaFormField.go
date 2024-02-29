package camunda

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewCamundaFormField ...
func NewCamundaFormField() CamundaFormFieldRepository {
	return &CamundaFormField{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (formfield *CamundaFormField) SetID(typ string, suffix interface{}) {
	switch typ {
	case "formfield":
		formfield.ID = fmt.Sprintf("FormField_%v", suffix)
		break
	case "id":
		formfield.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetLabel ...
func (formfield *CamundaFormField) SetLabel(label string) {
	formfield.Label = label
}

// SetType ...
// can be: string, long, boolean, date, enum, custom type
func (formfield *CamundaFormField) SetType(typ string) {
	formfield.Typ = typ
}

// SetDefaultValue ...
func (formfield *CamundaFormField) SetDefaultValue(defaultValue string) {
	formfield.DefaultValue = defaultValue
}

/*** Make Elements ***/

/** Camunda **/

// SetCamundaProperties ...
func (formfield *CamundaFormField) SetCamundaProperties() {
	formfield.CamundaProperties = make([]CamundaProperties, 1)
}

// SetCamundaValidation ...
func (formfield *CamundaFormField) SetCamundaValidation() {
	formfield.CamundaValidation = make([]CamundaValidation, 1)
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (formfield CamundaFormField) GetID() gobpmnTypes.STR_PTR {
	return &formfield.ID
}

// GetLabel ...
func (formfield CamundaFormField) GetLabel() gobpmnTypes.STR_PTR {
	return &formfield.Label
}

// GetType ...
// can be: string, long, boolean, date, enum, custom type
func (formfield CamundaFormField) GetType() gobpmnTypes.STR_PTR {
	return &formfield.Typ
}

// GetDefaultValue ...
func (formfield CamundaFormField) GetDefaultValue() gobpmnTypes.STR_PTR {
	return &formfield.DefaultValue
}

/*** Make Elements ***/

/** Camunda **/

// GetCamundaProperties ...
func (formfield CamundaFormField) GetCamundaProperties() *CamundaProperties {
	return &formfield.CamundaProperties[0]
}

// GetCamundaValidation ...
func (formfield CamundaFormField) GetCamundaValidation() *CamundaValidation {
	return &formfield.CamundaValidation[0]
}
