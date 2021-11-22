package usecases

import "go.uber.org/fx"

var (
	factories = fx.Provide(NewReader)
	Module    = fx.Options(factories)
)
