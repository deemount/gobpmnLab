package data

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// DataObjectRepository ...
type DataObjectRepository interface {
	gobpmnTypes.IFBaseID
}

// DataObjectReferenceRepository ...
type DataObjectReferenceRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseName
	attributes.AttributesBaseElements
	SetDataObjectRef(suffix interface{})
	GetDataObjectRef() *string
}

// DataStoreRepository ...
type DataStoreReferenceRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseName
	attributes.AttributesBaseElements
}
