package attributes

import (
	"github.com/deemount/gobpmnLab/models/bpmn/camunda"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

type AttributesBaseElements interface {
	SetDocumentation()
	GetDocumentation() *Documentation
	SetExtensionElements()
	GetExtensionElements() *ExtensionElements
}

// DocumentationRepository ...
type DocumentationRepository interface {
	SetData(data string)
	GetData() *string
}

// ExtensionElementsRepository ...
type ExtensionElementsRepository interface {
	camunda.ExtensionElementsRepository
}

// PropertyRepository ...
type PropertyRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseName
}
