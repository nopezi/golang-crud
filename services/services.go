package services

import (
	ap "infolelang/services/access_places"
	asset "infolelang/services/asset"
	category "infolelang/services/category"
	faq "infolelang/services/faq"
	kpknl "infolelang/services/kpknl"
	subCategory "infolelang/services/sub_category"
	user "infolelang/services/user"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(user.NewUserService),
	fx.Provide(NewJWTAuthService),
	fx.Provide(NewTransactionService),
	fx.Provide(ap.NewAccessPlaceService),
	fx.Provide(faq.NewFaqService),
	fx.Provide(kpknl.NewKpknlService),
	fx.Provide(category.NewCategoryService),
	fx.Provide(subCategory.NewSubCategoryService),
	fx.Provide(asset.NewAssetService),
)
