package helper

import (
	"context"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func CheckProposalInput() (string, error) {

	question := ``

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