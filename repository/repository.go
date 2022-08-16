package repository

import (
	activity "riskmanagement/repository/activity"
	incident "riskmanagement/repository/incident"
	product "riskmanagement/repository/product"
	riskIndicator "riskmanagement/repository/riskindicator"
	riskIssue "riskmanagement/repository/riskissue"
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
	fx.Provide(riskIssue.NewRiskIssueRepository),
	fx.Provide(riskIndicator.NewRiskIndicatorRepository),
	fx.Provide(incident.NewIncidentRepository),
)
