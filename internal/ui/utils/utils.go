package utils

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// EmptyBoxSpace returns simple Box without border with bgColor as background.
func EmptyBoxSpace(bgColor tcell.Color) *tview.Box {
	box := tview.NewBox()
	box.SetBackgroundColor(bgColor)
	box.SetBorder(false)

	return box
}
