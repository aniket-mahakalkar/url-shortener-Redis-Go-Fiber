package routes

import "github.com/gofiber/fiber/v2"

func SentHealth(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Your server is running"})
}