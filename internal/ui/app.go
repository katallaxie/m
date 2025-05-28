package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/katallaxie/m/internal/app"
	"github.com/katallaxie/m/internal/ui/dialogs"
	"github.com/katallaxie/m/internal/ui/layout"
	"github.com/katallaxie/m/internal/ui/pages"
	"github.com/katallaxie/m/internal/ui/status"
	"github.com/katallaxie/pkg/slices"
)

type application struct {
	app *app.App

	pages        map[pages.ID]tea.Model
	currentPage  pages.ID
	previousPage pages.ID

	showThemeDialog bool
	showHelpDialog  bool
	themeDialog     tea.Model
	helpDialog      dialogs.Help
	status          status.Status

	width, height int
}

type keyMap struct {
	Quit        key.Binding
	SwitchTheme key.Binding
	Help        key.Binding
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	SwitchTheme: key.NewBinding(
		key.WithKeys("ctrl+t"),
		key.WithHelp("ctrl+t", "switch theme"),
	),
	Help: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "toggle help"),
	),
}

// New returns a new application instance.
func New(app *app.App) tea.Model {
	startPage := pages.Chat

	a := new(application)
	a.app = app

	a.currentPage = startPage

	a.pages = make(map[pages.ID]tea.Model)
	a.pages[pages.Chat] = pages.NewChat(app)

	a.status = status.NewStatus()
	a.helpDialog = dialogs.NewHelp()
	a.themeDialog = dialogs.NewTheme()

	return a
}

// Init initializes the application and its components.
func (a application) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmd := a.pages[a.currentPage].Init()
	cmds = append(cmds, cmd)

	cmd = a.status.Init()
	cmds = append(cmds, cmd)

	cmd = a.themeDialog.Init()
	cmds = append(cmds, cmd)

	cmd = a.helpDialog.Init()
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

// Update handles the application state updates based on messages.
func (a application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		msg.Height -= 1 // Make space for the status bar
		a.width, a.height = msg.Width, msg.Height

		s, _ := a.status.Update(msg)
		a.status = s.(status.Status)
		a.pages[a.currentPage], cmd = a.pages[a.currentPage].Update(msg)
		cmds = append(cmds, cmd)

		help, helpCmd := a.helpDialog.Update(msg)
		a.helpDialog = help.(dialogs.Help)
		cmds = append(cmds, helpCmd)

		return a, tea.Batch(cmds...)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			if a.showThemeDialog {
				a.showThemeDialog = false
			}

			if a.showHelpDialog {
				a.showHelpDialog = false
			}

			return a, tea.Quit
		case key.Matches(msg, keys.SwitchTheme):
			a.showThemeDialog = true
			return a, a.themeDialog.Init()
		case key.Matches(msg, keys.Help):
			a.showHelpDialog = true
			return a, a.helpDialog.Init()
		}
	}

	if a.showThemeDialog {
		d, themeCmd := a.themeDialog.Update(msg)
		a.themeDialog = d.(dialogs.ThemeDialog)
		cmds = append(cmds, themeCmd)

		if _, ok := msg.(tea.KeyMsg); ok {
			return a, tea.Batch(cmds...)
		}
	}

	if a.showHelpDialog {
		d, helpCmd := a.helpDialog.Update(msg)
		a.helpDialog = d.(dialogs.Help)
		cmds = append(cmds, helpCmd)

		if _, ok := msg.(tea.KeyMsg); ok {
			return a, tea.Batch(cmds...)
		}
	}

	s, _ := a.status.Update(msg)
	a.status = s.(status.Status)
	a.pages[a.currentPage], cmd = a.pages[a.currentPage].Update(msg)
	cmds = append(cmds, cmd)

	return a, tea.Batch(cmds...)
}

// View renders the application view.
func (a application) View() string {
	components := []string{
		a.pages[a.currentPage].View(),
	}
	components = slices.Append(components, a.status.View())

	appView := lipgloss.JoinVertical(lipgloss.Top, components...)

	if a.showThemeDialog {
		overlay := a.themeDialog.View()
		row := lipgloss.Height(appView) / 2
		row -= lipgloss.Height(overlay) / 2
		col := lipgloss.Width(appView) / 2
		col -= lipgloss.Width(overlay) / 2
		appView = layout.PlaceOverlay(
			col,
			row,
			overlay,
			appView,
			true,
		)
	}

	if a.showHelpDialog {
		bindings := layout.KeyMapToSlice(keys)
		if p, ok := a.pages[a.currentPage].(layout.Bindings); ok {
			bindings = append(bindings, p.BindingKeys()...)
		}

		a.helpDialog.SetBindings(bindings)

		overlay := a.helpDialog.View()
		row := lipgloss.Height(appView) / 2
		row -= lipgloss.Height(overlay) / 2
		col := lipgloss.Width(appView) / 2
		col -= lipgloss.Width(overlay) / 2
		appView = layout.PlaceOverlay(
			col,
			row,
			overlay,
			appView,
			true,
		)
	}

	return appView
}
