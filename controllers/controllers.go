package controllers

import (
	activity "riskmanagement/controllers/activity"
	auth "riskmanagement/controllers/auth"
	briefing "riskmanagement/controllers/briefing"
	coaching "riskmanagement/controllers/coaching"
	incident "riskmanagement/controllers/incident"
	materi "riskmanagement/controllers/materi"
	product "riskmanagement/controllers/product"
	riskIndicator "riskmanagement/controllers/riskindicator"
	riskIssue "riskmanagement/controllers/riskissue"
	riskType "riskmanagement/controllers/risktype"
	subactivity "riskmanagement/controllers/subactivity"
	subIncident "riskmanagement/controllers/subincident"
	unitKerja "riskmanagement/controllers/unitkerja"
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
	fx.Provide(riskType.NewRiskTypeController),
	fx.Provide(unitKerja.NewUnitKerjaController),
	fx.Provide(briefing.NewBriefingController),
	fx.Provide(materi.NewMateriController),
	fx.Provide(coaching.NewCoachingController),
)
