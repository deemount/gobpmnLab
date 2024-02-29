package camunda

import (
	"fmt"

	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewCamundaConnectorID ...
func NewCamundaConnectorID() CamundaConnectorIDRepository {
	return &CamundaConnectorID{}
}

/*
 * Default Setters
 */

/* Content */

// SetID ...
// Notice: has no hash
func (cconnectorId *CamundaConnectorID) SetID(typ string, suffix interface{}) {
	switch typ {
	case "id":
		cconnectorId.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/*
 * Default Getters
 */

/* Content */

// GetID ...
func (cconnectorId CamundaConnectorID) GetID() gobpmnTypes.STR_PTR {
	return &cconnectorId.ID
}
