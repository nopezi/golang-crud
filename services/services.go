package services

import (
	auth "riskmanagement/services/auth"
	contents "riskmanagement/services/contents"
	user "riskmanagement/services/user"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(user.NewUserService),
	fx.Provide(auth.NewJWTAuthService),
	fx.Provide(contents.NewContentService),
)
