package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-auth-notes/utils"
	"strings"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required",
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}

		claims, err := utils.ValidateToken(tokenParts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		c.Locals("userID", claims.UserID)
		return c.Next()
	}
}
