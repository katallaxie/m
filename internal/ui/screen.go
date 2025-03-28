package ui

func (ui *UI) switchToScreen(name string) {
	ui.pages.SwitchToPage(name)
	ui.setPageFocus(name)
	ui.updatePageData(name)

	ui.currentPage = name
}

func (ui *UI) setPageFocus(page string) {
	switch page {
	case ui.help.GetTitle():
		ui.App.SetFocus(ui.help)
	}
}

func (ui *UI) updatePageData(page string) {
	switch page {

	}
}
