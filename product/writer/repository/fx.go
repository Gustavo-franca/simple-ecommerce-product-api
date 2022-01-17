package repository

import "go.uber.org/fx"

var (
	factories = fx.Provide(NewWriterMock)
	Module    = fx.Options(factories)
)
