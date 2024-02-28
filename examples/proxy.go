package examples

import "github.com/deemount/gobpmnLab/factory"

var Builder factory.Builder

// Creator ...
// TODO: make it usable. bpmnfactory.set, build method must be overthinked
type Creator interface{ Create() interface{} }
