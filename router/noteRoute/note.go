package noteRoute

import (
	"github.com/gofiber/fiber/v2"
	notehandler "github.com/rikyhidayat21/fiber-rest-api/handler/noteHandler"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	note.Post("/", notehandler.CreateNotes)

	note.Get("/", notehandler.GetNotes)
	
	note.Get("/:noteId", notehandler.GetNote)
	
	note.Put("/:noteId", notehandler.UpdateNote)
	
	note.Delete("/:noteId", notehandler.DeleteNote)
}