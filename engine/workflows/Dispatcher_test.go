package workflows

import (
	"testing"

	"github.com/deemount/gobpmnLab/models/bpmn/events/elements"
)

// /opt/homebrew/opt/go/libexec/bin/go test -v -timeout 30s -run ^TestRegister$ github.com/deemount/gobpmnLab/models/bpmn/events
func TestRegister(t *testing.T) {

	t.Run("test -> register single", func(t *testing.T) {
		d := NewDispatcher()
		d.Register(elements.TStartEvent{})
		t.Logf("single: %+v", d.events)
	})

	t.Run("test -> register many", func(t *testing.T) {
		d := NewDispatcher()
		d.Register(
			elements.TStartEvent{},
			elements.TEndEvent{},
			elements.TBoundaryEvent{},
			elements.TIntermediateCatchEvent{},
			elements.TIntermediateThrowEvent{})
		t.Logf("many: %+v", d.events)
	})
}

// /opt/homebrew/opt/go/libexec/bin/go test -v -timeout 30s -run ^TestDispatch$ github.com/deemount/gobpmnLab/models/bpmn/events
/*
func TestDispatch(t *testing.T) {

	t.Run("test dispatch", func(t *testing.T) {
		d := NewDispatcher()
		d.Register(elements.TStartEvent{})
		err := d.Dispatch(context.Background(), elements.TStartEvent{
			BaseAttributes: gobpmnTypes.BaseAttributes{
				ID:   "Event_" + utils.GenerateHash(),
				Name: "Start",
			}})
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", d)
	})

}
*/
