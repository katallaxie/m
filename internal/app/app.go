package app

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/katallaxie/m/internal/api"
	"github.com/katallaxie/m/internal/cmd"
	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/entity"
	"github.com/katallaxie/m/internal/keymap"
	"github.com/katallaxie/m/internal/model"
	"github.com/katallaxie/m/internal/store"
	"github.com/katallaxie/m/internal/ui/activity"
	"github.com/katallaxie/m/internal/ui/chat"
	"github.com/katallaxie/m/internal/ui/help"
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

	theme   *entity.Theme
	winMan  *winman.Manager
	pages   *tview.Pages
	menu    *tview.TextView
	chat    *chat.Chat
	prompt  *chat.Prompt
	infoBar *infobar.InfoBar
	config  *config.Config
	state   redux.Store[store.State]
	api     *api.Api
	ctx     context.Context
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
		infoBar:     infobar.NewInfoBar("M", "0.1.0"),
		api:         api.NewApi(client),
		config:      cfg,
	}

	// State machine
	app.state = redux.New(
		store.NewState(),
		store.AddMessageReducer,
		store.UpdateMessageReducer,
		store.SetStatusReducer,
		store.AddNotebookReducer,
	)

	// Chat panel
	app.chat = chat.NewChat(app, "M", "0.1.0")

	// Prompt panel
	app.prompt = chat.NewPrompt(app, app.api)

	// menu items
	menuItems := [][]string{
		{utils.HelpScreenKey.Label(), "Help"},
		{utils.NewNotebookKey.Label(), "New Notebook"},
		{utils.AppExitKey.Label(), "Quit"},
	}
	app.menu = newMenu(menuItems)

	sidebarPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(chat.NewNotebookList(app), 0, 1, true).
		AddItem(activity.NewActivity(app), 3, 0, false)

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

	// app.SetRoot(window, true)
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

		if command == cmd.NewNotebook {
			app.GetStore().Dispatch(store.NewAddNotebook(model.NewNotebook()))
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
