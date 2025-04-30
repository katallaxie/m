package app

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/katallaxie/m/internal/api"
	"github.com/katallaxie/m/internal/cmd"
	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/entity"
	"github.com/katallaxie/m/internal/keymap"
	"github.com/katallaxie/m/internal/models"
	"github.com/katallaxie/m/internal/store"
	"github.com/katallaxie/m/internal/ui/activity"
	"github.com/katallaxie/m/internal/ui/chat"
	"github.com/katallaxie/m/internal/ui/help"
	"github.com/katallaxie/m/internal/ui/history"
	"github.com/katallaxie/m/internal/ui/infobar"
	"github.com/katallaxie/m/internal/ui/modals"
	"github.com/katallaxie/m/internal/ui/utils"

	"github.com/epiclabs-io/winman"
	"github.com/katallaxie/pkg/redux"
	"github.com/rivo/tview"
)

// App is the main application.
type App struct {
	*tview.Application

	api        *api.Api
	chat       *chat.Chat
	config     *config.Config
	ctx        context.Context
	history    *history.History
	activities *activity.Activity
	infoBar    *infobar.InfoBar
	menu       *tview.TextView
	pages      *tview.Pages
	prompt     *chat.Prompt
	state      redux.Store[store.State]
	theme      *entity.Theme
	winMan     *winman.Manager
}

// New returns a new application.
func New(ctx context.Context, appName, version string, cfg *config.Config) *App {
	a := tview.NewApplication()
	wm := winman.NewWindowManager()

	client := api.ClientFactory(cfg.Spec.Api.Provider, cfg.Spec.Api.Model, cfg.Spec.Api.URL, cfg.Spec.Api.Key)

	app := &App{
		Application: a,
		ctx:         ctx,
		winMan:      wm,
		theme:       &entity.TerminalTheme,
		pages:       tview.NewPages(),
		api:         api.NewApi(client),
		config:      cfg,
	}

	state := store.NewState()
	state.History.Next()

	// State machine
	app.state = redux.New(
		ctx,
		state,
		store.ChatMessageReducer,
		// store.UpdateMessageReducer,
		// store.SetStatusReducer,
		// store.AddNotebookReducer,
	)

	// Chat panel
	app.chat = chat.NewChat(app, "M", "0.1.0")

	// Prompt panel
	app.prompt = chat.NewPrompt(app, app.api)

	// History panel
	app.history = history.NewHistory(app)

	// Activity panel
	app.activities = activity.NewActivity(app)

	// Info bar
	app.infoBar = infobar.NewInfoBar(appName, version)

	// menu items
	menuItems := [][]string{
		{utils.HelpScreenKey.Label(), "Help"},
		{utils.NewChat.Label(), "New"},
		{utils.AppExitKey.Label(), "Quit"},
	}
	app.menu = newMenu(menuItems)

	sidebarPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(app.history, 0, 1, false).
		AddItem(app.activities, 3, 0, false)

	mainPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(app.chat, 0, 3, false).
		AddItem(app.prompt, 0, 1, true)

	mainLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(sidebarPanel, 35, 1, false).
		AddItem(mainPanel, 0, 4, true)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(app.infoBar, 3, 1, false).
		AddItem(mainLayout, 0, 1, true).
		AddItem(app.menu, 1, 1, false)

	app.pages.AddPage("Main", layout, true, true)
	app.pages.AddPage("Help", help.NewHelpModal(app), false, false)
	app.pages.AddPage("Quit", modals.NewQuitModal(app), false, false)

	window := wm.NewWindow().
		Show().
		SetRoot(app.pages).
		SetBorder(false)

	app.EnableMouse(true)
	app.EnablePaste(false)

	app.SetRoot(window, true)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		command := keymap.Keymaps.Group(keymap.HomeGroup).Resolve(event)

		if command == cmd.HelpPopup {
			app.pages.ShowPage("Help")
		}

		if command == cmd.Quit {
			app.pages.ShowPage("Quit")
		}

		if command == cmd.FocusPrompt {
			app.SetFocus(app.prompt)
		}

		if command == cmd.NewChat {
			app.GetStore().Dispatch(store.NewAddChat(models.NewChat()))
		}

		return event
	})

	app.SetFocus(app.chat)

	return app
}

// Context returns the context of the application.
func (a *App) Context() context.Context {
	return a.ctx
}

// Pages returns the pages of the application.
func (a *App) Pages() *tview.Pages {
	return a.pages
}

// StateUpdates returns the state updates.
func (a *App) GetState() store.State {
	return a.state.State()
}

// GetStore returns the state store.
func (a *App) GetStore() redux.Store[store.State] {
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

// Config returns the configuration of the application.
func (a *App) Config() *config.Config {
	return a.config
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
