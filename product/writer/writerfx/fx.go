package writerfx

import (
	"go.uber.org/fx"
	"simpleecommerceproductapi/product/writer/handler"
	"simpleecommerceproductapi/product/writer/repository"
	"simpleecommerceproductapi/product/writer/usecases"
)

var Module = fx.Options(
	handler.Module,
	usecases.Module,
	repository.Module,

	fx.Provide(
		func(u usecases.Writer) handler.CreateService { return &u },
		func(m repository.WriterMock) usecases.RepositoryWriter { return &m },
	),
)
