package usecases

import "go.uber.org/fx"

var (
	factories = fx.Provide(NewCreate)
	Module    = fx.Options(factories)
)
