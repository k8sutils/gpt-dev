package openai

import (
	"context"

	openai "github.com/openai/openai-go/v2"
)

// Define the OpenAI API credentials
var (
	openaiAPIKey = "" // Set your OpenAI API key here
	modelID      = "text-davinci-002"
)

// OpenAI is a struct that represents the OpenAI API client
type OpenAI struct {
	Client *openai.Client
}

// NewOpenAI creates a new instance of the OpenAI API client
func NewOpenAI() (*OpenAI, error) {
	client, err := openai.NewClient(openaiAPIKey)
	if err != nil {
		return nil, err
	}

	return &OpenAI{Client: client}, nil
}

// GenerateCode generates code using the OpenAI API
func (o *OpenAI) GenerateCode(ctx context.Context, prompt string) (string, error) {
	resp, err := o.Client.Completions.Create(&openai.CompletionRequest{
		Model:      modelID,
		Prompt:     prompt,
		MaxTokens:  1024,
		Temperature: 0.7,
		N: 1,
		Stop:       "",
	})
	if err != nil {
		return "", err
	}

	// Parse the response and extract the generated code
	code := ""
	for _, choice := range resp.Choices {
		code += choice.Text
	}

	return code, nil
}
