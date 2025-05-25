package utils

import "github.com/rivo/tview"

type (
	InfoMsg struct {
	}
	ClearStatusMsg struct{}
)

func SetActive(box *tview.Box, title string, active bool) {
	// if active {
	// 	box.SetBorderColor(STYLE_BORDER_FOCUS.Fg)
	// 	box.SetTitleAlign(STYLE_TITLE_ACTIVE.Align)
	// 	title = dao.StyleFormat(title, STYLE_TITLE_ACTIVE.FormatStr)
	// 	if title != "" {
	// 		title = ColorizeTitle(title, *TUITheme.TitleActive)
	// 		box.SetTitle(title)
	// 	}
	// } else {
	// 	box.SetBorderColor(STYLE_BORDER.Fg)
	// 	box.SetTitleAlign(STYLE_TITLE.Align)
	// 	title = dao.StyleFormat(title, STYLE_TITLE.FormatStr)
	// 	if title != "" {
	// 		title = ColorizeTitle(title, *TUITheme.Title)
	// 		box.SetTitle(title)
	// 	}
	// }
}
