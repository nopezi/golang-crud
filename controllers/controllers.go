package controllers

import (
	activity "riskmanagement/controllers/activity"
	auth "riskmanagement/controllers/auth"
	product "riskmanagement/controllers/product"
	subactivity "riskmanagement/controllers/subactivity"
	user "riskmanagement/controllers/user"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(user.NewUserController),
	fx.Provide(auth.NewJWTAuthController),
	fx.Provide(activity.NewActivityController),
	fx.Provide(subactivity.NewSubActivityController),
	fx.Provide(product.NewProductController),
)
