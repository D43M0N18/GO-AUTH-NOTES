package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-auth-notes/models"
	"go-auth-notes/utils"
	"gorm.io/gorm"
)

// CreateNote creates a new note for the authenticated user
func CreateNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if note.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}

	note.UserID = userID
	if err := utils.DB.Create(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create note",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(note)
}

// GetNotes retrieves all notes for the authenticated user
func GetNotes(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var notes []models.Note
	if err := utils.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notes",
		})
	}

	return c.JSON(notes)
}

// GetNote retrieves a specific note by ID
func GetNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	var note models.Note
	if err := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Note not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch note",
		})
	}

	return c.JSON(note)
}

// UpdateNote updates a specific note
func UpdateNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	var note models.Note
	if err := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Note not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch note",
		})
	}

	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if note.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}

	if err := utils.DB.Save(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update note",
		})
	}

	return c.JSON(note)
}

// DeleteNote deletes a specific note
func DeleteNote(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	noteID := c.Params("id")

	result := utils.DB.Where("id = ? AND user_id = ?", noteID, userID).Delete(&models.Note{})
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete note",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
