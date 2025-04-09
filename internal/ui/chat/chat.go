package chat

import (
	"strings"

	"github.com/katallaxie/m/internal/store"
	"github.com/katallaxie/m/internal/ui"
	"github.com/rivo/tview"
)

// Chat is a chat primitive dialog.
type Chat struct {
	*tview.TextView
	title string
}

// NewChat returns a chat screen primitive.
func NewChat(app ui.Application[store.State], appName string, appVersion string) *Chat {
	chat := &Chat{
		TextView: tview.NewTextView(),
		title:    "💬 Chat",
	}

	chat.SetTitle(chat.title)
	chat.SetBorder(true)

	go func() {
		store := app.GetStore()

		for s := range store.Subscribe() {
			app.QueueUpdateDraw(func() {
				chat.SetText(strings.Join(s.Curr().Messages, ""))
				chat.ScrollToEnd()
			})
		}
	}()

	return chat
}
