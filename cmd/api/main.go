package main

import (
	"go.uber.org/fx"
	"simpleecommerceproductapi/internal"
	"simpleecommerceproductapi/product/search/searchfx"
)

func main() {
	app := newApp()
	app.Run()
}

func newApp() *fx.App {
	modules := fx.Options(
		internal.Module,
		searchfx.Module,
	)

	return fx.New(modules)
}
