package repository

import "go.uber.org/fx"

var (
	factories = fx.Provide(NewReaderMock)
	Module    = fx.Options(factories)
)
