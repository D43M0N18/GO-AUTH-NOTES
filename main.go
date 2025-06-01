package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "log"
    "os"
    "go-auth-notes/routes"
    "go-auth-notes/utils"
)

func main() {
    // Try to load .env file, but don't fail if it doesn't exist
    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: .env file not found: %v", err)
    }

    if err := utils.ConnectDB(); err != nil {
        log.Fatal("Database connection failed:", err)
    }

    app := fiber.New(fiber.Config{
        AppName: "Go Auth Notes API",
    })

    routes.SetupRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    log.Printf("Server starting on port %s", port)
    log.Fatal(app.Listen("0.0.0.0:" + port))
}
