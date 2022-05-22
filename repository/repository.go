package repository

import (
	ap "infolelang/repository/access_places"
	asset "infolelang/repository/asset"
	category "infolelang/repository/categories"
	faq "infolelang/repository/faq"
	kpknl "infolelang/repository/kpknl"
	subCategory "infolelang/repository/sub_categories"
	user "infolelang/repository/user"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(user.NewUserRepository),
	fx.Provide(NewTransactionRepository),
	fx.Provide(ap.NewAccessPlaceReporitory),
	fx.Provide(faq.NewFaqReporitory),
	fx.Provide(asset.NewAssetReporitory),
	fx.Provide(category.NewCategoryReporitory),
	fx.Provide(kpknl.NewKpknlReporitory),
	fx.Provide(subCategory.NewSubCategoryReporitory),
)
