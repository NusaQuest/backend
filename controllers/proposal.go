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

// AddProposal handles POST api/proposals
// @notice Adds a new proposal to the database with AI-based validation.
// @dev Parses request body, uses OpenAI to validate location and description context, inserts if valid.
// @param c Fiber context containing the request body.
// @return JSON response with insert result or error.
func AddProposal(c *fiber.Ctx) error {

	var proposal models.Proposal

	err := c.BodyParser(&proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	valid, err := helper.CheckProposalInput(proposal.ProposalName, proposal.ProposalDescription, proposal.BeachName, proposal.City, proposal.Province)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	if !valid {
		return output.GetError(c, fiber.StatusBadRequest, string(constants.ProposalValidationFailed))
	}

	res, err := helper.InsertData(string(constants.Proposals), &proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res.InsertedID,
	})

}

// UpdateProposal handles PUT/PATCH api/proposals/:id
// @notice Updates an existing proposal by its ID with AI-based validation..
// @dev Parses request body, uses OpenAI to validate location and description context, updates if valid.
// @param c Fiber context with the proposal ID and body.
// @return JSON response with update result or error.
func UpdateProposal(c *fiber.Ctx) error {

	id := c.Params("id")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	var proposal models.Proposal

	err = c.BodyParser(&proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	valid, err := helper.CheckProposalInput(proposal.ProposalName, proposal.ProposalDescription, proposal.BeachName, proposal.City, proposal.Province)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	if !valid {
		return output.GetError(c, fiber.StatusBadRequest, string(constants.ProposalValidationFailed))
	}

	_, err = helper.UpdateData(string(constants.Proposals), "_id", objId, &proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessUpdateMessage), fiber.Map{})

}

// DeleteProposal handles DELETE /api/proposals/:id
// @notice Deletes a proposal based on the provided ID.
// @param id The ObjectID of the proposal to be deleted (from URL path).
// @return JSON response indicating success or failure of the deletion.
func DeleteProposal(c *fiber.Ctx) error {

	id := c.Params("id")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	_, err = helper.DeleteData(string(constants.Proposals), "_id", objId)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessDeleteMessage), fiber.Map{})

}

// GetProposals handles GET api/proposals
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
