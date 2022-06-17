package internal

import (
	"go.uber.org/fx"
	"simpleecommerceproductapi/internal/database"
	"simpleecommerceproductapi/internal/server"
	"simpleecommerceproductapi/internal/tracer"
)

var Module = fx.Options(
	tracer.Module,
	server.Module,
	database.Module,
)
