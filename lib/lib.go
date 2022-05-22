package lib

import (
	env "infolelang/lib/env"

	minio "gitlab.com/golang-package-library/minio"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(env.NewEnv),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewElastic),
	fx.Provide(NewDatabases),
	fx.Provide(minio.NewMinio),
)
