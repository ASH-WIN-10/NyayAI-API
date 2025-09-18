package ai

import (
	"context"
	"time"

	"google.golang.org/genai"
)

func (c *AIClient) GenerateLegalAdvice(prompt string) (string, error) {
	systemInstruction := `
        You are a helpful legal assistant.
        Provide clear and concise legal advice based on the user's prompt.
        Ensure that your advice is accurate and easy to understand.
        Don't provide any disclaimers or additional information beyond the legal advice requested.
        Dont't use technical terms or jargon that a layperson might not understand.
        Keep your responses brief and to the point.
    `
	temperature := float32(0.3)
	prompt = prompt + "\n\nDo not mention their region or country name unless explicitly asked like (In India, In USA etc.). "

	config := &genai.GenerateContentConfig{
		MaxOutputTokens:   500,
		Temperature:       &temperature,
		SystemInstruction: genai.NewContentFromText(systemInstruction, ""),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	result, err := c.client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		config,
	)

	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
