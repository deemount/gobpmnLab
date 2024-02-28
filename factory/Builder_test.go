package factory

import (
	"testing"

	"github.com/deemount/gobpmnLab/models/bpmn/core"
)

func TestBuilder_countEndEvent(t *testing.T) {
	var Def core.DefinitionsRepository
	type fields struct {
		Reflect  Reflect
		Settings Settings
		Suffix   string
		Map      map[string][]interface{}
	}
	type args struct {
		builderField string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Event",
			fields: fields{
				Reflect: Reflect{
					IF: Def,
				},
				Settings: Settings{},
				Suffix:   "ID",
				Map:      make(map[string][]interface{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Builder{
				Reflect:  tt.fields.Reflect,
				Settings: tt.fields.Settings,
				Suffix:   tt.fields.Suffix,

				Map: tt.fields.Map,
			}
			h.countEndEvent(tt.args.builderField)
		})
	}
}

func TestBuilder_countTask(t *testing.T) {
	var Def core.DefinitionsRepository
	type fields struct {
		Reflect   Reflect
		Settings  Settings
		Suffix    string
		HashTable []string
	}
	type args struct {
		builderField string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "count_task",
			fields: fields{
				Reflect: Reflect{
					IF: Def,
				},
				Settings:  Settings{},
				Suffix:    "",
				HashTable: []string{},
			},
			args: args{
				builderField: "Task",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Builder{
				Reflect:   tt.fields.Reflect,
				Settings:  tt.fields.Settings,
				Suffix:    tt.fields.Suffix,
				HashTable: tt.fields.HashTable,
			}
			h.countTask(tt.args.builderField)
		})
	}
}

func TestBuilder_countStartEvent(t *testing.T) {
	var Def core.DefinitionsRepository
	type fields struct {
		Reflect   Reflect
		Settings  Settings
		Suffix    string
		HashTable []string
	}
	type args struct {
		builderField string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "count_start_event",
			fields: fields{
				Reflect: Reflect{
					IF: Def,
				},
				Settings:  Settings{},
				Suffix:    "",
				HashTable: []string{},
			},
			args: args{
				builderField: "StartEvent",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Builder{
				Reflect:   tt.fields.Reflect,
				Settings:  tt.fields.Settings,
				Suffix:    tt.fields.Suffix,
				HashTable: tt.fields.HashTable,
			}
			h.countStartEvent(tt.args.builderField)
		})
	}
}
