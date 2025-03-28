package chat

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Chat is a chat primitive dialog.
type Chat struct {
	*tview.Box
	title         string
	layout        *tview.Flex
	prompt        *Prompt
	notebooksList *NotebookList
}

// NewChat returns a chat screen primitive.
func NewChat(appName string, appVersion string) *Chat {
	chat := &Chat{
		Box:           tview.NewBox(),
		notebooksList: NewNotebookList(),
		prompt:        NewPrompt(),
		title:         "Chat",
	}

	splitSidebar := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(chat.notebooksList, 15, 1, false)

	splitMainPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(chat.prompt, 15, 1, true)

	playout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(splitSidebar, 35, 1, false).
		AddItem(splitMainPanel, 0, 4, true)

	// layout
	chat.layout = playout

	return chat
}

// GetTitle returns primitive title.
func (c *Chat) GetTitle() string {
	return c.title
}

// HasFocus returns whether or not this primitive has focus.
func (c *Chat) HasFocus() bool {
	return c.Box.HasFocus() || c.layout.HasFocus()
}

// Focus is called when this primitive receives focus.
func (c *Chat) Focus(delegate func(p tview.Primitive)) {
	delegate(c.layout)
}

// Draw draws this primitive onto the screen.
func (c *Chat) Draw(screen tcell.Screen) {
	x, y, width, height := c.Box.GetInnerRect()
	if height <= 3 { //nolint:mnd
		return
	}

	c.Box.DrawForSubclass(screen, c)
	c.layout.SetRect(x, y, width, height)
	c.layout.Draw(screen)
}
