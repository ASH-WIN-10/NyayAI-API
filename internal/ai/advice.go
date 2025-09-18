package ai

import (
	"context"

	"google.golang.org/genai"
)

func (c *AIClient) GenerateLegalAdvice(prompt string) (string, error) {
	result, err := c.client.Models.GenerateContent(
		context.Background(),
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	return result.Candidates[0].Content.Parts[0].Text, nil
}
