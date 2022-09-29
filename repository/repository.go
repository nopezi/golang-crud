package repository

import (
	content "crud/repository/content"
	user "crud/repository/user"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(user.NewUserRepository),
	fx.Provide(content.NewContentRepository),
)
