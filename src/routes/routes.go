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
	adminAuthenticated.Get("/user", controllers.User)
	adminAuthenticated.Post("/logout", controllers.Logout)
	adminAuthenticated.Post("/updateInfo", controllers.UpdateInfo)
	adminAuthenticated.Post("/updatePassword", controllers.UpdatePassword)
	adminAuthenticated.Get("/ambassadors", controllers.Ambassadors)
	adminAuthenticated.Get("/populateAmbassadors", controllers.PopulateAmbassadors)
}
