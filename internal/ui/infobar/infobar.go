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
	ctx *context.ProgramContext
}

// NewInfoBar returns a new InfoBar.
func NewInfoBar(ctx *context.ProgramContext, app ui.Application[store.State]) *InfoBar {
	infoBar := &InfoBar{
		TextView: tview.NewTextView(),
		app:      app,
		ctx:      ctx,
	}

	infoBar.SetBorder(true)
	infoBar.SetTextAlign(tview.AlignCenter)
	infoBar.SetDynamicColors(true)

	sub := app.GetStore().Subscribe()
	infoBar.onUpdate(app.GetStore().State())

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
	i.SetText(fmt.Sprintf("%s (%s) - %s", i.ctx.GetAppName(), i.ctx.GetAppVersion(), s.History.Active().Name))
}
