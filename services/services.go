package services

import (
	activity "riskmanagement/services/activity"
	auth "riskmanagement/services/auth"
	product "riskmanagement/services/product"
	riskIssue "riskmanagement/services/riskissue"
	subactivity "riskmanagement/services/subactivity"
	user "riskmanagement/services/user"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(user.NewUserService),
	fx.Provide(auth.NewJWTAuthService),
	fx.Provide(activity.NewActivityService),
	fx.Provide(subactivity.NewSubActivityService),
	fx.Provide(product.NewProductService),
	fx.Provide(riskIssue.NewRiskIssueService),
)
