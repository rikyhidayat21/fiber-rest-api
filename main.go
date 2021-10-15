package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rikyhidayat21/fiber-rest-api/database"
)

func main() {
	// inisialisasi fiber
	app := fiber.New() 

	// Connect to the Database
	database.ConnectDB()

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