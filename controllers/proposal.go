package controllers

import (
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/controllers/helper"
	"github.com/NusaQuest/backend.git/models"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProposal(c *fiber.Ctx) error {

	var proposal models.Proposal

	res, err := helper.InsertData(c, "proposals", proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res,
	})

}

func UpdateProposal(c *fiber.Ctx) error {

	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	var proposal models.Proposal

	res, err := helper.UpdateData(c, objId, proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessUpdateMessage), fiber.Map{
		"result": res,
	})

}

func GetProposals(c *fiber.Ctx) error {

	var proposals []models.Proposal

	res, err := helper.RetrieveData(bson.M{}, "proposals", proposals)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"proposals": res,
	})

}

func GetWalletProposals(c *fiber.Ctx) error {

	wallet := c.Params("wallet")

	var proposals []models.Proposal
	filter := bson.M{"wallet": wallet}

	res, err := helper.RetrieveData(filter, "proposals", proposals)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"proposals": res,
	})

}
