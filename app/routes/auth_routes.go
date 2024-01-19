package routes

import (
	"Todo-Go/app/api"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App) {
	app.Post("/register", api.RegisterHandler)
	app.Post("/login", api.LoginHandler)
}
