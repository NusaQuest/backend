package helper

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func CheckProposalInput(proposalName string, proposalDescription string, beachName string, city string, province string) (string, error) {

	question := fmt.Sprintf(`
		Please check the following input and answer clearly and concisely:
			1. Is there a beach named "%s" located in the city of "%s", province of "%s", Indonesia?
			2. Does the proposal description below indicate a beach cleanup activity?
				Proposal Name: %s
				Proposal Description: %s
				
		Reply in this JSON format:
			{
				"valid": true or false,
				"reason": "brief explanation"
			}			
		`,
		beachName, city, province, proposalName, proposalDescription)

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(option.WithAPIKey(OPENAI_API_KEY))

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		},
		Model: openai.ChatModelGPT3_5Turbo,
	})
	if err != nil {
		return "", err
	}

	return chatCompletion.Choices[0].Message.Content, nil

}
