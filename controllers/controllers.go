package controllers

import (
	auth "crud/controllers/auth"
	content "crud/controllers/content"
	user "crud/controllers/user"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(user.NewUserController),
	fx.Provide(auth.NewJWTAuthController),
	fx.Provide(content.NewContentController),
)
