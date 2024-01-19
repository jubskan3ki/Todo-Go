package routes

import (
	"Todo-Go/app/api"
	"Todo-Go/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterTodoRoutes(app *fiber.App) {
	app.Post("/todo", middleware.AuthMiddleware, api.CreateTodoHandler)
	app.Get("/todos/:id?", middleware.AuthMiddleware, api.GetTodosHandler)
	app.Put("/todo/:id", middleware.AuthMiddleware, api.UpdateTodoHandler)
	app.Delete("/todo/:id", middleware.AuthMiddleware, api.DeleteTodoHandler)
}
