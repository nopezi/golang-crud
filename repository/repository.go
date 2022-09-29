package repository

import (
	content "riskmanagement/repository/content"
	user "riskmanagement/repository/user"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(user.NewUserRepository),
	fx.Provide(content.NewContentRepository),
)
