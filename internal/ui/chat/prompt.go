package chat

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/katallaxie/m/internal/state"
	"github.com/katallaxie/m/internal/ui"
	"github.com/katallaxie/pkg/slices"
	"github.com/katallaxie/prompts"
	"github.com/katallaxie/prompts/ollama"
	"github.com/katallaxie/streams"
	"github.com/katallaxie/streams/sinks"
	"github.com/katallaxie/streams/sources"
	"github.com/rivo/tview"
)

const (
	example = `Write a concise summary of the following in less then 20 words:
	
	"Artificial intelligence (AI) is technology that enables computers and machines to simulate human learning, comprehension, problem solving, decision making, creativity and autonomy."
	
	CONCISE SUMMARY:`
)

const defaultPrompt = "🐶 >"

type PromptState struct {
	isFocused bool
}

// Prompt is a chat prompt.
type Prompt struct {
	*tview.TextArea
	state *PromptState
	app   ui.Application
}

func mapCompletionMessages(msg prompts.Completion) string {
	f := slices.First(msg.Choices...)
	return f.Message.GetContent()
}

// NewPrompt returns a new chat prompt.
func NewPrompt(app ui.Application) *Prompt {
	prompt := &Prompt{
		TextArea: tview.NewTextArea(),
		state: &PromptState{
			isFocused: false,
		},
		app: app,
	}

	prompt.SetTitle(" ✍️ Prompt ")
	prompt.SetBorder(true)
	prompt.SetWordWrap(true)
	prompt.SetWrap(true)
	prompt.SetBorderPadding(1, 1, 1, 1)
	prompt.SetPlaceholder("Enter your message here...")
	prompt.SetInputCapture(prompt.onInputCapture)

	prompt.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			prompt.onEnter()
		}

		return event
	})

	return prompt
}

func (p *Prompt) onEnter() {
	api, err := ollama.New(ollama.WithBaseURL("http://localhost:7869"), ollama.WithModel("smollm"))
	if err != nil {
		panic(err)
	}

	prompt := prompts.Prompt{
		Model: prompts.Model("smollm"),
		Messages: []prompts.Message{
			&prompts.UserMessage{
				Content: example,
			},
		},
	}

	res, err := api.Complete(context.Background(), &prompt)
	if err != nil {
		panic(err)
	}

	in := make(chan any, 1)

	source := sources.NewChanSource(res)
	sink := sinks.NewChanSink(in)

	go func() {
		for msg := range in {
			p.app.GetState().Dispatch(state.NewAddMessage(msg.(string)))
		}
	}()

	source.Pipe(streams.NewPassThrough()).Pipe(streams.NewMap(mapCompletionMessages)).To(sink)
}

func (p *Prompt) onInputCapture(event *tcell.EventKey) *tcell.EventKey {
	return event
}

// GetIsFocused returns the prompt focus state.
func (p *Prompt) GetIsFocused() bool {
	return p.state.isFocused
}

// SetIsFocused sets the prompt focus state.
func (p *Prompt) SetIsFocused(focused bool) {
	p.state.isFocused = focused
}

// Deactivate deactivates the prompt.
func (p *Prompt) Deactivate() {}
