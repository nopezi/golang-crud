package middlewares

import "go.uber.org/fx"

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewJWTAuthMiddleware),
	// fx.Provide(NewDatabaseTrx),
	fx.Provide(NewMiddlewares),
	fx.Provide(NewLogActivityMiddleware),
)

// IMiddleware middleware interface
type IMiddleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []IMiddleware

// NewMiddlewares creates new middlewares
// Register the middleware that should be applied directly (globally)
func NewMiddlewares(
	corsMiddleware CorsMiddleware,
	// dbTrxMiddleware DatabaseTrx,
	logActivityMiddleware LogActivityMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
		// dbTrxMiddleware,
		logActivityMiddleware,
	}
}

// Setup sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
