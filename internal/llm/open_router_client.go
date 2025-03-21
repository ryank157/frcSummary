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
		model:  "google/gemini-2.0-flash-001",
	}
}

const MatchSystemMessage = `Create a summary of the match in the past that follows the format of the following summary.
This is the 2025 First Robotics competition. There are 3 robots on 2 alliances (red, blue) that are facing off against each other.
There are three distinct phases to the match: Autonomous, tele operated, endgame. The autonomous is preprogrammed using sensors and is worht more points compared to teleop
Have a breakdown of the points between the 2 teams at each of the three phases. The prompt will hold the data for the match


The upcoming match #14 should be a victory for team #1778. All
of the teams have only two matches in this first event of the season
so it is easily possible for their to be an upset.

1778 has a strong autonomous routine scoring at least the
equal of 9450 the strongest team on the opposing red alliance. In
teleop team 1778 has dominated. The blue alliances only weakness is
its complete inability to climb. If it is close at the end the red alliance
could win with 9450's climb.

If red gets out to an early lead it could be exciting.`

func (c *OpenRouterClient) Generate(prompt string) (string, error) {
	completion := c.client.
		NewChatCompletion().
		WithDebug(true).    // Enable debug mode to see the request and response in the console
		WithModel(c.model). // Change the model if you want
		WithSystemMessage(MatchSystemMessage).
		WithUserMessage(prompt)

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
