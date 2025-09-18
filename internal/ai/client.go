package ai

import (
	"context"

	"google.golang.org/genai"
)

type AIClient struct {
	client *genai.Client
}

func NewAIClient(apiKey string) (*AIClient, error) {
	config := &genai.ClientConfig{
		APIKey: apiKey,
	}

	client, err := genai.NewClient(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &AIClient{client}, nil
}
