package api

import (
	"context"

	"github.com/katallaxie/m/internal/state"
	"github.com/katallaxie/streams"

	"github.com/katallaxie/pkg/fsmx"
	"github.com/katallaxie/pkg/slices"
	"github.com/katallaxie/prompts"
	"github.com/katallaxie/prompts/ollama"
	"github.com/katallaxie/streams/sinks"
	"github.com/katallaxie/streams/sources"
)

func mapCompletionMessages(msg prompts.Completion) string {
	f := slices.First(msg.Choices...)
	return f.Message.GetContent()
}

func completionAction(msg any) fsmx.Action {
	return state.NewAddMessage(msg.(string))
}

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
func (a *Api) CreatePrompt(ctx context.Context, store fsmx.Store[state.State], p string) error {
	prompt := prompts.Prompt{
		Model: prompts.Model("smollm"),
		Messages: []prompts.Message{
			&prompts.UserMessage{
				Content: p,
			},
		},
	}

	res, err := a.client.Complete(context.Background(), &prompt)
	if err != nil {
		panic(err)
	}
	source := sources.NewChanSource(res)
	sink := sinks.NewFSMStore(store, completionAction)

	source.Pipe(streams.NewPassThrough()).Pipe(streams.NewMap(mapCompletionMessages)).To(sink)

	return nil
}
