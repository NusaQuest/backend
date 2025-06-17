package controller

import (
	"context"
	"time"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/constant"
	"github.com/NusaQuest/backend.git/model"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
)

func ConnectWallet(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req *model.AuthRequest
	err := c.BodyParser(&req)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	collection := config.GetDatabase().Collection("auth_requests")
	_, err = collection.InsertOne(ctx, req)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	return output.GetSuccess(c, "", fiber.Map{
		"auth_request": req,
	})

}