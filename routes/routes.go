package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoutes),
	// fx.Provide(NewTransactionRoutes),
	fx.Provide(NewAccessPlaceRoutes),
	fx.Provide(NewFaqRoutes),
	fx.Provide(NewAssetRoutes),
	fx.Provide(NewKpknlRoutes),
	fx.Provide(NewCategoryRoutes),
	fx.Provide(NewSubCategoryRoutes),
	fx.Provide(NewFileManagerRoutes),
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
	// transactionRoutes TransactionRoutes,
	accessPlaceRoutes AccessPlaceRoutes,
	faqRoutes FaqRoutes,
	assetRoutes AssetRoutes,
	kpknlRoutes KpknlRoutes,
	categoryRoutes CategoryRoutes,
	subCategoryRoutes SubCategoryRoutes,
	fileManager FileManagerRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
		// transactionRoutes,
		accessPlaceRoutes,
		faqRoutes,
		assetRoutes,
		kpknlRoutes,
		categoryRoutes,
		subCategoryRoutes,
		fileManager,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
