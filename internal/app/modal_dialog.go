package app

import (
	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type winSize struct {
	x      int
	y      int
	width  int
	height int
}

type CreateModalDialogParam struct {
	title         string
	rootView      tview.Primitive
	draggable     bool
	resizeable    bool
	size          winSize
	fallbackFocus tview.Primitive
}

// CreateModalDialogParam is a helper to create a modal dialog window
func (a *App) CreateModalDialog(param CreateModalDialogParam) *winman.WindowBase {
	wnd := winman.NewWindow().Show()

	wnd.SetTitle(param.title)
	wnd.SetRoot(param.rootView)
	wnd.SetDraggable(param.draggable)
	wnd.SetResizable(param.resizeable)
	wnd.SetModal(true)
	// wnd.SetBackgroundColor(a.Theme.Colors.WindowColor)

	wnd.SetRect(param.size.x, param.size.y, param.size.width, param.size.height)
	wnd.AddButton(&winman.Button{
		Symbol: 'X',
		OnClick: func() {
			// Close current window and get back focus to the fallback primitive
			a.CloseModalDialog(wnd, param.fallbackFocus)
		},
	})

	a.winMan.AddWindow(wnd)
	a.winMan.Center(wnd)
	a.SetFocus(wnd)

	return wnd
}

// CloseModalDialog is a helper to close the modal dialog window
func (a *App) CloseModalDialog(wnd *winman.WindowBase, focus tview.Primitive) {
	a.winMan.RemoveWindow(wnd)
	a.SetFocus(focus)
}
