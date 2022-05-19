package services

import (
	ap "infolelang/services/access_places"
	faq "infolelang/services/faq"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewJWTAuthService),
	fx.Provide(NewTransactionService),
	fx.Provide(ap.NewAccessPlaceService),
	fx.Provide(faq.NewFaqService),
)
