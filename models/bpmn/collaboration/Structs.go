package collaboration

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/flow"
	"github.com/deemount/gobpmnLab/models/bpmn/loop"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// collaboratiion
type DelegateParameter struct {
	PPT *Participant // element pointer
	T   string       // typ
	PR  string       // typProcessRef
	N   string       // name
	H   []string     // hash
}

/*
 * Elementary
 */

// Collaboration ...
type Collaboration struct {
	gobpmnTypes.CoreID
	attributes.Attributes
	Participant PARTICIPANT_SLC    `xml:"bpmn:participant" json:"participant,omitempty"`
	MessageFlow []flow.MessageFlow `xml:"bpmn:messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// TCollaboration ...
type TCollaboration struct {
	gobpmnTypes.CoreID
	attributes.TAttributes
	Participant TPARTICIPANT_SLC    `xml:"participant" json:"participant,omitempty"`
	MessageFlow []flow.TMessageFlow `xml:"messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// Participant ...
type Participant struct {
	gobpmnTypes.BaseAttributes
	ProcessRef              string                         `xml:"processRef,attr,omitempty" json:"processRef,omitempty" csv:"PROCESS_REF"`
	Documentation           []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"DOCUMENTATION"`
	ParticipantMultiplicity []loop.ParticipantMultiplicity `xml:"bpmn:participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty" csv:"PARTICIPANT_MULTIPLICITY"`
}

// TParticipant ...
type TParticipant struct {
	gobpmnTypes.BaseAttributes
	ProcessRef              string                          `xml:"processRef,attr" json:"processRef,omitempty"`
	Documentation           []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ParticipantMultiplicity []loop.TParticipantMultiplicity `xml:"participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty"`
}
