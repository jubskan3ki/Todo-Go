package main

import (
	"Todo-Go/app/config"
	"Todo-Go/app/db"
	"Todo-Go/app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()
	db.Init(cfg)
	app := fiber.New()

	routes.RegisterAuthRoutes(app)
	routes.RegisterTodoRoutes(app)

	app.Listen(":8080")
}
