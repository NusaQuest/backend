package middlewares

import (
	"context"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
)

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