package modals

import (
	"github.com/katallaxie/m/internal/ui"
	"github.com/katallaxie/pkg/redux"
	"github.com/rivo/tview"
)

// QuitModal is a modal dialog that asks the user if they want to quit the application.
type QuitModal[S redux.State] struct {
	*tview.Modal
	app ui.Application[S]
}

// NewQuitModal creates a new quit modal.
func NewQuitModal[S redux.State](a ui.Application[S]) *QuitModal[S] {
	modal := &QuitModal[S]{
		Modal: tview.NewModal(),
		app:   a,
	}

	modal.SetBorder(true)
	modal.SetText("Do you want to quit the application?")
	modal.AddButtons([]string{"Yes", "No"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Yes" {
			a.Stop()
		} else {
			a.Pages().HidePage("Quit")
		}
	})

	return modal
}
