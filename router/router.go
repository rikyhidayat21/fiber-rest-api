package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rikyhidayat21/fiber-rest-api/router/noteRoute"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())


	// setup the node routes
	noteRoute.SetupNoteRoutes(api)
}