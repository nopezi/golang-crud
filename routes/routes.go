package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoutes),
	fx.Provide(NewActivityRoutes),
	fx.Provide(NewSubActivityRoutes),
	fx.Provide(NewProductRoutes),
	fx.Provide(NewRiskIssueRoutes),
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
	activityRoutes ActivityRoutes,
	subActivityRoutes SubActivityRoutes,
	productRoutes ProductRoutes,
	riskIssueRoutes RiskIssueRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
		activityRoutes,
		subActivityRoutes,
		productRoutes,
		riskIssueRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
