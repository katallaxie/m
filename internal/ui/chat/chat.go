package chat

import (
	"github.com/rivo/tview"
)

// Chat is a chat primitive dialog.
type Chat struct {
	*tview.TextView
	title string
}

// NewChat returns a chat screen primitive.
func NewChat(appName string, appVersion string) *Chat {
	chat := &Chat{
		TextView: tview.NewTextView(),
		title:    "💬 Chat",
	}

	chat.SetTitle(chat.title)
	chat.SetBorder(true)

	return chat
}
