package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-auth-notes/handlers"
	"go-auth-notes/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Public routes
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Protected routes
	notes := app.Group("/notes", middleware.AuthMiddleware())
	notes.Post("/", handlers.CreateNote)
	notes.Get("/", handlers.GetNotes)
	notes.Get("/:id", handlers.GetNote)
	notes.Put("/:id", handlers.UpdateNote)
	notes.Delete("/:id", handlers.DeleteNote)
}
