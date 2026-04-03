package routes

import (
		"simple_crud/Controllers"
		"simple_crud/middleware"

		"github.com/gofiber/fiber/v2"
		"github.com/uptrace/bun"
	)

	func Setup(app *fiber.App, db *bun.DB) {

	userHandler := controllers.NewUserHandler(db)
	categoryHandler := controllers.NewCategoryHandler(db)


	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
	app.Post("/add-category", categoryHandler.CreateCategory)
	
	app.Get("/", fiber.Handler(func(c *fiber.Ctx) error {
		return c.SendString("Backend")
	}))

	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.Post("/logout", userHandler.Logout)
	api.Get("/users", userHandler.GetCurrentUser)
	api.Put("/users", userHandler.UpdateUser)

}	