package llm

import (
	"fmt"
	"log"

	openrouter "github.com/eduardolat/openroutergo"
)

type OpenRouterClient struct {
	client *openrouter.Client
	model  string // You can configure the model
}

func NewOpenRouterClient(apiKey string) LLMClient {
	client, err := openrouter.NewClient().WithAPIKey(apiKey).Create()

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &OpenRouterClient{
		client: client,
		model:  "mistralai/mistral-medium",
	}
}

func (c *OpenRouterClient) Generate(prompt string) (string, error) {
	completion := c.client.
		NewChatCompletion().
		WithDebug(true).    // Enable debug mode to see the request and response in the console
		WithModel(c.model). // Change the model if you want
		WithSystemMessage("You are a helpful assistant expert in geography.").
		WithUserMessage("What is the capital of France?")

	_, resp, err := completion.Execute()
	if err != nil {
		log.Fatalf("Failed to execute completion: %v", err)
	}

	fmt.Println("Response:", resp.Choices[0].Message.Content)

	if err != nil {
		log.Printf("Error during OpenRouter ChatCompletion: %v", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
