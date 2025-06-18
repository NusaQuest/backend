package controllers

import (
	"context"
	"time"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/models"
	"github.com/NusaQuest/backend.git/output"
	"github.com/NusaQuest/backend.git/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddDestination(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var destination models.Destination
	err := c.BodyParser(&destination)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constants.FailedToParseData))
	}

	err = utils.GetValidator().Struct(destination) 
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constants.ValidationError))
	}

	collection := config.GetDatabase().Collection("destinations")
	_, err = collection.InsertOne(ctx, destination)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constants.FailedToInsertData))
	}

	return output.GetSuccess(c, string(constants.SuccessPostData), fiber.Map{
		"destination": destination,
	})

}

func GetDestinations(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var destinations []models.Destination

	collection := config.GetDatabase().Collection("destinations")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constants.FailedToRetrieveData))
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &destinations)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constants.FailedToDecodeData))
	}

	return output.GetSuccess(c, string(constants.SuccessReturnData), fiber.Map{
		"destinations": destinations,
	})

}