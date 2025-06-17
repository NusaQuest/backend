package controller

import (
	"context"
	"time"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/constant"
	"github.com/NusaQuest/backend.git/model"
	"github.com/NusaQuest/backend.git/output"
	"github.com/NusaQuest/backend.git/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ConnectWallet(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var auth *model.Auth
	err := c.BodyParser(&auth)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	err = utils.GetValidator().Struct(auth) 
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	collection := config.GetDatabase().Collection("auths")
	_, err = collection.InsertOne(ctx, auth)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToInsertData))
	}

	return output.GetSuccess(c, string(constant.SuccessPostData), fiber.Map{
		"auth": auth,
	})

}

func DisconnectWallet(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wallet := c.Params("wallet")

	collection := config.GetDatabase().Collection("auths")
	filter := bson.M{
		"wallet": wallet,
	}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToDeleteData))
	}

	return output.GetSuccess(c, string(constant.SuccessDeleteData), nil)

}