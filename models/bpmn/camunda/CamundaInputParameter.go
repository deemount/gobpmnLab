package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewCamundaInputParameter
func NewCamundaInputParameter() CamundaInputParameterRepository {
	return &CamundaInputParameter{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetLocalVariableName ...
func (inputParameter *CamundaInputParameter) SetLocalVariableName(variable string) {
	inputParameter.LocalVariableName = variable
}

// SetVariableAssignmentValue ...
func (inputParameter *CamundaInputParameter) SetVariableAssignmentValue(value string) {
	inputParameter.VariableAssignmentValue = value
}

/* Elements */

/** Camunda **/

// SetCamundaScript ...
func (inputParameter *CamundaInputParameter) SetCamundaScript() {
	inputParameter.CamundaScript = make([]CamundaScript, 1)
}

// SetCamundaList ...
func (inputParameter *CamundaInputParameter) SetCamundaList() {
	inputParameter.CamundaList = make([]CamundaList, 1)
}

// SetCamundaMap ...
func (inputParameter *CamundaInputParameter) SetCamundaMap() {
	inputParameter.CamundaMap = make([]CamundaMap, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetLocalVariableName ...
func (inputParameter CamundaInputParameter) GetLocalVariableName() gobpmnTypes.STR_PTR {
	return &inputParameter.LocalVariableName
}

// GetVariableAssignmentValue ...
func (inputParameter CamundaInputParameter) GetVariableAssignmentValue() gobpmnTypes.STR_PTR {
	return &inputParameter.VariableAssignmentValue
}

/* Elements */

/** Camunda **/

// GetCamundaScript ...
func (inputParameter CamundaInputParameter) GetCamundaScript() *CamundaScript {
	return &inputParameter.CamundaScript[0]
}

// GetCamundaList ...
func (inputParameter CamundaInputParameter) GetCamundaList() *CamundaList {
	return &inputParameter.CamundaList[0]
}

// GetCamundaMap ...
func (inputParameter CamundaInputParameter) GetCamundaMap() *CamundaMap {
	return &inputParameter.CamundaMap[0]
}
