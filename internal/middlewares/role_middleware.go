package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func OnlyCategory(category string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*map[string]interface{})
		if (*user)["category"] != category {
			return c.Status(403).JSON(fiber.Map{
				"success": false,
				"message": "Forbidden for this category",
			})
		}
		return c.Next()
	}
}
