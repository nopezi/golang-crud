package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	// fx.Provide(NewUserRoutes),
	// fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(NewTransactionRoutes),
	fx.Provide(NewAccessPlaceRoutes),
	fx.Provide(NewFaqRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	// userRoutes UserRoutes,
	// authRoutes AuthRoutes,
	transactionRoutes TransactionRoutes,
	accessPlaceRoutes AccessPlaceRoutes,
	faqRoutes FaqRoutes,
) Routes {
	return Routes{
		// userRoutes,
		// authRoutes,
		transactionRoutes,
		accessPlaceRoutes,
		faqRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}