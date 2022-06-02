package routes

import (
	"github.com/0xsp4c3/core/app/controllers"
	"github.com/0xsp4c3/core/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)           // create a new book
    route.Post("/coin", middleware.JWTProtected(), controllers.CreateCoin)
    route.Post("/exchange", middleware.JWTProtected(), controllers.CreateExchange)

	route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens

	// Routes for PUT method:
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID
    route.Put("/coin", middleware.JWTProtected(), controllers.UpdateCoin)
    route.Put("/exchange", middleware.JWTProtected(), controllers.UpdateExchange)

	// Routes for DELETE method:
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
    route.Delete("/coin", middleware.JWTProtected(), controllers.DeleteCoin)
    route.Delete("/exchange", middleware.JWTProtected(), controllers.DeleteExchange)
}
