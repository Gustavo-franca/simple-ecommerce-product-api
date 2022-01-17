package handler

import "go.uber.org/fx"

var (
	factories = fx.Provide(NewCreateHandler)
	invoke    = fx.Invoke(RegisterCreateHandlers)
	Module    = fx.Options(factories, invoke)
)
