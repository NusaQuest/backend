package controller

import (
	"context"
	"time"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/constant"
	"github.com/NusaQuest/backend.git/model"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddDestination(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var destination model.Destination
	err := c.BodyParser(&destination)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	collection := config.GetDatabase().Collection("destinations")
	_, err = collection.InsertOne(ctx, destination)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToInsertData))
	}

	return output.GetSuccess(c, string(constant.SuccessPostData), fiber.Map{
		"destination": destination,
	})

}

func GetDestinations(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var destinations []model.Destination

	collection := config.GetDatabase().Collection("destinations")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToRetrieveData))
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &destinations)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToDecodeData))
	}

	return output.GetSuccess(c, string(constant.SuccessReturnData), fiber.Map{
		"destinations": destinations,
	})

}