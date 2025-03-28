package help

import (
	"fmt"

	"github.com/katallaxie/m/internal/ui/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Chat is a chat primitive dialog.
type Chat struct {
	*tview.Box
	title  string
	layout *tview.Flex
}

// NewChat returns a chat screen primitive.
func NewChat(appName string, appVersion string) *Chat {
	chat := &Chat{
		Box:   tview.NewBox(),
		title: "Chat",
	}

	// colors
	headerColor := tcell.ColorWhite
	fgColor := tcell.ColorWhite
	bgColor := tcell.ColorDefault
	borderColor := tcell.ColorWhite

	// application keys description table
	keyinfo := tview.NewTable()
	// keyinfo.SetBackgroundColor(bgColor)
	keyinfo.SetFixed(1, 1)
	keyinfo.SetSelectable(false, false)

	// application description and version text view
	appinfo := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignLeft)
	appinfo.SetBackgroundColor(bgColor)

	appInfoText := fmt.Sprintf("%s\n\n%s %s\n\n%s")

	appinfo.SetText(appInfoText)
	appinfo.SetTextColor(headerColor)

	// help table items
	// the items will be divided into two separate tables
	rowIndex := 0
	colIndex := 0
	needInit := true
	maxRowIndex := len(utils.UIKeysBindings)/2 + 1 //nolint:mnd

	for i := range utils.UIKeysBindings {
		if i >= maxRowIndex {
			if needInit {
				colIndex = 2
				rowIndex = 0
				needInit = false
			}
		}

		keyinfo.SetCell(rowIndex, colIndex,
			tview.NewTableCell(fmt.Sprintf("%s:", utils.UIKeysBindings[i].KeyLabel)). //nolint:perfsprint
													SetAlign(tview.AlignRight).
													SetBackgroundColor(bgColor).
													SetSelectable(true).SetTextColor(headerColor))

		keyinfo.SetCell(rowIndex, colIndex+1,
			tview.NewTableCell(utils.UIKeysBindings[i].KeyDesc).
				SetAlign(tview.AlignLeft).
				SetBackgroundColor(bgColor).
				SetSelectable(true).SetTextColor(fgColor))

		rowIndex++
	}

	// appinfo and appkeys layout
	mlayout := tview.NewFlex().SetDirection(tview.FlexRow)
	mlayout.AddItem(appinfo, 2, 0, false) //nolint:mnd
	mlayout.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, false)
	mlayout.AddItem(keyinfo, 0, 1, false)
	mlayout.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, false)

	// layout
	// help.layout = tview.NewFlex().SetDirection(tview.FlexColumn)
	// help.layout.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, false)
	// help.layout.AddItem(mlayout, 0, 1, false)
	// help.layout.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, false)
	// help.layout.SetBorder(true)
	// help.layout.SetBackgroundColor(bgColor)
	chat.layout.SetBorderColor(borderColor)

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
