package searchfx

import (
	"go.uber.org/fx"
	"simpleecommerceproductapi/product/search/handler"
	"simpleecommerceproductapi/product/search/repository"
	"simpleecommerceproductapi/product/search/usecases"
)

var Module = fx.Options(
	handler.Module,
	usecases.Module,
	repository.Module,

	fx.Provide(
		func(u usecases.Reader) handler.Service { return &u },
		func(m repository.RelationalReader) usecases.RepositoryReader { return &m },
	),
)
