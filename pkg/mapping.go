package pkg

import (
	"dx/pkg/mp"
	"dx/pkg/mp/plat_telemetry"
)

// MPtype is the MP type that dx is going to generate cfg2 to.
type mpType int

const (
	// PonfTelemetry used in Ponf data extracting to CFG2
	platTelemetry mpType = iota
)

type actionType int

const (
	ponf actionType = iota
)

var mpToMp = map[string]mpType{
	"plat-telemetry": platTelemetry,
}

var actionToAction = map[string]actionType{
	"ponf": ponf,
}

var mpActionRunner = map[mpType]map[actionType]mp.Actioner{
	platTelemetry: {
		ponf: &plat_telemetry.Runner{},
	},
}
