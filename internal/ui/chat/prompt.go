package chat

import (
	"github.com/katallaxie/m/internal/api"
	"github.com/katallaxie/m/internal/store"
	"github.com/katallaxie/m/internal/ui"
	"github.com/katallaxie/prompts"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// PromptState holds the state of the prompt.
type PromptState struct {
	isFocused bool
}

// Prompt is a chat prompt.
type Prompt struct {
	*tview.TextArea
	state *PromptState
	api   *api.Api
	app   ui.Application[store.State]
}

// NewPrompt returns a new chat prompt.
func NewPrompt(app ui.Application[store.State], api *api.Api) *Prompt {
	prompt := &Prompt{
		TextArea: tview.NewTextArea(),
		state: &PromptState{
			isFocused: false,
		},
		app: app,
		api: api,
	}

	prompt.SetTitle(" ✍️ Prompt ctrl-p ")
	prompt.SetBorder(true)
	prompt.SetWordWrap(true)
	prompt.SetWrap(true)
	prompt.SetBorderPadding(1, 1, 1, 1)
	prompt.SetPlaceholder("Enter your message here...")
	prompt.SetInputCapture(prompt.onInputCapture)

	prompt.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			prompt.onEnter(prompt.GetText())
		}

		return event
	})

	return prompt
}

func (p *Prompt) onEnter(prompt string) {
	go func() {
		p.app.GetStore().Dispatch(store.NewSetStatus(store.Loading))

		fn := func(res *prompts.ChatCompletionResponse) error {
			p.app.GetStore().Dispatch()

			return nil
		}

		_ = p.api.CreatePrompt(p.app.Context(), p.app.Config().Spec.Api.Model, prompt, fn)
		p.app.GetStore().Dispatch(store.NewSetStatus(store.Success))
	}()
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
