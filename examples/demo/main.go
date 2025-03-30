// Demo code for the Form primitive.
package main

import (
	"fmt"
	"strconv"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

func calculator() *winman.WindowBase {

	value := []float64{0.0, 0.0}
	i := 0
	decimal := 1.0
	op := ' '
	display := tview.NewTextView().
		SetText("0.").
		SetTextAlign(tview.AlignRight)

	keyPressed := func(char rune) {
		if char >= '0' && char <= '9' {
			digit := (float64)(char - '0')
			if decimal == 1.0 {
				value[i] = value[i]*10 + digit
			} else {
				value[i] = value[i] + digit*decimal
				decimal /= 10
			}
		} else {
			switch char {
			case '.':
				if decimal == 1.0 {
					decimal = decimal / 10
				}
			case '=':
				if i == 1 {
					switch op {
					case '+':
						value[0] = value[0] + value[1]
					case '-':
						value[0] = value[0] - value[1]
					case 'x':
						value[0] = value[0] * value[1]
					case '/':
						if value[1] == 0.0 {
							display.SetText("Err")
							value[0] = 0.0
						} else {
							value[0] = value[0] / value[1]
						}
					}
					i = 0
					decimal = 1.0
				} else {
					value[0] = 0.0
				}
				op = ' '
			default:
				op = char
				i = 1
				decimal = 1.0
				value[1] = 0
			}
		}
		display.SetText(fmt.Sprintf("%g", value[i]))
	}

	newCalcButton := func(char rune) *tview.Button {
		return tview.NewButton(string(char)).SetSelectedFunc(func() {
			keyPressed(char)
		})
	}

	grid := tview.NewGrid().
		SetRows(2, 0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		SetBorders(true).
		AddItem(display, 0, 0, 1, 4, 2, 0, false)

	buttons := []rune{'7', '8', '9', '/', '4', '5', '6', 'x', '1', '2', '3', '-', '0', '.', '=', '+'}

	for i, b := range buttons {
		row := 1 + i/4
		col := i % 4
		grid.AddItem(newCalcButton(b), row, col, 1, 1, 1, 1, true)
	}

	wnd := winman.NewWindow().SetRoot(grid)
	wnd.AddButton(&winman.Button{
		Symbol:    'X',
		Alignment: winman.ButtonLeft,
		OnClick:   func() { wnd.Hide() },
	})
	wnd.SetRect(0, 0, 30, 15)
	wnd.Draggable = true
	wnd.Resizable = true

	return wnd
}

// MsgBox creates a new modal message box
func MsgBox(title, text string, buttons []string, callback func(clicked string)) *winman.WindowBase {

	msgBox := winman.NewWindow()
	message := tview.NewTextView().SetText(text).SetTextAlign(tview.AlignCenter)
	buttonBar := tview.NewFlex().
		SetDirection(tview.FlexColumn)

	content := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(message, 0, 1, false).
		AddItem(buttonBar, 1, 0, true)

	msgBox.SetRoot(content)
	msgBox.SetTitle(title).
		SetRect(4, 2, 30, 6)
	msgBox.Draggable = true
	msgBox.Modal = true

	for _, buttonText := range buttons {
		button := func(buttonText string) *tview.Button {
			return tview.NewButton(buttonText).SetSelectedFunc(func() {
				msgBox.Hide()
				callback(buttonText)
			})
		}(buttonText)
		buttonBar.AddItem(button, 0, 1, true)
	}

	return msgBox
}

func main() {

	app := tview.NewApplication()
	wm := winman.NewWindowManager()

	quitMsgBox := MsgBox("Confirmation", "Really quit?", []string{"Yes", "No"}, func(clicked string) {
		if clicked == "Yes" {
			app.Stop()
		}
	})
	wm.AddWindow(quitMsgBox)

	calc := calculator()
	wm.AddWindow(calc)

	setFocus := func(p tview.Primitive) {
		go app.QueueUpdateDraw(func() {
			app.SetFocus(p)
		})
	}

	var createForm func(modal bool) *winman.WindowBase
	var counter = 0

	setZ := func(wnd *winman.WindowBase, newZ int) {
		go app.QueueUpdateDraw(func() {
			newTopWindow := wm.Window(wm.WindowCount() - 2)
			if newTopWindow != nil {
				app.SetFocus(newTopWindow)
				wm.SetZ(wnd, newZ)
			}
		})
	}

	createForm = func(modal bool) *winman.WindowBase {
		counter++
		form := tview.NewForm()
		window := winman.NewWindow().
			SetRoot(form).
			SetResizable(true).
			SetDraggable(true).
			SetModal(modal)

		quit := func() {
			if wm.WindowCount() == 3 {
				quitMsgBox.Show()
				wm.Center(quitMsgBox)
				setFocus(quitMsgBox)
			} else {
				wm.RemoveWindow(window)
				setFocus(wm)
			}
		}

		form.AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
			AddInputField("First name", "", 20, nil, nil).
			AddPasswordField("Password", "", 10, '*', nil).
			AddCheckbox("Draggable", window.IsDraggable(), func(checked bool) {
				window.SetDraggable(checked)
			}).
			AddCheckbox("Resizable", window.IsResizable(), func(checked bool) {
				window.SetResizable(checked)
			}).
			AddCheckbox("Modal", window.Modal, func(checked bool) {
				window.SetModal(checked)
			}).
			AddCheckbox("Border", window.Draggable, func(checked bool) {
				window.SetBorder(checked)
			}).
			AddInputField("Z-Index", "", 20, func(text string, char rune) bool {
				return char >= '0' && char <= '9'
			}, nil).
			AddButton("Set Z", func() {
				zIndexField := form.GetFormItemByLabel("Z-Index").(*tview.InputField)
				z, _ := strconv.Atoi(zIndexField.GetText())
				setZ(window, z)
			}).
			AddButton("New", func() {
				newWnd := createForm(false).Show()
				wm.AddWindow(newWnd)
				setFocus(newWnd)
			}).
			AddButton("Modal", func() {
				newWnd := createForm(true).Show()
				newWnd.Modal = true
				wm.AddWindow(newWnd)
				setFocus(newWnd)
			}).
			AddButton("Calc", func() {
				calc.Show()
				wm.Center(calc)
				setFocus(calc)
			}).
			AddButton("Close", quit)

		title := fmt.Sprintf("Window%d", counter)
		window.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)
		window.SetRect(2+counter*2, 2+counter, 50, 30)
		window.AddButton(&winman.Button{
			Symbol:    'X',
			Alignment: winman.ButtonLeft,
			OnClick:   quit,
		})

		var maxMinButton *winman.Button
		maxMinButton = &winman.Button{
			Symbol:    '▴',
			Alignment: winman.ButtonRight,
			OnClick: func() {
				if window.IsMaximized() {
					window.Restore()
					maxMinButton.Symbol = '▴'
				} else {
					window.Maximize()
					maxMinButton.Symbol = '▾'
				}
			},
		}
		window.AddButton(maxMinButton)
		wm.AddWindow(window)
		return window
	}

	for i := 0; i < 1; i++ {
		createForm(false).Show()
	}

	if err := app.SetRoot(wm, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
