package routes

import (
	"github.com/0xsp4c3/core/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	// route.Get("/books", controllers.GetBooks)   // get list of all books
	// route.Get("/book/:id", controllers.GetBook) // get one book by ID
    // Coins
    route.Get("/coins", controllers.GetCoins)
    route.Get("/coin/:id", controllers.GetCoin)
    route.Get("/coinsbyexchangeid/:id", controllers.GetCoinsByExchangeID)
    // Exchanges
    route.Get("/exchanges", controllers.GetExchanges)
    route.Get("/exchange/:id", controllers.GetExchange)

	// Routes for POST method:
	route.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	route.Post("/user/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens
}
