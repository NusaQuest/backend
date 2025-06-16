package middleware

import (
	"context"

	"github.com/NusaQuest/backend.git/config"
	"github.com/gofiber/fiber/v2"
)

func CheckDBConnection(c *fiber.Ctx) error {
	
	if config.Client == nil {
		
	}

	err := config.Client.Ping(context.Background(), nil)
	if err != nil {

	}
	
	return c.Next()

}