package repository

import (
	repoAccessPlace "infolelang/repository/access_places"
	faq "infolelang/repository/faq"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewTransactionRepository),
	fx.Provide(repoAccessPlace.NewAccessPlaceReporitory),
	fx.Provide(faq.NewFaqReporitory),
)
