package controllers

import (
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/controllers/helper"
	"github.com/NusaQuest/backend.git/models"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// RegisterIdentity handles POST /api/identities
// @notice Stores a new identity (hashed KTP registration) associated with a wallet.
// @param c Fiber context containing the identity data in the request body.
// @return JSON response with insert result or error.
func RegisterIdentity(c *fiber.Ctx) error {
	var identity models.Identity

	res, err := helper.InsertData(c, string(constants.Identities), &identity)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res,
	})
}

// GetIdentity handles GET api/identities/:wallet
// @notice Retrieves a registered identity associated with the given wallet address.
// @param c Fiber context with the wallet parameter.
// @return JSON response with the identity data or error.
func GetIdentity(c *fiber.Ctx) error {
	wallet := c.Params("wallet")

	var identites []models.Identity
	filter := bson.M{"wallet": wallet}

	_, err := helper.RetrieveData(filter, string(constants.Identities), &identites)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	if len(identites) == 0 {
		return output.GetError(c, fiber.StatusNotFound, string(constants.IdentityNotFound))
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"identity": identites[0],
	})
}
