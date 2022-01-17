package product

import (
	"go.uber.org/fx"
)

var Module = fx.Invoke(Migrate)
