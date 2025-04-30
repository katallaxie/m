package chat

import (
	"fmt"

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
		title:    " 💬 Chat ",
	}

	chat.SetTitle(chat.title)
	chat.SetBorder(true)
	chat.SetDynamicColors(true)
	chat.SetWrap(true)
	chat.SetScrollable(true)

	go func() {
		store := app.GetStore()

		for range store.Subscribe() {
			app.QueueUpdateDraw(func() {
				w := chat.BatchWriter()
				defer w.Close()
				w.Clear()

				curr := store.State()
				msgs := curr.Notebooks[curr.CurrentNotebook].Messages

				for _, msg := range msgs {
					fmt.Fprintln(w, "[red::] 👨‍💻 You:[-]")
					fmt.Fprintln(w, msg.Content())
				}

				chat.ScrollToEnd()
			})
		}
	}()

	return chat
}
