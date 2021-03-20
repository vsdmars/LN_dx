package pkg

// MPtype is the MP type that dx is going to generate cfg2 to.
type mp int

const (
	// PonfTelemetry used in Ponf data extracting to CFG2
	platTelemetry mp = iota
)

type actionType int

const (
	ponf actionType = iota
)

var mpToMp = map[string]mp{
	"plat-telemetry": platTelemetry,
}

var actionToAction = map[string]actionType{
	"ponf": ponf,
}

var mpActionType = map[mp][]actionType{
	platTelemetry: []actionType{ponf},
}
