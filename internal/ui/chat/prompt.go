package chat

import (
	"github.com/katallaxie/m/internal/api"
	"github.com/katallaxie/m/internal/effects"
	"github.com/katallaxie/m/internal/state"
	"github.com/katallaxie/m/internal/ui"

	"github.com/gdamore/tcell/v2"
	"github.com/katallaxie/pkg/fsmx"
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
	api   *api.Api
	app   ui.Application[state.State]
}

// NewPrompt returns a new chat prompt.
func NewPrompt(app ui.Application[state.State], api *api.Api) *Prompt {
	prompt := &Prompt{
		TextArea: tview.NewTextArea(),
		state: &PromptState{
			isFocused: false,
		},
		app: app,
		api: api,
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
			prompt.onEnter(example)
		}

		return event
	})

	return prompt
}

func (p *Prompt) onEnter(prompt string) {
	fsmx.Effect(p.app.GetStore(), effects.IsLoading())
	go fsmx.Effect(p.app.GetStore(), effects.FetchChatCompletion(p.api, prompt))
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
