package handlers

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).Render("not_found", fiber.Map{})
}

func Welcome(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("welcome", fiber.Map{})
}
