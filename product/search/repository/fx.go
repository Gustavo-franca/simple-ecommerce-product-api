package repository

import "go.uber.org/fx"

var (
	factories = fx.Provide(NewReaderMock, NewRelationalReader)
	Module    = fx.Options(factories)
)
