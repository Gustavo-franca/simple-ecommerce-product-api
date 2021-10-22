package server

import "go.uber.org/fx"

var (
	factories = fx.Provide(New, NewRouter)
	invoke    = fx.Invoke(StartServer)
	Module    = fx.Options(factories, invoke)
)
