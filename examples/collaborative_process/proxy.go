package collaborative_process

import gobpmn_reflection "github.com/deemount/gobpmnReflection"

var Reflection gobpmn_reflection.Reflection

type Proxy interface{ Build() Process }
