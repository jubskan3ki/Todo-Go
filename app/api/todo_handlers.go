package api

import (
	"Todo-Go/app/db"
	"Todo-Go/app/model"

	"github.com/gofiber/fiber/v2"
)

func CreateTodoHandler(c *fiber.Ctx) error {
	var todo model.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := c.Locals("user").(*model.User)
	todo.UserID = user.ID

	db.DB.Create(&todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func GetTodosHandler(c *fiber.Ctx) error {
	todoID := c.Params("id")
	user := c.Locals("user").(*model.User)

	var todos []model.Todo

	if todoID == "" {
		db.DB.Where("user_id = ?", user.ID).Find(&todos)
	} else {
		db.DB.Where("id = ? AND user_id = ?", todoID, user.ID).First(&todos)
	}

	return c.JSON(todos)
}

func UpdateTodoHandler(c *fiber.Ctx) error {
	todoID := c.Params("id")
	user := c.Locals("user").(*model.User)

	var todo model.Todo
	if err := db.DB.Where("id = ? AND user_id = ?", todoID, user.ID).First(&todo).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db.DB.Save(&todo)

	return c.JSON(todo)
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	todoID := c.Params("id")
	user := c.Locals("user").(*model.User)

	var todo model.Todo
	if err := db.DB.Where("id = ? AND user_id = ?", todoID, user.ID).First(&todo).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	db.DB.Delete(&todo)

	return c.SendStatus(fiber.StatusNoContent)
}
