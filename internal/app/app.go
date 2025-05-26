package app

import (
	"context"

	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/ports"
	"github.com/katallaxie/pkg/dbx"
	"github.com/katallaxie/prompts"
)

// App is the main application.
type App struct {
	// *tview.Application

	client prompts.Chat
	store  dbx.Database[ports.ReadTx, ports.ReadWriteTx]

	// api     *api.Api
	// chat    *chat.Chat
	// config  *config.Config
	// ctx     *context.ProgramContext
	// history *history.History
	// infoBar *infobar.InfoBar
	// menu    *menu.Menu
	// pages   *tview.Pages
	// prompt  *prompt.Prompt
	// state   redux.Store[store.State]
	// theme   *entity.Theme
	// winMan  *winman.Manager
}

// New returns a new application.
func New(ctx context.Context, store dbx.Database[ports.ReadTx, ports.ReadWriteTx], client prompts.Chat, cfg config.Config) (*App, error) {
	app := new(App)
	app.store = store
	app.client = client

	err := app.Init()
	if err != nil {
		return nil, err
	}

	return app, nil

	// a := tview.NewApplication()
	// wm := winman.NewWindowManager()

	// client := api.ClientFactory(cfg.Spec.Api.Provider, cfg.Spec.Api.Model, cfg.Spec.Api.URL, cfg.Spec.Api.Key)
	// api := api.NewApi(client)

	// app := &App{
	// 	api:         api,
	// 	Application: a,
	// 	config:      cfg,
	// 	ctx:         ctx,
	// 	pages:       tview.NewPages(),
	// 	theme:       &entity.TerminalTheme,
	// 	winMan:      wm,
	// }

	// state := store.NewState()
	// state.History.Next()

	// // State machine
	// app.state = redux.New(
	// 	ctx.Context(),
	// 	state,
	// 	store.ChatMessageReducer,
	// 	// store.UpdateMessageReducer,
	// 	// store.SetStatusReducer,
	// 	// store.AddNotebookReducer,
	// )

	// // Chat panel
	// app.chat = chat.NewChat(app, "M", "0.1.0")

	// // Prompt panel
	// app.prompt = prompt.NewPrompt(app, app.api)

	// // History panel
	// app.history = history.NewHistory(app)

	// // Activity panel
	// // app.activities = activity.NewActivity(app)

	// // Info bar
	// app.infoBar = infobar.NewInfoBar(ctx, app)

	// // menu items
	// menuItems := [][]string{
	// 	{utils.HelpScreenKey.Label(), "Help"},
	// 	{utils.NewChat.Label(), "New"},
	// 	{utils.AppExitKey.Label(), "Quit"},
	// }
	// app.menu = menu.NewMenu(cfg.AppName, cfg.Version, menuItems)

	// sidebarPanel := tview.NewFlex().SetDirection(tview.FlexRow).
	// 	AddItem(app.history, 0, 1, false)

	// mainPanel := tview.NewFlex().SetDirection(tview.FlexRow).
	// 	AddItem(app.chat, 0, 3, false).
	// 	AddItem(app.prompt, 0, 1, true)

	// mainLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
	// 	AddItem(sidebarPanel, 35, 1, false).
	// 	AddItem(mainPanel, 0, 4, true)

	// layout := tview.NewFlex().
	// 	SetDirection(tview.FlexRow).
	// 	AddItem(app.infoBar, 3, 1, false).
	// 	AddItem(mainLayout, 0, 1, true).
	// 	AddItem(app.menu, 1, 1, false)

	// app.pages.AddPage("Main", layout, true, true)
	// app.pages.AddPage("Help", help.NewHelpModal(app), false, false)
	// app.pages.AddPage("Quit", modals.NewQuitModal(app), false, false)

	// window := wm.NewWindow().
	// 	Show().
	// 	SetRoot(app.pages).
	// 	SetBorder(false)

	// app.EnableMouse(true)
	// app.EnablePaste(false)

	// app.SetRoot(window, true)
	// app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	command := keymap.Keymaps.Group(keymap.HomeGroup).Resolve(event)

	// 	if command == cmd.HelpPopup {
	// 		app.pages.ShowPage("Help")
	// 	}

	// 	if command == cmd.Quit {
	// 		app.pages.ShowPage("Quit")
	// 	}

	// 	if command == cmd.FocusPrompt {
	// 		app.SetFocus(app.prompt)
	// 	}

	// 	if command == cmd.NewChat {
	// 		app.GetStore().Dispatch(store.NewAddChat(models.NewChat()))
	// 	}

	// 	return event
	// })

	// app.SetFocus(app.chat)

	// return app
}

// Init initializes the application.
func (a *App) Init() error {
	return nil
}

// Dispose cleans up the application resources.
func (a *App) Dispose() {
}

// // Context returns the context of the application.
// func (a *App) Context() *context.ProgramContext {
// 	return a.ctx
// }

// // Pages returns the pages of the application.
// func (a *App) Pages() *tview.Pages {
// 	return a.pages
// }

// // StateUpdates returns the state updates.
// func (a *App) GetState() store.State {
// 	return a.state.State()
// }

// // GetStore returns the state store.
// func (a *App) GetStore() redux.Store[store.State] {
// 	return a.state
// }

// // Stop stops the application.
// func (a *App) Stop() {
// 	a.Application.Stop()
// }

// // Draw draws the application.
// func (a *App) Draw() {
// 	a.Application.Draw()
// }

// // Config returns the configuration of the application.
// func (a *App) Config() *config.Config {
// 	return a.config
// }

// // Run runs the application.
// func (a *App) Run() error {
// 	// a.Application.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
// 	// 	if event.Key() == utils.AppExitKey.Key {
// 	// 		a.Stop()
// 	// 		os.Exit(0)
// 	// 	}

// 	// 	event = utils.ParseKeyEventKey(event)

// 	// 	if !a.fontScreenHasActiveDialog() {
// 	// 		// previous and next screen keys
// 	// 		switch event.Rune() {
// 	// 		case utils.NextScreenKey.Rune():
// 	// 			return nil
// 	// 		default:
// 	// 		}

// 	// 		// normal page key switch
// 	// 		switch event.Key() { //nolint:exhaustive
// 	// 		case utils.HelpScreenKey.EventKey():
// 	// 			a.switchToScreen(a.help.GetTitle())

// 	// 			return nil
// 	// 		case utils.ChatScreenKey.EventKey():
// 	// 			a.switchToScreen(a.chat.GetTitle())

// 	// 			return nil
// 	// 		}
// 	// 	}

// 	// 	return event
// 	// })

// 	a.Init()

// 	return a.Application.Run()
// }

// // Init initializes the application.
// func (a *App) Init() error {
// 	return nil
// }

// // QueueUpdate queues up a ui action.
// func (a *App) QueueUpdate(f func()) {
// 	go func() {
// 		a.Application.QueueUpdate(f)
// 	}()
// }

// // QueueUpdateDraw queues up a ui action and redraw the ui.
// func (a *App) QueueUpdateDraw(f func()) {
// 	go func() {
// 		a.Application.QueueUpdateDraw(f)
// 	}()
// }
