package repository

import (
	activity "riskmanagement/repository/activity"
	briefing "riskmanagement/repository/briefing"
	coaching "riskmanagement/repository/coaching"
	files "riskmanagement/repository/files"
	incident "riskmanagement/repository/incident"
	materi "riskmanagement/repository/materi"
	product "riskmanagement/repository/product"
	riskIndicator "riskmanagement/repository/riskindicator"
	riskIssue "riskmanagement/repository/riskissue"
	riskType "riskmanagement/repository/risktype"
	subactivity "riskmanagement/repository/subactivity"
	subIncident "riskmanagement/repository/subincident"
	unitKerja "riskmanagement/repository/unitkerja"
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
	fx.Provide(subIncident.NewSubIncidentRepository),
	fx.Provide(riskType.NewRiskTypeRepository),
	fx.Provide(unitKerja.NewUnitKerjaRepository),
	fx.Provide(briefing.NewBriefingRepository),
	fx.Provide(briefing.NewBriefingMateriRepository),
	fx.Provide(files.NewFilesRepository),
	fx.Provide(materi.NewMateriRepository),
	fx.Provide(coaching.NewCoachingRepository),
	fx.Provide(coaching.NewCoachingActivityRepository),
)
