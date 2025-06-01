package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-auth-notes/models"
	"go-auth-notes/utils"
)

func Register(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Validate input
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, email, and password are required",
		})
	}

	// Check if user already exists
	var existingUser models.User
	if err := utils.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email already registered",
		})
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := utils.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Validate input
	if input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	var user models.User
	if err := utils.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if err := user.CheckPassword(input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func GetUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var user models.User
	if err := utils.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"email": user.Email,
	})
}

func Logout(c *fiber.Ctx) error {
	// In a real application, you might want to blacklist the token
	// For now, we'll just return a success message
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
