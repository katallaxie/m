package app

import (
	"github.com/katallaxie/m/internal/api"
	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/entity"
	"github.com/katallaxie/m/internal/state"
	"github.com/katallaxie/m/internal/ui/activity"
	"github.com/katallaxie/m/internal/ui/chat"
	"github.com/katallaxie/m/internal/ui/help"
	"github.com/katallaxie/m/internal/ui/infobar"
	"github.com/katallaxie/m/internal/ui/utils"

	"github.com/epiclabs-io/winman"
	"github.com/katallaxie/pkg/fsmx"
	"github.com/katallaxie/prompts/ollama"
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
	state       fsmx.Store[state.State]
	api         *api.Api
}

// New returns a new application.
func New(appName, version string, cfg *config.Config) *App {
	a := tview.NewApplication()
	wm := winman.NewWindowManager()

	client, err := ollama.New(ollama.WithBaseURL("http://localhost:7869"), ollama.WithModel("smollm"))
	if err != nil {
		panic(err)
	}

	app := &App{
		Application: a,
		winMan:      wm,
		theme:       &entity.TerminalTheme,
		pages:       tview.NewPages(),
		infoBar:     infobar.NewInfoBar("M", "0.1.0"),
		help:        help.NewHelp("M", "0.1.0"),
		api:         api.NewApi(client),
	}

	// State machine
	app.state = fsmx.New(state.NewState(), state.AddMessageReducer, state.SetStatusReducer)

	// Chat panel
	app.chat = chat.NewChat(app, "M", "0.1.0")

	// menu items
	menuItems := [][]string{
		{utils.HelpScreenKey.Label(), app.help.GetTitle()},
		{utils.ChatScreenKey.Label(), app.chat.GetTitle()},
	}
	app.menu = newMenu(menuItems)

	modal := tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})

	sidebarPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(chat.NewNotebookList(), 0, 1, true).
		AddItem(activity.NewActivity(app), 3, 0, false)

	mainPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(app.chat, 0, 3, false).
		AddItem(chat.NewPrompt(app, app.api), 0, 1, true)

	mainLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(sidebarPanel, 35, 1, false).
		AddItem(mainPanel, 0, 4, true)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(app.infoBar, 3, 1, false).
		AddItem(mainLayout, 0, 1, true).
		AddItem(app.menu, 1, 1, false)

	app.pages.AddPage("Quit", modal, true, false)
	app.pages.AddPage("Main", mainLayout, false, true)

	window := wm.NewWindow().
		Show().
		SetRoot(layout).
		SetBorder(false)

	app.SetRoot(window, true)
	app.EnableMouse(true)
	app.EnablePaste(false)

	return app
}

// StateUpdates returns the state updates.
func (a *App) GetState() state.State {
	return a.state.State()
}

// GetStore returns the state store.
func (a *App) GetStore() fsmx.Store[state.State] {
	return a.state
}

// Stop stops the application.
func (a *App) Stop() {
	a.Application.Stop()
}

// Draw draws the application.
func (a *App) Draw() {
	a.Application.Draw()
}

// Run runs the application.
func (a *App) Run() error {
	// a.Application.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	if event.Key() == utils.AppExitKey.Key {
	// 		a.Stop()
	// 		os.Exit(0)
	// 	}

	// 	event = utils.ParseKeyEventKey(event)

	// 	if !a.fontScreenHasActiveDialog() {
	// 		// previous and next screen keys
	// 		switch event.Rune() {
	// 		case utils.NextScreenKey.Rune():
	// 			return nil
	// 		default:
	// 		}

	// 		// normal page key switch
	// 		switch event.Key() { //nolint:exhaustive
	// 		case utils.HelpScreenKey.EventKey():
	// 			a.switchToScreen(a.help.GetTitle())

	// 			return nil
	// 		case utils.ChatScreenKey.EventKey():
	// 			a.switchToScreen(a.chat.GetTitle())

	// 			return nil
	// 		}
	// 	}

	// 	return event
	// })

	a.Init()

	return a.Application.Run()
}

// Init initializes the application.
func (a *App) Init() {
}

// QueueUpdateDraw queues up a ui action and redraw the ui.
func (a *App) QueueUpdateDraw(f func()) {
	go func() {
		a.Application.QueueUpdateDraw(f)
	}()
}
