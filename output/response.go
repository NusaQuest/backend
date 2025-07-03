package output

import "github.com/gofiber/fiber/v2"

// @notice Returns a standardized JSON error response to the client.
// @param c Fiber context for the current request.
// @param status HTTP status code to return (e.g. 400, 500).
// @param err Error message to include in the response.
func GetError(c *fiber.Ctx, status int, err string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": err,
	})
}

// @notice Returns a standardized JSON success response with optional data.
// @param c Fiber context for the current request.
// @param msg Success message to include in the response.
// @param data Optional key-value map containing response data.
func GetSuccess(c *fiber.Ctx, msg string, data fiber.Map) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": msg,
		"data":    data,
	})
}
