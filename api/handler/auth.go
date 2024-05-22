package handler

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"result": "OK"})
}
