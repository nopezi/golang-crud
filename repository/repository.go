package repository

import (
	ap "infolelang/repository/access_places"
	address "infolelang/repository/address"
	approval "infolelang/repository/approvals"
	asset "infolelang/repository/asset"
	category "infolelang/repository/categories"
	certificate_type "infolelang/repository/certificate_type"
	contact "infolelang/repository/contacts"
	facility "infolelang/repository/facilities"
	faq "infolelang/repository/faq"
	image "infolelang/repository/images"
	kpknl "infolelang/repository/kpknl"
	postalcode "infolelang/repository/postalcode"
	subCategory "infolelang/repository/sub_categories"
	user "infolelang/repository/user"
	vehicle_brand "infolelang/repository/vehicle_brand"
	vehicle_capacity "infolelang/repository/vehicle_capacity"
	vehicle_category "infolelang/repository/vehicle_category"
	vehicle_color "infolelang/repository/vehicle_color"
	vehicle_transmission "infolelang/repository/vehicle_transmission"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(user.NewUserRepository),
	// fx.Provide(NewTransactionRepository),
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
	fx.Provide(postalcode.NewPostalcodeReporitory),
	fx.Provide(facility.NewFacilitiesReporitory),
	fx.Provide(certificate_type.NewCertificateTypeReporitory),
	fx.Provide(vehicle_brand.NewVehicleBrandReporitory),
	fx.Provide(vehicle_capacity.NewVehicleCapacityReporitory),
	fx.Provide(vehicle_category.NewVehicleCategoryReporitory),
	fx.Provide(vehicle_color.NewVehicleColorReporitory),
	fx.Provide(vehicle_transmission.NewVehicleTransmissionReporitory),
)
