package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(NewAccessPlaceRoutes),
	fx.Provide(NewFaqRoutes),
	fx.Provide(NewAssetRoutes),
	fx.Provide(NewKpknlRoutes),
	fx.Provide(NewCategoryRoutes),
	fx.Provide(NewSubCategoryRoutes),
	fx.Provide(NewFileManagerRoutes),
	fx.Provide(NewPostalcodeRoutes),
	fx.Provide(NewFacilityRoutes),
	fx.Provide(NewCertificateTypeRoutes),
	fx.Provide(NewMcsRoutes),
	fx.Provide(NewVehicleBrandRoutes),
	fx.Provide(NewVehicleCapacityRoutes),
	fx.Provide(NewVehicleCategoryRoutes),
	fx.Provide(NewVehicleColorRoutes),
	fx.Provide(NewVehicleTransmissionRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes UserRoutes,
	authRoutes AuthRoutes,
	accessPlaceRoutes AccessPlaceRoutes,
	faqRoutes FaqRoutes,
	assetRoutes AssetRoutes,
	kpknlRoutes KpknlRoutes,
	categoryRoutes CategoryRoutes,
	subCategoryRoutes SubCategoryRoutes,
	fileManager FileManagerRoutes,
	postalcode PostalcodeRoutes,
	facility FacilityRoutes,
	certificateType CertificateTypeRoutes,
	mcs McsRoutes,
	vehicleBrand VehicleBrandRoutes,
	vehicleCapacity VehicleCapacityRoutes,
	vehicleCategory VehicleCategoryRoutes,
	vehicleColor VehicleColorRoutes,
	vehicleTransmission VehicleTransmissionRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
		accessPlaceRoutes,
		faqRoutes,
		assetRoutes,
		kpknlRoutes,
		categoryRoutes,
		subCategoryRoutes,
		fileManager,
		postalcode,
		facility,
		certificateType,
		mcs,
		vehicleBrand,
		vehicleCapacity,
		vehicleCategory,
		vehicleColor,
		vehicleTransmission,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
