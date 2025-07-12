package controllers

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/controllers/helper"
	"github.com/NusaQuest/backend.git/models"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
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

	res, err := helper.InsertData(string(constants.Proposals), &proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res.InsertedID,
	})

}

// CheckProposalInput handles POST api/proposals/check
// @notice Validates a beach cleanup proposal using OpenAI's language model.
// @dev Ensures that the provided beach location exists in Indonesia and that the proposal specifically describes a beach cleanup activity.
// @param c Fiber context containing the proposal data in the request body.
// @return JSON response indicating whether the input is valid, or returns an error message if validation fails or input is unclear.
func CheckProposal(c *fiber.Ctx) error {

	var proposal models.Proposal

	err := c.BodyParser(&proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	question := fmt.Sprintf(`
		You are validating a beach cleanup proposal in Indonesia. Analyze the following:

			- Claimed beach: "%s"
			- City: "%s"
			- Province: "%s"
			- Proposal Name: "%s"
			- Description: "%s"

		Determine if BOTH conditions are met:
			1. The location refers to a real beach in Indonesia (not mountains or non-coastal areas).
			2. The proposal clearly describes a beach cleanup (not forest, mountain, or general environmental action).

		If BOTH are clearly true, reply only: true  
		If either one is false or unclear, reply only: false  

		⚠️ Reply with only: true or false — no explanations or extra words.
		`, proposal.BeachName, proposal.City, proposal.Province, proposal.ProposalName, proposal.ProposalDescription)

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(option.WithAPIKey(OPENAI_API_KEY))

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		},
		Model: openai.ChatModelGPT3_5Turbo,
	})
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	boolean, err := strconv.ParseBool(chatCompletion.Choices[0].Message.Content)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	if !boolean {
		return output.GetError(c, fiber.StatusBadRequest, string(constants.ProposalValidationFailed))
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{})

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
