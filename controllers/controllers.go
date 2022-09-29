package controllers

import (
	auth "riskmanagement/controllers/auth"
	content "riskmanagement/controllers/content"
	user "riskmanagement/controllers/user"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(user.NewUserController),
	fx.Provide(auth.NewJWTAuthController),
	fx.Provide(content.NewContentController),
)
