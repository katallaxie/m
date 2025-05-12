package infobar

import (
	"fmt"

	"github.com/katallaxie/m/internal/context"
	"github.com/katallaxie/m/internal/store"
	"github.com/katallaxie/m/internal/ui"
	"github.com/rivo/tview"
)

// InfoBar is a primitive for the app.
type InfoBar struct {
	*tview.TextView
	app ui.Application[store.State]
}

// NewInfoBar returns a new InfoBar.
func NewInfoBar(ctx *context.ProgramContext, app ui.Application[store.State]) *InfoBar {
	infoBar := &InfoBar{
		TextView: tview.NewTextView(),
		app:      app,
	}

	infoBar.SetBorder(true)
	infoBar.SetText(fmt.Sprintf("%s (%s)", ctx.GetAppName(), ctx.GetAppVersion()))
	infoBar.SetTextAlign(tview.AlignCenter)
	infoBar.SetDynamicColors(true)

	sub := app.GetStore().Subscribe()

	go func() {
		for change := range sub {
			app.QueueUpdateDraw(func() {
				infoBar.onUpdate(change.Curr())
			})
		}
	}()

	return infoBar
}

func (i *InfoBar) onUpdate(s store.State) {
}
