package data

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/impl"
)

// DataObjectRepository ...
type DataObjectRepository interface {
	impl.IFBaseID
}

// DataObjectReferenceRepository ...
type DataObjectReferenceRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
	SetDataObjectRef(suffix interface{})
	GetDataObjectRef() *string
}

// DataStoreRepository ...
type DataStoreReferenceRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
}
