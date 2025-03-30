package chat

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

// Chat is a chat primitive dialog.
type Chat struct {
	*tview.TextView
	title string
}

// NewChat returns a chat screen primitive.
func NewChat(app *tview.Application, appName string, appVersion string) *Chat {
	chat := &Chat{
		TextView: tview.NewTextView(),
		title:    "💬 Chat",
	}

	chat.SetTitle(chat.title)
	chat.SetBorder(true)

	go func() {
		var (
			newBuffer string
		)

		ticker := time.NewTicker(1 * time.Second)

		for _ = range ticker.C {
			timeHeader := time.Now().Format("15:04:05 02/01/2006")

			newBuffer += fmt.Sprintln(timeHeader)
			app.QueueUpdateDraw(func() {
				chat.SetText(newBuffer)
				chat.ScrollToEnd()
			})
		}
	}()

	return chat
}
