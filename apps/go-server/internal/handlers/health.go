package handlers

import "github.com/gofiber/fiber/v3"

func GetHealth(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}
