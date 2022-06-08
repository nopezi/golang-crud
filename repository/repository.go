package repository

import (
	ap "infolelang/repository/access_places"
	address "infolelang/repository/address"
	approval "infolelang/repository/approvals"
	asset "infolelang/repository/asset"
	category "infolelang/repository/categories"
	contact "infolelang/repository/contacts"
	faq "infolelang/repository/faq"
	image "infolelang/repository/images"
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
	fx.Provide(asset.NewBuildingAssetReporitory),
	fx.Provide(asset.NewVehicleAssetReporitory),
	fx.Provide(asset.NewAssetFacilityReporitory),
	fx.Provide(asset.NewAssetAccessPlaceReporitory),
	fx.Provide(asset.NewAssetImageReporitory),
	fx.Provide(category.NewCategoryReporitory),
	fx.Provide(kpknl.NewKpknlReporitory),
	fx.Provide(subCategory.NewSubCategoryReporitory),
	fx.Provide(address.NewAddressReporitory),
	fx.Provide(approval.NewApprovalReporitory),
	fx.Provide(contact.NewContactReporitory),
	fx.Provide(image.NewImageReporitory),
)
