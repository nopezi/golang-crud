package repository

import (
	activity "riskmanagement/repository/activity"
	product "riskmanagement/repository/product"
	subactivity "riskmanagement/repository/subactivity"
	user "riskmanagement/repository/user"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(user.NewUserRepository),
	fx.Provide(activity.NewActivityRepository),
	fx.Provide(subactivity.NewSubActivityRepository),
	fx.Provide(product.NewProductRepository),
)
