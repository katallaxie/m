package chat

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const defaultPrompt = "🐶 >"

type PromptState struct {
	isFocused bool
}

// Prompt is a chat prompt.
type Prompt struct {
	*tview.TextArea
	state *PromptState
}

// NewPrompt returns a new chat prompt.
func NewPrompt() *Prompt {
	prompt := &Prompt{
		TextArea: tview.NewTextArea(),
		state: &PromptState{
			isFocused: false,
		},
	}

	prompt.SetTitle(" ✍️ Prompt ")
	prompt.SetBorder(true)
	prompt.SetWordWrap(true)
	prompt.SetWrap(true)
	prompt.SetBorderPadding(1, 1, 1, 1)
	prompt.SetPlaceholder("Enter your message here...")
	prompt.SetInputCapture(prompt.onInputCapture)

	prompt.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})

	return prompt
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
