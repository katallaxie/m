package chat

import (
	"fmt"
	"strings"
	"time"

	"github.com/katallaxie/m/internal/state"
	"github.com/katallaxie/m/internal/ui"
	"github.com/rivo/tview"
)

// Chat is a chat primitive dialog.
type Chat struct {
	*tview.TextView
	title string
}

// NewChat returns a chat screen primitive.
func NewChat(app ui.Application, appName string, appVersion string) *Chat {
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
			app.GetState().Dispatch(state.NewAddMessage(newBuffer))
		}
	}()

	go func() {
		for s := range app.GetState().Subscribe() {
			s := s.(*state.State)
			app.QueueUpdateDraw(func() {
				chat.SetText(strings.Join(s.Messages, ""))
				chat.ScrollToEnd()
			})
		}
	}()

	return chat
}
