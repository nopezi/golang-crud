package controllers

import (
	ap "infolelang/controllers/access_places"
	faq "infolelang/controllers/faq"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewJWTAuthController),
	fx.Provide(NewTransactionController),
	fx.Provide(ap.NewAccessPlaceController),
	fx.Provide(faq.NewFaqController),
)
