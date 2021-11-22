package handler

import "go.uber.org/fx"

var (
	factories = fx.Provide(New)
	invoke    = fx.Invoke(RegisterWebHandlers)
	Module    = fx.Options(factories, invoke)
)
