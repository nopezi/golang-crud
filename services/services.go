package services

import (
	activity "riskmanagement/services/activity"
	auth "riskmanagement/services/auth"
	incident "riskmanagement/services/incident"
	product "riskmanagement/services/product"
	riskIndicator "riskmanagement/services/riskindicator"
	riskIssue "riskmanagement/services/riskissue"
	subactivity "riskmanagement/services/subactivity"
	subincident "riskmanagement/services/subincident"
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
	fx.Provide(riskIndicator.NewRiskIndicatorService),
	fx.Provide(incident.NewIncidentService),
	fx.Provide(subincident.NewSubIncidentService),
)
