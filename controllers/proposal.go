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

// AddProposal handles POST /proposals
// @notice Adds a new proposal to the database.
// @param c Fiber context containing the request body.
// @return JSON response with insert result or error.
func AddProposal(c *fiber.Ctx) error {

	var proposal models.Proposal

	res, err := helper.InsertData(c, string(constants.Proposals), &proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res,
	})

}

// UpdateProposal handles PUT/PATCH /proposals/:id
// @notice Updates an existing proposal by its ID.
// @param c Fiber context with the proposal ID and body.
// @return JSON response with update result or error.
func UpdateProposal(c *fiber.Ctx) error {

	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	var proposal models.Proposal

	res, err := helper.UpdateData(c, string(constants.Proposals), objId, &proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessUpdateMessage), fiber.Map{
		"result": res,
	})

}

// GetProposals handles GET /proposals
// @notice Retrieves all proposals from the database.
// @return JSON response with the list of proposals or error.
func GetProposals(c *fiber.Ctx) error {

	var proposals []models.Proposal

	_, err := helper.RetrieveData(bson.M{}, string(constants.Proposals), &proposals)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"proposals": proposals,
	})

}

// GetWalletProposals handles GET /proposals/:wallet
// @notice Retrieves proposals associated with a specific wallet address.
// @param c Fiber context with the wallet parameter.
// @return JSON response with filtered proposals or error.
func GetWalletProposals(c *fiber.Ctx) error {

	wallet := c.Params("wallet")

	var proposals []models.Proposal
	filter := bson.M{"wallet": wallet}

	_, err := helper.RetrieveData(filter, string(constants.Proposals), &proposals)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"proposals": proposals,
	})

}
