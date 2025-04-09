package help

import (
	"github.com/gdamore/tcell/v2"
	"github.com/katallaxie/m/internal/cmd"
	"github.com/katallaxie/m/internal/keymap"
	"github.com/katallaxie/m/internal/ui"
	"github.com/katallaxie/pkg/redux"
	"github.com/rivo/tview"
)

type HelpModal[S redux.State] struct {
	app ui.Application[S]
	tview.Primitive
}

// NewHelpModal creates a new help modal.
func NewHelpModal[S redux.State](a ui.Application[S]) *HelpModal[S] {
	// Returns a new primitive which puts the provided primitive in the center and
	// sets its size to the given width and height.
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	table := tview.NewTable()

	// table.SetBorders(true)
	table.SetBorder(true)
	table.SetTitle(" Keybindings ")
	table.SetSelectable(true, false)

	table.Select(3, 0)

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		command := keymap.Keymaps.Group(keymap.HelpGroup).Resolve(event)

		if command == cmd.Close {
			a.Pages().HidePage("Help")
		}

		return event
	})

	r := &HelpModal[S]{a, modal(table, 0, 30)}

	return r
}
