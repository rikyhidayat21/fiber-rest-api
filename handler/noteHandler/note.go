package notehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rikyhidayat21/fiber-rest-api/database"
	"github.com/rikyhidayat21/fiber-rest-api/model"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	// find all notes in the db
	db.Find(&notes)

	// if no note is present return an error
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "No notes present",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Notes found",
		"data": notes,
	})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// store the body in the note and return error if encountered
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Review your input",
			"data": err,
		})
	}

	// Add a uuid to the note
	note.ID = uuid.New()

	// create the note and return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Could not create note",
			"data": err,
		})
	}

	// return the created note
	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Created Note",
		"data": note,
	})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// read the param noteId
	id := c.Params("noteId")

	// find the note with the given id
	db.Find(&note, "id = ?", id)

	// if not found, return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "Not found!",
			"data": nil,
		})
	}

	// return the note with id
	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Note found",
		"data": note,
	})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title			string `json:"title"`
		SubTitle		string `json:"subTitle"`
		Text			string `json:"text"`
	}

	db := database.DB
	var note model.Note

	// read the param
	id := c.Params("noteId")

	// find the note with the given id
	db.Find(&note, "id = ?", id)

	// if no such note, return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "Not found!",
			"data": nil,
		})
	}

	// store the body containing the updated data and return error if encourage
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Review your input",
			"data": err,
		})
	}

	// edit the note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	// save the changes
	db.Save(&note)

	// return the updated note
	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Notes found",
		"data": note,
	})
}

func DeleteNote(c *fiber.Ctx) error {
    db := database.DB
    var note model.Note

    // Read the param noteId
    id := c.Params("noteId")

    // Find the note with the given Id
    db.Find(&note, "id = ?", id)

    // If no such note present return an error
    if note.ID == uuid.Nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
    }

    // Delete the note and return error if encountered
    err := db.Delete(&note, "id = ?", id).Error

    if err != nil {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
    }

    // Return success message
    return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}