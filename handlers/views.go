package handlers

import "github.com/gofiber/fiber/v2"

// NotFound renders a 404 page.
// @notice Called when a route is not found.
// @param c Fiber context for the request.
// @return Rendered "not_found" page with 404 status.
func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).Render("not_found", fiber.Map{})
}

// Welcome renders the welcome page.
// @notice Called for the root or home page.
// @param c Fiber context for the request.
// @return Rendered "welcome" page with 200 OK status.
func Welcome(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("welcome", fiber.Map{})
}
