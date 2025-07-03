package controllers

import (
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/controllers/helper"
	"github.com/NusaQuest/backend.git/models"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// AddTransaction handles POST /transactions
// @notice Adds a new transaction to the database.
// @param c Fiber context containing the request body.
// @return JSON response with insert result or error.
func AddTransaction(c *fiber.Ctx) error {

	var transaction models.Transaction

	res, err := helper.InsertData(c, string(constants.Transactions), &transaction)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res,
	})

}

// GetWalletTransactions handles GET /transactions/:wallet
// @notice Retrieves transactions based on wallet address.
// @param c Fiber context with the wallet parameter.
// @return JSON response with the list of transactions or error.
func GetWalletTransactions(c *fiber.Ctx) error {

	wallet := c.Params("wallet")

	var transactions []models.Transaction
	filter := bson.M{"wallet": wallet}

	_, err := helper.RetrieveData(filter, string(constants.Transactions), &transactions)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"transactions": transactions,
	})

}
