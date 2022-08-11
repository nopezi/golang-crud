package controllers

import (
	ap "infolelang/controllers/access_places"
	asset "infolelang/controllers/asset"
	auth "infolelang/controllers/auth"
	banner "infolelang/controllers/banner"
	category "infolelang/controllers/category"
	certificate_type "infolelang/controllers/certificate_type"
	facility "infolelang/controllers/facility"
	faq "infolelang/controllers/faq"
	file "infolelang/controllers/file_manager"
	kpknl "infolelang/controllers/kpknl"
	mcs "infolelang/controllers/mcs"
	postalcode "infolelang/controllers/postalcode"
	subCategory "infolelang/controllers/sub_category"
	user "infolelang/controllers/user"
	vehicle_brand "infolelang/controllers/vehicle_brand"
	vehicle_capacity "infolelang/controllers/vehicle_capacity"
	vehicle_category "infolelang/controllers/vehicle_category"
	vehicle_color "infolelang/controllers/vehicle_color"
	vehicle_transmission "infolelang/controllers/vehicle_transmission"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(user.NewUserController),
	fx.Provide(auth.NewJWTAuthController),
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
	fx.Provide(mcs.NewMcsController),
	fx.Provide(vehicle_brand.NewVehicleBrandController),
	fx.Provide(vehicle_capacity.NewVehicleCapacityController),
	fx.Provide(vehicle_category.NewVehicleCategoryController),
	fx.Provide(vehicle_color.NewVehicleColorController),
	fx.Provide(vehicle_transmission.NewVehicleTransmissionController),
	fx.Provide(banner.NewBannerController),
)
