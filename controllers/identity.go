package controllers

import (
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/controllers/helper"
	"github.com/NusaQuest/backend.git/models"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

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

func GetIdentity(c *fiber.Ctx) error {

	wallet := c.Params("wallet")

	var identites []models.Identity
	filter := bson.M{"wallet": wallet}

	_, err := helper.RetrieveData(filter, string(constants.Identities), &identites)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"identity": identites[0],
	})

}
