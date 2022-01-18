package repository

import "go.uber.org/fx"

var (
	factories = fx.Provide(NewRelationalWriter, NewWriterMock)
	Module    = fx.Options(factories)
)
