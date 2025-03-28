package ui

import (
	"sync"

	"github.com/rivo/tview"
)

// Prompt captures users free from command input.
type Prompt struct {
	*tview.TextView

	ui      *UI
	noIcons bool
	icon    rune
	spacer  int
	mx      sync.RWMutex
}

// NewPrompt returns a new command view.
func NewPrompt(ui *UI, noIcons bool) *Prompt {
	p := Prompt{
		ui:       ui,
		noIcons:  noIcons,
		TextView: tview.NewTextView(),
	}

	p.SetWordWrap(true)
	p.SetWrap(true)
	p.SetDynamicColors(true)
	p.SetBorder(true)
	p.SetBorderPadding(0, 0, 1, 1)
	// styles.AddListener(&p)
	// p.SetInputCapture(p.)

	return &p
}
