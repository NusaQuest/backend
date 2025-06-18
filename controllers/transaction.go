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
)

func AddTransaction(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var transaction models.Transaction
	err := c.BodyParser(&transaction)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constants.FailedToParseData))
	}

	err = utils.GetValidator().Struct(transaction) 
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constants.ValidationError))
	}

	collection := config.GetDatabase().Collection("transactions")
	_, err = collection.InsertOne(ctx, &transaction)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constants.FailedToInsertData))
	}

	return output.GetSuccess(c, string(constants.SuccessPostData), fiber.Map{
		"transaction": transaction,
	})

}

func GetTransactions(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var transactions []models.Transaction

	collection := config.GetDatabase().Collection("transactions")
	cursor, err := collection.Find(ctx, fiber.Map{})
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constants.FailedToRetrieveData))
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &transactions)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constants.FailedToDecodeData))
	}

	return output.GetSuccess(c, string(constants.SuccessReturnData), fiber.Map{
		"transactions": transactions,
	})

}