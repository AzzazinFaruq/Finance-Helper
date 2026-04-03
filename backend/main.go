package main

import (
	"simple_crud/Config"
	"simple_crud/Routes"
	"simple_crud/Setup"
    "simple_crud/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()
	Setup.Connect()
	Setup.Migrate()

	app := fiber.New()
	routes.Setup(app, Setup.DB)
	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.SecurityMiddleware())

	app.Listen(":" + config.Get("APP_PORT"))
}