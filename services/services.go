package services

import (
	activity "riskmanagement/services/activity"
	auth "riskmanagement/services/auth"
	briefing "riskmanagement/services/briefing"
	coaching "riskmanagement/services/coaching"
	incident "riskmanagement/services/incident"
	materi "riskmanagement/services/materi"
	product "riskmanagement/services/product"
	riskControl "riskmanagement/services/riskcontrol"
	riskIndicator "riskmanagement/services/riskindicator"
	riskIssue "riskmanagement/services/riskissue"
	riskType "riskmanagement/services/risktype"
	subactivity "riskmanagement/services/subactivity"
	subincident "riskmanagement/services/subincident"
	unitKerja "riskmanagement/services/unitkerja"
	user "riskmanagement/services/user"
	verifikasi "riskmanagement/services/verifikasi"

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
	fx.Provide(riskType.NewRiskTypeService),
	fx.Provide(unitKerja.NewUnitKerjaService),
	fx.Provide(briefing.NewBriefingService),
	fx.Provide(materi.NewMateriService),
	fx.Provide(coaching.NewCoachingService),
	fx.Provide(verifikasi.NewVerifikasiService),
	fx.Provide(riskControl.NewRiskControService),
)
