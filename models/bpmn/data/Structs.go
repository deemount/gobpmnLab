package data

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// DataObject ...
type DataObject struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// DataObjectReference ...
type DataObjectReference struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	DataObjectRef string `xml:"dataObjectRef,attr,omitempty" json:"dataObjectRef"`
}

// TDataObjectReference ...
type TDataObjectReference struct {
	gobpmnTypes.BaseAttributes
	attributes.TAttributes
	DataObjectRef string `xml:"dataObjectRef,attr,omitempty" json:"dataObjectRef"`
}

// DataStoreReference ...
type DataStoreReference struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
}

// TDataStoreReference ...
type TDataStoreReference struct {
	gobpmnTypes.BaseAttributes
	attributes.TAttributes
}
