package main

import (
	"go.uber.org/fx"
	"simpleecommerceproductapi/internal"
	"simpleecommerceproductapi/search"
)

func main() {
	app := newApp()
	app.Run()
}

func newApp() *fx.App {
	modules := fx.Options(
		internal.Module,
		search.Module,
	)

	return fx.New(modules)
}
