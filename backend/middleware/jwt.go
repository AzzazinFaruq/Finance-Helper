package middleware

import (
	"simple_crud/Utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tokenString string

		authHeader := c.Get("Authorization")
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			tokenString = c.Cookies("Authorization")
		}

		if tokenString == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Token not found"})
		}

		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["sub"].(float64))

		c.Locals("user", userID)

		return c.Next()
	}
}