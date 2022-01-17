package internal

import (
	"go.uber.org/fx"
	"simpleecommerceproductapi/internal/database"
	"simpleecommerceproductapi/internal/server"
)

var Module = fx.Options(
	server.Module,
	database.Module,
)
