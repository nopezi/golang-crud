package controllers

import (
	activity "riskmanagement/controllers/activity"
	auth "riskmanagement/controllers/auth"
	incident "riskmanagement/controllers/incident"
	product "riskmanagement/controllers/product"
	riskIndicator "riskmanagement/controllers/riskindicator"
	riskIssue "riskmanagement/controllers/riskissue"
	subactivity "riskmanagement/controllers/subactivity"
	subIncident "riskmanagement/controllers/subincident"
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
	fx.Provide(riskIssue.NewRiskIssueController),
	fx.Provide(riskIndicator.NewRiskIndicatorController),
	fx.Provide(incident.NewIncidentController),
	fx.Provide(subIncident.NewSubIncidentController),
)
