package collaboration

import (
	"github.com/deemount/gobpmnLab/models/bpmn/attributes"
	"github.com/deemount/gobpmnLab/models/bpmn/flow"
	"github.com/deemount/gobpmnLab/models/bpmn/impl"
	"github.com/deemount/gobpmnLab/models/bpmn/loop"
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
	impl.CoreID
	attributes.Attributes
	Participant PARTICIPANT_SLC    `xml:"bpmn:participant" json:"participant,omitempty"`
	MessageFlow []flow.MessageFlow `xml:"bpmn:messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// TCollaboration ...
type TCollaboration struct {
	impl.CoreID
	attributes.TAttributes
	Participant TPARTICIPANT_SLC    `xml:"participant" json:"participant,omitempty"`
	MessageFlow []flow.TMessageFlow `xml:"messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// Participant ...
type Participant struct {
	impl.BaseAttributes
	ProcessRef              string                         `xml:"processRef,attr,omitempty" json:"processRef,omitempty" csv:"PROCESS_REF"`
	Documentation           []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"DOCUMENTATION"`
	ParticipantMultiplicity []loop.ParticipantMultiplicity `xml:"bpmn:participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty" csv:"PARTICIPANT_MULTIPLICITY"`
}

// TParticipant ...
type TParticipant struct {
	impl.BaseAttributes
	ProcessRef              string                          `xml:"processRef,attr" json:"processRef,omitempty"`
	Documentation           []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ParticipantMultiplicity []loop.TParticipantMultiplicity `xml:"participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty"`
}
