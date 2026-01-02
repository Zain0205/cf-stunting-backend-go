package utils

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, data any) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func Error(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}
