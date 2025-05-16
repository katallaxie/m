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
	infoBar.SetText(fmt.Sprintf("%s (%s) - %s", infoBar.ctx.GetAppName(), infoBar.ctx.GetAppVersion(), infoBar.app.GetState().History.Active().Name))

	go infoBar.onUpdate()

	return infoBar
}

func (i *InfoBar) onUpdate() {
	for {
		select {
		case <-i.ctx.Context().Done():
			return
		case change := <-i.app.GetStore().Subscribe():
			i.app.QueueUpdateDraw(func() {
				i.SetText(fmt.Sprintf("%s (%s) - %s", i.ctx.GetAppName(), i.ctx.GetAppVersion(), change.Curr().History.Active().Name))
			})
		}
	}
}
