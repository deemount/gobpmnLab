package core

import (
	"github.com/deemount/gobpmnLab/models/bpmn/canvas"
	"github.com/deemount/gobpmnLab/models/bpmn/collaboration"
	"github.com/deemount/gobpmnLab/models/bpmn/events"
	"github.com/deemount/gobpmnLab/models/bpmn/impl"
	"github.com/deemount/gobpmnLab/models/bpmn/marker"
	"github.com/deemount/gobpmnLab/models/bpmn/process"
)

type DefinitionsElements interface {
	SetCollaboration()
	GetCollaboration() collaboration.COLLABORATION_PTR
	SetProcess(num int)
	GetProcess(num int) process.PROCESS_PTR
	SetCategory(num int)
	GetCategory(num int) marker.CATEGORY_PTR
	events.CoreEventsElementsRepository
	SetDiagram(num int)
	GetDiagram(num int) canvas.DIAGRAM_PTR
}

// DefinitionsRepository ...
type DefinitionsRepository interface {
	impl.IFBaseID
	DefinitionsElements
	SetBpmn()
	SetBpmnDI()
	SetOmgDI()
	SetDC()
	SetOmgDC()
	SetBioc()
	SetXSD()
	SetXSI()
	SetXsiSchemaLocation()
	SetXsiSchemaLocationHTTPS()

	SetTargetNamespace()
	SetCamundaSchema()
	SetZeebeSchema()
	SetModelerSchema()
	SetModelerExecutionPlatform()
	SetModelerExecutionPlatformVersion(version string)
	SetExporter()
	SetExporterVersion(version string)

	SetMainElements(num int)
	SetDefaultAttributes()
}