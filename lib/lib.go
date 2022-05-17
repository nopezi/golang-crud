package lib

import (
	env "infolelang/lib/env"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(env.NewEnv),
	fx.Provide(NewLogger),
	// 5
	fx.Provide(NewDatabase),
	fx.Provide(NewElastic),
	fx.Provide(NewDatabases),
)
