package api

import (
	"context"

	"github.com/katallaxie/prompts"
	"github.com/katallaxie/prompts/ollama"
)

type Api struct {
	client *ollama.Ollama
}

// NewApi returns a new api.
func NewApi(client *ollama.Ollama) *Api {
	return &Api{
		client: client,
	}
}

// CreatePrompt creates a new prompt.
func (a *Api) CreatePrompt(ctx context.Context, p string, cb ...func(res *prompts.ChatCompletionResponse) error) error {
	msg := []prompts.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "You are a helpful assistant. You start every answers with 'Sure!'",
		},
		{
			Role:    "user",
			Content: p,
		},
	}

	req := ollama.NewStreamCompletionRequest()
	req.AddMessages(msg...)

	err := a.client.SendStreamCompletionRequest(context.Background(), req, cb...)
	if err != nil {
		return err
	}

	return nil
}
