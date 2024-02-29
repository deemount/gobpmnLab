package time

import gobpmnTypes "github.com/deemount/gobpmnTypes"

type TimeBaseDefintionType interface {
	SetTimerDefinitionType()
	GetTimerDefinitionType() gobpmnTypes.STR_PTR
}

type TimeBaseDefinition interface {
	SetTimerDefinition(timerDefinition string)
	GetTimerDefinition() gobpmnTypes.STR_PTR
}

type TimeBaseCoreElements interface {
	TimeBaseDefintionType
	TimeBaseDefinition
}
type TimeBase interface {
	TimeBaseCoreElements
}

/*
 * @Repositories
 */

// TimeCycle ...
type TimeCycleRepository interface{ TimeBase }

// TimeDateRepository ...
type TimeDateRepository interface{ TimeBase }

// TimeDurationRepository ...
type TimeDurationRepository interface{ TimeBase }
