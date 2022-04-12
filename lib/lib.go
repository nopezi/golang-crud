package lib

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewEnv),
	fx.Provide(NewLogger),
	// 5
	// fx.Provide(NewDatabase),
	fx.Provide(NewElastic),
)
