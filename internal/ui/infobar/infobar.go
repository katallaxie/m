package infobar

import (
	"fmt"

	"github.com/rivo/tview"
)

// InfoBar is a primitive for the app.
type InfoBar struct {
	*tview.TextView
}

// NewInfoBar returns a new InfoBar.
func NewInfoBar(appName string, appVersion string) *InfoBar {
	chat := &InfoBar{
		TextView: tview.NewTextView(),
	}

	chat.SetBorder(true)
	chat.SetText(fmt.Sprintf("%s (%s)", appName, appVersion))
	chat.SetTextAlign(tview.AlignCenter)
	chat.SetDynamicColors(true)

	return chat
}
