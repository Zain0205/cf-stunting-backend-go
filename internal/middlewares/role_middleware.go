package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func OnlyCategory(category string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Tambahkan pengecekan 'ok' agar lebih aman
		user, ok := c.Locals("user").(*map[string]any)

		if !ok || (*user)["category"] != category {
			return c.Status(403).JSON(fiber.Map{
				"success": false,
				"message": "Forbidden for this category",
			})
		}
		return c.Next()
	}
}
