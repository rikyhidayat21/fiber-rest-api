package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// inisialisasi fiber
	app := fiber.New() 

	// middleware logger
	app.Use(logger.New())

	// Get route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"Hello": "world",
		})
	})

	app.Listen(":3000")
}