package infobar

import (
	"fmt"

	"github.com/rivo/tview"
)

const (
	InfoBarViewHeight = 5
)

// InfoBar is a primitive for the app.
type InfoBar struct {
	*tview.TextView
	title string
}

// NewInfoBar returns a new InfoBar.
func NewInfoBar(appName string, appVersion string) *InfoBar {
	chat := &InfoBar{
		TextView: tview.NewTextView(),
		title:    "Chat",
	}

	chat.SetBorder(true)
	chat.SetText(fmt.Sprint(" 🤖 M"))
	chat.SetTextAlign(tview.AlignCenter)

	return chat
}
