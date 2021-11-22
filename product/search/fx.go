package search

import (
	"go.uber.org/fx"
	"simpleecommerceproductapi/search/handler"
)

var Module = fx.Options(
	handler.Module,
)
