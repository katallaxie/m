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
func (a *Api) CreatePrompt(ctx context.Context, p string) (chan any, error) {
	prompt := prompts.Prompt{
		Model: prompts.Model("smollm"),
		Messages: []prompts.Message{
			&prompts.UserMessage{
				Content: p,
			},
		},
	}

	// res, err := a.client.Complete(context.Background(), &prompt)
	// if err != nil {
	// 	panic(err)
	// }
	// source := sources.NewChanSource(res)
	// sink := sinks.NewFSMStore(store, completionAction)

	// source.Pipe(streams.NewPassThrough()).Pipe(streams.NewMap(mapCompletionMessages)).To(sink)

	return a.client.Complete(context.Background(), &prompt)
}
