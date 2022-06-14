package services

import (
	ap "infolelang/services/access_places"
	asset "infolelang/services/asset"
	category "infolelang/services/category"
	certificate_type "infolelang/services/certificate_type"
	facility "infolelang/services/facility"
	faq "infolelang/services/faq"
	file "infolelang/services/file_manager"
	kpknl "infolelang/services/kpknl"
	mcs "infolelang/services/mcs"
	postalcode "infolelang/services/postalcode"
	subCategory "infolelang/services/sub_category"
	user "infolelang/services/user"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(user.NewUserService),
	fx.Provide(NewJWTAuthService),
	fx.Provide(ap.NewAccessPlaceService),
	fx.Provide(faq.NewFaqService),
	fx.Provide(kpknl.NewKpknlService),
	fx.Provide(category.NewCategoryService),
	fx.Provide(subCategory.NewSubCategoryService),
	fx.Provide(asset.NewAssetService),
	fx.Provide(file.NewFileManagerService),
	fx.Provide(postalcode.NewPostalcodeService),
	fx.Provide(facility.NewFacilityService),
	fx.Provide(certificate_type.NewCertificateTypeService),
	fx.Provide(mcs.NewMcsService),
)
