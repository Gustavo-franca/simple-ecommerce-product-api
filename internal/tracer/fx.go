package tracer

import "go.uber.org/fx"

var (
	factories = fx.Provide(NewTracer)
	invoke    = fx.Invoke(StartTracer)
	Module    = fx.Options(factories, invoke)
)
