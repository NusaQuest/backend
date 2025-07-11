package helper

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// CheckProposalInput validates a beach cleanup proposal using OpenAI's language model.
// @notice Checks if the provided beach name and location are valid, and if the proposal content is related to beach cleanup.
// @dev Sends a structured prompt to OpenAI and parses the JSON response for UI-friendly validation feedback.
// @param proposalName The title of the proposed beach cleanup.
// @param proposalDescription The detailed explanation of the proposed activity.
// @param beachName The name of the beach to be cleaned.
// @param city The city where the beach is located.
// @param province The province where the beach is located.
// @return bool Indicates whether the proposal is valid.
// @return error Any error encountered during the API call or JSON parsing.
func CheckProposalInput(proposalName string, proposalDescription string, beachName string, city string, province string) (bool, error) {

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
		`, beachName, city, province, proposalName, proposalDescription)

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(option.WithAPIKey(OPENAI_API_KEY))

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		},
		Model: openai.ChatModelGPT3_5Turbo,
	})
	if err != nil {
		return false, err
	}

	boolean, err := strconv.ParseBool(chatCompletion.Choices[0].Message.Content)
	if err != nil {
		return false, err
	}

	return boolean, nil

}
