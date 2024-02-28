package small_process

import "github.com/deemount/gobpmnLab/factory"

var Builder factory.Builder

type Proxy interface{ Build() Process }
