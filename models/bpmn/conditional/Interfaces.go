package conditional

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// ConditionalScriptFormat ...
type ConditionalScriptFormat interface {
	SetScriptFormat(format string)
	GetScriptFormat() gobpmnTypes.STR_PTR
}

// ConditionalScript ...
type ConditionalScript interface {
	SetScript(script string)
	GetScript() gobpmnTypes.STR_PTR
}

// ConditionalType ...
type ConditionalType interface {
	SetConditionType()
	GetConditionType() gobpmnTypes.STR_PTR
}

type ConditionalExpression interface {
	SetExpression(expression string)
	GetExpression() gobpmnTypes.STR_PTR
}

// CompletionConditionRepository ...
type CompletionConditionRepository interface{}

// ConditionRepository ...
type ConditionRepository interface {
	ConditionalScriptFormat
	ConditionalScript
	ConditionalType
}

// ConditionExpressionRepository ...
type ConditionExpressionRepository interface {
	ConditionalScriptFormat
	ConditionalScript
	ConditionalType
	ConditionalExpression
}
