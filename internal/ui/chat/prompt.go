package chat

import (
	"github.com/rivo/tview"
)

const defaultPrompt = "> [::b]"

// Prompt is a chat prompt.
type Prompt struct {
	*tview.TextArea
}

// NewPrompt returns a new chat prompt.
func NewPrompt() *Prompt {
	prompt := &Prompt{
		TextArea: tview.NewTextArea(),
	}

	prompt.SetTitle(" ✍️ Prompt ")
	prompt.SetBorder(true)
	prompt.SetWordWrap(true)
	prompt.SetWrap(true)
	prompt.SetBorderPadding(1, 1, 1, 1)
	prompt.SetText("help", true)

	return prompt
}

// Deactivate deactivates the prompt.
func (p *Prompt) Deactivate() {}
