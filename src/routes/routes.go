package routes

import (
	"ambassador/src/controllers"
	"ambassador/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	// ADMIN ROUTE
	admin := api.Group("admin")
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)

	//Users
	adminAuthenticated.Get("/user", controllers.User)
	adminAuthenticated.Post("/logout", controllers.Logout)
	adminAuthenticated.Post("/updateInfo", controllers.UpdateInfo)
	adminAuthenticated.Post("/updatePassword", controllers.UpdatePassword)
	adminAuthenticated.Get("/ambassadors", controllers.Ambassadors)
	adminAuthenticated.Get("/populateAmbassadors", controllers.PopulateAmbassadors)

	// Products
	adminAuthenticated.Get("/products", controllers.Products)
	adminAuthenticated.Post("/products", controllers.CreateProducts)
	adminAuthenticated.Get("/products/:id", controllers.GetProduct)
	adminAuthenticated.Put("/products/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("/products/:id", controllers.DeleteProduct)
	adminAuthenticated.Get("/productsPopulate", controllers.PopulateProducts)

	// LINKS
	adminAuthenticated.Get("/user/:id/links", controllers.Link)

	// ORDERS
	adminAuthenticated.Get("/orders", controllers.Orders)
	adminAuthenticated.Get("/generateOrders", controllers.GenerateOrders)

	// AMBASSADORS
	ambassador := api.Group("ambassador")
	ambassador.Post("/register", controllers.Register)
	ambassador.Post("/login", controllers.Login)
	ambassador.Get("/products/frontend", controllers.ProductsFrontend)
	ambassador.Get("/products/backend", controllers.ProductsBackend)

	ambassadorAuthenticated := ambassador.Use(middlewares.IsAuthenticated)
	ambassadorAuthenticated.Get("/user", controllers.User)
	ambassadorAuthenticated.Post("/logout", controllers.Logout)
	ambassadorAuthenticated.Put("/users/info", controllers.UpdateInfo)
	ambassadorAuthenticated.Put("/users/password", controllers.UpdatePassword)

}
