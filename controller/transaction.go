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

func AddTransaction(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var transaction model.Transaction
	err := c.BodyParser(&transaction)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	collection := config.GetDatabase().Collection("transactions")
	_, err = collection.InsertOne(ctx, &transaction)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToInsertData))
	}

	return output.GetSuccess(c, string(constant.SuccessPostData), fiber.Map{
		"transaction": transaction,
	})

}

func GetTransactions(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var transactions []model.Transaction

	collection := config.GetDatabase().Collection("transactions")
	cursor, err := collection.Find(ctx, fiber.Map{})
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToRetrieveData))
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &transactions)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, string(constant.FailedToDecodeData))
	}

	return output.GetSuccess(c, string(constant.SuccessReturnData), fiber.Map{
		"transactions": transactions,
	})

}