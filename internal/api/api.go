package api

import (
	"context"

	"github.com/katallaxie/prompts"
	"github.com/katallaxie/prompts/ollama"
	"github.com/katallaxie/prompts/perplexity"
)

// Provider is the provider for the model.
type Provider string

const (
	// ProviderOllama is the provider for the Ollama model.
	ProviderOllama Provider = "ollama"
	// ProviderPerplexity is the provider for the Perplexity model.
	ProviderPerplexity Provider = "perplexity"
)

// ClientFactory is a factory for creating a new client.
func ClientFactory(provider, model, url, key string) prompts.Chat {
	switch Provider(provider) {
	case ProviderOllama:
		return ollama.New(ollama.WithBaseURL(url))
	case ProviderPerplexity:
		return perplexity.New(perplexity.WithApiKey(key))
	default:
		return nil
	}
}

// Api is the API for the application.
type Api struct {
	chat prompts.Chat
}

// NewApi returns a new api.
func NewApi(chat prompts.Chat) *Api {
	return &Api{
		chat: chat,
	}
}

// CreatePrompt creates a new prompt.
func (a *Api) CreatePrompt(ctx context.Context, model, p string, cb ...func(res *prompts.ChatCompletionResponse) error) error {
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

	req := prompts.NewStreamChatCompletionRequest()
	req.SetModel(model)
	req.AddMessages(msg...)

	err := a.chat.SendStreamCompletionRequest(ctx, req, cb...)
	if err != nil {
		return err
	}

	return nil
}
