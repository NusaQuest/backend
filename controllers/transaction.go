package controllers

import (
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/controllers/helper"
	"github.com/NusaQuest/backend.git/models"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddTransaction(c *fiber.Ctx) error {

	var transaction models.Transaction

	res, err := helper.InsertData(c, "transaction", transaction)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res,
	})

}

func GetWalletTransactions(c *fiber.Ctx) error {

	wallet := c.Params("wallet")

	var transactions []models.Transaction
	filter := bson.M{"wallet": wallet}

	res, err := helper.RetrieveData(filter, "transactions", transactions)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"transactions": res,
	})

}
