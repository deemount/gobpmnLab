package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

/*
 * Base
 */

// CamundaBaseScriptElements ...
type CamundaBaseScriptElements interface {
	SetCamundaScript()
	GetCamundaScript() *CamundaScript
}

// CamundaDefaultAttributes ...
type CamundaDefaultAttributes interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
}

type CamundaProcessAttributes interface {
	SetCamundaVersionTag(tag string)
	GetCamundaVersionTag() *string
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
	SetCamundaTaskPriority(priority int)
	GetCamundaTaskPriority() *int
	SetCamundaCandidateStarterGroups(groups string)
	GetCamundaCandidateStarterGroups() *string
	SetCamundaCandiddateStarterUsers(users string)
	GetCamundaCandiddateStarterUsers() *string
	SetCamundaHistoryTimeToLive(tolive string)
	GetCamundaHistoryTimeToLive() *string
}

/*
 * Repositories
 */

// ExtensionElementsRepository ...
type ExtensionElementsRepository interface {
	SetCamundaProperties()
	GetCamundaProperties() *CamundaProperties
	SetCamundaFailedJobRetryCycle()
	GetCamundaFailedJobRetryCycle() *CamundaFailedJobRetryCycle
	SetCamundaFormData()
	GetCamundaFormData() *CamundaFormData
	SetCamundaInputOutput()
	GetCamundaInputOutput() *CamundaInputOutput
	SetCamundaTaskListener(num int)
	GetCamundaTaskListener(num int) *CamundaTaskListener
	SetCamundaExecutionListener(num int)
	GetCamundaExecutionListener(num int) *CamundaExecutionListener
	SetCamundaIn(num int)
	GetCamundaIn(num int) *CamundaIn
	SetCamundaOut(num int)
	GetCamundaOut(num int) *CamundaOut
}

// CamundaConnectorRepository ...
type CamundaConnectorRepository interface{}

// CamundaConnectorIDRepository ...
type CamundaConnectorIDRepository interface{}

// CamundaConstraintRepository ...
type CamundaConstraintRepository interface {
	gobpmnTypes.IFBaseName
	gobpmnTypes.IFBaseConfig
}

// CamundaEntryRepository ...
type CamundaEntryRepository interface {
	gobpmnTypes.IFBaseValue
	gobpmnTypes.IFBaseKey
}

// CamundaExecutionListener ...
type CamundaExecutionListenerRepository interface {
	gobpmnTypes.IFBaseEvent
	gobpmnTypes.IFBaseClass
	CamundaBaseScriptElements
	SetDelegateExpression(expr string)
	GetDelegateExpression() *string
	SetCamundaField(num int)
	GetCamundaField(num int) *CamundaField
}

// CamundaExpressionRepository ...
type CamundaExpressionRepository interface{}

// CamundaFailedJobyRetryCycleRepository  ...
type CamundaFailedJobRetryCycleRepository interface{}

// CamundaFieldRepository ...
type CamundaFieldRepository interface {
	gobpmnTypes.IFBaseName
	SetCamundaExpression()
	GetCamundaExpression() *CamundaExpression
	SetCamundaString()
	GetCamundaString() *CamundaString
}

// CamundaFormDataRepository ...
type CamundaFormDataRepository interface {
	SetCamundaFormField(num int)
	GetCamundaFormField(num int) *CamundaFormField
}

// CamundaFormField ...
type CamundaFormFieldRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseLabel
	gobpmnTypes.IFBaseType
	gobpmnTypes.IFBaseDefaultValue
	SetCamundaProperties()
	GetCamundaProperties() *CamundaProperties
	SetCamundaValidation()
	GetCamundaValidation() *CamundaValidation
}

// CamundaInRepository ...
type CamundaInRepository interface{}

// CamundaInputOutputRepository ...
type CamundaInputOutputRepository interface {
	SetCamundaInputParameter(num int)
	GetCamundaInputParameter(num int) *CamundaInputParameter
	SetCamundaOutputParameter(num int)
	GetCamundaOutputParameter(num int) *CamundaOutputParameter
}

// CamundaInputParameterRepository ...
type CamundaInputParameterRepository interface {
	CamundaBaseScriptElements
	gobpmnTypes.IFBaseLocalVariableName
	gobpmnTypes.IFBaseVariableAssignmentValue
	SetCamundaList()
	GetCamundaList() *CamundaList
	SetCamundaMap()
	GetCamundaMap() *CamundaMap
}

// CamundaListRepository ...
type CamundaListRepository interface {
	SetCamundaValue(num int)
	GetCamundaValue(num int) *CamundaValue
}

// CamundaMapRepository ...
type CamundaMapRepository interface {
	SetCamundaEntry(num int)
	GetCamundaEntry(num int) *CamundaEntry
}

// CamundaOutRepository ...
type CamundaOutRepository interface{}

// CamundaOutputParameter ...
type CamundaOutputParameterRepository interface {
	CamundaBaseScriptElements
	gobpmnTypes.IFBaseLocalVariableName
	gobpmnTypes.IFBaseVariableAssignmentValue
	SetCamundaList()
	GetCamundaList() *CamundaList
	SetCamundaMap()
	GetCamundaMap() *CamundaMap
}

// CamundaPropertiesRepository ...
type CamundaPropertiesRepository interface {
	SetCamundaProperty(num int)
	GetCamundaProperty(num int) *CamundaProperty
}

// CamundaPropertyRepository ...
type CamundaPropertyRepository interface {
	gobpmnTypes.IFBaseName
	gobpmnTypes.IFBaseValue
}

// CamundaScriptRepository ...
type CamundaScriptRepository interface {
	gobpmnTypes.IFBaseScriptFormat
}

// CamundaStringRepository ...
type CamundaStringRepository interface{}

// CamundaTaskListener ...
type CamundaTaskListenerRepository interface {
	gobpmnTypes.IFBaseEvent
	gobpmnTypes.IFBaseClass
	gobpmnTypes.IFBaseListenerID
	SetCamundaField(num int)
	GetCamundaField(num int) *CamundaField
}

// CamundaValidationRepository ...
type CamundaValidationRepository interface {
	SetCamundaConstraint(num int)
	GetCamundaConstraint(num int) *CamundaConstraint
}

// CamundaValueRepository ...
type CamundaValueRepository interface {
	gobpmnTypes.IFBaseValue
}
