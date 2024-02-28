package attributes

import (
	"github.com/deemount/gobpmnLab/models/bpmn/camunda"
	"github.com/deemount/gobpmnLab/models/bpmn/impl"
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
	impl.IFBaseID
	impl.IFBaseName
}
