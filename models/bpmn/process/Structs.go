package process

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/data"
	"github.com/deemount/gobpmnLab/models/bpmn/events"
	"github.com/deemount/gobpmnLab/models/bpmn/flow"
	"github.com/deemount/gobpmnLab/models/bpmn/gateways"
	"github.com/deemount/gobpmnLab/models/bpmn/marker"
	"github.com/deemount/gobpmnLab/models/bpmn/pool"
	"github.com/deemount/gobpmnLab/models/bpmn/subprocesses"
	"github.com/deemount/gobpmnLab/models/bpmn/tasks"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// Process ...
type Process struct {
	gobpmnTypes.BaseAttributes
	attributes.Attributes
	events.ProcessEvents
	tasks.Tasks
	subprocesses.Subprocesses
	gateways.Gateways
	// Base
	IsExecutable bool `xml:"isExecutable,attr" json:"isExecutable,omitempty" csv:"IS_EXECUTABLE"`
	// Camunda
	CamundaVersionTag             string `xml:"camunda:versionTag,attr,omitempty" json:"versionTag,omitempty" csv:"VERSION_TAG"`
	CamundaJobPriority            int    `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty" csv:"JOB_PRIORITY"`
	CamundaTaskPriority           int    `xml:"camunda:taskPriority,attr,omitempty" json:"taskPriority,omitempty" csv:"TASK_PRIORITY"`
	CamundaCandidateStarterGroups string `xml:"camunda:candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty" csv:"CANDIDATE_STARTER_GROUPS"`
	CamundaCandidateStarterUsers  string `xml:"camunda:candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty" csv:"CANDIDATE_STARTER_USERS"`
	CamundaHistoryTimeToLive      string `xml:"camunda:historyTimeToLive,attr,omitempty" json:"historyTimeToLive,omitempty" csv:"HISTORY_TIME_TO_LIVE"`
	// Pool
	LaneSet []pool.LaneSet `xml:"bpmn:laneSet,omitempty" json:"laneSet,omitempty" csv:"-"`
	// Flow
	Association  []flow.Association  `xml:"bpmn:association,omitempty" json:"association,omitempty"`
	SequenceFlow []flow.SequenceFlow `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty" csv:"-"`
	// Marker
	Group []marker.Group `xml:"bpmn:group,omitempty" json:"group,omitempty" csv:"-"`
	// Data
	DataObject []data.DataObject `xml:"bpmn:dataObject,omitempty" json:"dataObject,omitempty"`
}

// TProcess ...
type TProcess struct {
	gobpmnTypes.BaseAttributes
	attributes.TAttributes
	events.TProcessEvents
	tasks.TTasks
	subprocesses.TSubprocesses
	gateways.TGateways
	// Base
	IsExecutable bool `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	// Camunda
	VersionTag             string `xml:"versionTag,attr,omitempty" json:"versionTag,omitempty"`
	JobPriority            int    `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	TaskPriority           int    `xml:"taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CandidateStarterGroups string `xml:"candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CandidateStarterUsers  string `xml:"candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	HistoryTimeToLive      string `xml:"historyTimeToLive,attr,omitempty" json:"historyTimeToLive,omitempty"`
	// Pool
	LaneSet []pool.LaneSet `xml:"laneSet,omitempty" json:"laneSet,omitempty"`
	// Flow
	Association  []flow.TAssociation  `xml:"association,omitempty" json:"association,omitempty"`
	SequenceFlow []flow.TSequenceFlow `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
	// Marker
	Group []marker.TGroup `xml:"group,omitempty" json:"group,omitempty" csv:"-"`
	// Data
	DataObject []data.DataObject `xml:"dataObject,omitempty" json:"dataObject,omitempty"`
}
