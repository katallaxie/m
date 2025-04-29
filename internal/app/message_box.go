package app

import (
	"fmt"

	"github.com/epiclabs-io/winman"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Button struct {
	Name    string
	OnClick func(msgboxWnd *winman.WindowBase)
}

type ShowMessageBoxParam struct {
	title   string
	message string
	buttons []Button
}

// ShowMessageBox is used to show message box with fixed size 50x11
func (a *App) ShowMessageBox(param ShowMessageBoxParam) {
	txtMessage := tview.NewTextView()

	form := tview.NewForm()
	fmt.Fprint(txtMessage, param.message)
	form.AddFormItem(txtMessage)
	form.SetButtonsAlign(tview.AlignCenter)
	// form.SetButtonStyle(u.Theme.Style.ButtonStyle)
	// form.SetBackgroundColor(u.Theme.Colors.WindowColor)

	wnd := a.CreateModalDialog(CreateModalDialogParam{
		title:      param.title,
		rootView:   form,
		draggable:  true,
		resizeable: false,
		size:       winSize{0, 0, 50, 11},
		// fallbackFocus: a.Layout.MenuList,
	})

	a.ShowMessageBox_SetComponentActions(wnd, form, param.buttons)
	a.ShowMessageBox_SetInputCapture(wnd, form)
}

func (a *App) ShowMessageBox_SetComponentActions(wnd *winman.WindowBase, form *tview.Form, buttons []Button) {
	// Populate buttons
	for _, button := range buttons {
		form.AddButton(button.Name, func() { button.OnClick(wnd) })
	}

	if (len(buttons)) > 0 {
		// Set focus to the first button
		form.SetFocus(1)
	}
}

func (a *App) ShowMessageBox_SetInputCapture(wnd *winman.WindowBase, parent *tview.Form) {
	parent.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			a.CloseModalDialog(wnd, a.chat)
		}
		return event
	})
}
