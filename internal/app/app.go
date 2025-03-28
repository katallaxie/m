package app

import (
	"os"

	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/entity"
	"github.com/katallaxie/m/internal/ui/chat"
	"github.com/katallaxie/m/internal/ui/help"
	"github.com/katallaxie/m/internal/ui/infobar"
	"github.com/katallaxie/m/internal/ui/utils"

	"github.com/epiclabs-io/winman"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// App is the main application.
type App struct {
	*tview.Application

	theme       *entity.Theme
	winMan      *winman.Manager
	currentPage string
	pages       *tview.Pages
	menu        *tview.TextView
	chat        *chat.Chat
	help        *help.Help
	infoBar     *infobar.InfoBar
	config      *config.Config
}

// New returns a new application.
func New(appName, version string, cfg *config.Config) *App {
	a := tview.NewApplication()
	wm := winman.NewWindowManager()

	app := &App{
		Application: a,
		winMan:      wm,
		theme:       &entity.TerminalTheme,
		pages:       tview.NewPages(),
		infoBar:     infobar.NewInfoBar("M", "0.1.0"),
		chat:        chat.NewChat("M", "0.1.0"),
		help:        help.NewHelp("M", "0.1.0"),
	}

	// menu items
	menuItems := [][]string{
		{utils.HelpScreenKey.Label(), app.help.GetTitle()},
		{utils.ChatScreenKey.Label(), app.chat.GetTitle()},
	}
	app.menu = newMenu(menuItems)

	app.pages.AddPage(app.help.GetTitle(), app.help, true, false)
	app.pages.AddPage(app.chat.GetTitle(), app.chat, true, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(app.infoBar, 3, 1, false).
		AddItem(app.pages, 0, 1, true).
		AddItem(app.menu, 1, 1, false)

	window := wm.NewWindow().
		Show().
		SetRoot(layout).
		SetBorder(false)

	app.SetRoot(window, true)

	// listen for user input
	app.Application.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == utils.AppExitKey.Key {
			app.Stop()
			os.Exit(0)
		}

		event = utils.ParseKeyEventKey(event)

		if !app.fontScreenHasActiveDialog() {
			// previous and next screen keys
			switch event.Rune() {
			case utils.NextScreenKey.Rune():
				return nil
			default:
			}

			// normal page key switch
			switch event.Key() { //nolint:exhaustive
			case utils.HelpScreenKey.EventKey():
				app.switchToScreen(app.help.GetTitle())

				return nil
			case utils.ChatScreenKey.EventKey():
				app.switchToScreen(app.chat.GetTitle())

				return nil
			}
		}

		return event
	})

	app.Init()

	return app
}

// Run runs the application.
func (a *App) Run() error {
	a.EnableMouse(true)
	return a.Application.Run()
}

// Init initializes the application.
func (a *App) Init() {
	a.SetFocus(a.chat)
}

func (a *App) switchToScreen(name string) {
	a.pages.SwitchToPage(name)
	a.setPageFocus(name)
	a.updatePageData(name)

	a.currentPage = name
}

func (a *App) setPageFocus(page string) {
	switch page {
	case a.help.GetTitle():
		a.Application.SetFocus(a.help)
	case a.chat.GetTitle():
		a.Application.SetFocus(a.chat)
	}
}

func (a *App) updatePageData(page string) {
	switch page {

	}
}

func (a *App) fontScreenHasActiveDialog() bool {
	switch a.currentPage {
	case a.help.GetTitle():
		return false
	case a.chat.GetTitle():
		return a.chat.HasFocus()
	}

	return false
}
