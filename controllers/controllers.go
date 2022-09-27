package controllers

import (
	activity "riskmanagement/controllers/activity"
	aplikasi "riskmanagement/controllers/aplikasi"
	auth "riskmanagement/controllers/auth"
	briefing "riskmanagement/controllers/briefing"
	coaching "riskmanagement/controllers/coaching"
	incident "riskmanagement/controllers/incident"
	materi "riskmanagement/controllers/materi"
	mcs "riskmanagement/controllers/mcs"
	product "riskmanagement/controllers/product"
	riskControl "riskmanagement/controllers/riskcontrol"
	riskIndicator "riskmanagement/controllers/riskindicator"
	riskIssue "riskmanagement/controllers/riskissue"
	riskType "riskmanagement/controllers/risktype"
	subactivity "riskmanagement/controllers/subactivity"
	subIncident "riskmanagement/controllers/subincident"
	unitKerja "riskmanagement/controllers/unitkerja"
	user "riskmanagement/controllers/user"
	verifikasi "riskmanagement/controllers/verifikasi"

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
	fx.Provide(verifikasi.NewVerifikasiController),
	fx.Provide(riskControl.NewRiskControlController),
	fx.Provide(aplikasi.NewAplikasiController),
	fx.Provide(mcs.NewMcsController),
)
