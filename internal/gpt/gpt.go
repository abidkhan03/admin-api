package gpt

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func (c *Client) request(prompt string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Temperature: 0,

		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	res, err := c.c.CreateChatCompletion(c.ctx, req)
	if err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("no response")
	}

	finishReason := res.Choices[0].FinishReason

	if finishReason == "length" {
		return "", fmt.Errorf("input too long")
	}

	if finishReason != "stop" {
		return "", fmt.Errorf("incomplete response with reason: %s", finishReason)
	}

	return res.Choices[0].Message.Content, nil
}
