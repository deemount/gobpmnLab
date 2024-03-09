package collaborative_process_factory

import "github.com/deemount/gobpmnLab/factory"

var Builder factory.Builder

type Proxy interface{ Build() Process }
