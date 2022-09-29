package services

import (
	auth "crud/services/auth"
	contents "crud/services/contents"
	user "crud/services/user"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(user.NewUserService),
	fx.Provide(auth.NewJWTAuthService),
	fx.Provide(contents.NewContentService),
)
