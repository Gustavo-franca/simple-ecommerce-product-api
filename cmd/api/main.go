package main

import (
	"go.uber.org/fx"
	"simpleecommerceproductapi/internal"
	"simpleecommerceproductapi/product/search/searchfx"
	"simpleecommerceproductapi/product/writer/writerfx"
)

func main() {
	app := newApp()
	app.Run()
}

func newApp() *fx.App {
	modules := fx.Options(
		internal.Module,
		searchfx.Module,
		writerfx.Module,
	)

	return fx.New(modules)
}
