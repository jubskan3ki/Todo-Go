package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	log.Printf("Started %s %s", c.Method(), c.Path())

	err := c.Next()

	log.Printf("Completed in %v", time.Since(start))
	return err
}
