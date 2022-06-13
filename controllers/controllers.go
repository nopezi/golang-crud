package controllers

import (
	ap "infolelang/controllers/access_places"
	asset "infolelang/controllers/asset"
	category "infolelang/controllers/category"
	certificate_type "infolelang/controllers/certificate_type"
	facility "infolelang/controllers/facility"
	faq "infolelang/controllers/faq"
	file "infolelang/controllers/file_manager"
	kpknl "infolelang/controllers/kpknl"
	postalcode "infolelang/controllers/postalcode"
	subCategory "infolelang/controllers/sub_category"
	user "infolelang/controllers/user"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(user.NewUserController),
	fx.Provide(NewJWTAuthController),
	fx.Provide(ap.NewAccessPlaceController),
	fx.Provide(faq.NewFaqController),
	fx.Provide(kpknl.NewKpknlController),
	fx.Provide(category.NewCategoryController),
	fx.Provide(subCategory.NewSubCategoryController),
	fx.Provide(asset.NewAssetController),
	fx.Provide(file.NewFileManagerController),
	fx.Provide(postalcode.NewPostalcodeController),
	fx.Provide(facility.NewFacilityController),
	fx.Provide(certificate_type.NewCertificateTypeController),
)
