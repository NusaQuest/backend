package middlewares

import (
	"context"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
)

// CheckDBConnection is a middleware to validate MongoDB connection status.
// @notice Ensures the MongoDB client is initialized and reachable before proceeding.
// @param c Fiber context for the current HTTP request.
// @return Proceeds to the next middleware/handler if connected, or returns 500 error if not.
func CheckDBConnection(c *fiber.Ctx) error {

	if config.Client == nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constants.DatabaseNotInitialized))
	}

	err := config.Client.Ping(context.Background(), nil)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constants.ErrorWhilePingToDatabase))
	}

	return c.Next()
}
