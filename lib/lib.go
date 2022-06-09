package lib

import (
	env "infolelang/lib/env"

	elastic "gitlab.com/golang-package-library/elasticsearch"
	minioEnv "gitlab.com/golang-package-library/env"
	logger "gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(env.NewEnv),
	fx.Provide(logger.NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(elastic.NewElastic),
	fx.Provide(NewDatabases),
	fx.Provide(minioEnv.NewEnv),
	fx.Provide(minio.NewMinio),
)
