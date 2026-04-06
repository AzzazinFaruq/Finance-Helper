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
	transactionHandler := controllers.NewTransactionHandler(db)


	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
	
	
	app.Get("/", fiber.Handler(func(c *fiber.Ctx) error {
		return c.SendString("Backend")
	}))

	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.Post("/add-category", categoryHandler.CreateCategory)
	api.Get("/get-category", categoryHandler.GetCategory)
	api.Put("/update-category/:id", categoryHandler.UpdateCategory)
	api.Delete("/delete-category/:id", categoryHandler.DeleteCategory)

	api.Post("/add-transaction", transactionHandler.CreateTransaction)
	api.Get("/get-transaction", transactionHandler.GetTransaction)
	// api.Put("/update-transaction/:id", transactionHandler.UpdateTransaction)
	// api.Delete("/delete-transaction/:id", transactionHandler.DeleteTransaction)

	api.Post("/logout", userHandler.Logout)
	api.Get("/users", userHandler.GetCurrentUser)
	api.Put("/users", userHandler.UpdateUser)

}	