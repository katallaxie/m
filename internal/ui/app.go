package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/katallaxie/m/internal/app"
	"github.com/katallaxie/m/internal/ui/pages"
	"github.com/katallaxie/m/internal/ui/status"
	"github.com/katallaxie/pkg/slices"
)

type application struct {
	app *app.App

	pages        map[pages.ID]tea.Model
	currentPage  pages.ID
	previousPage pages.ID

	status status.Status

	width, height int
}

type keyMap struct {
	Quit key.Binding
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
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

	return a
}

func (a application) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmd := a.pages[a.currentPage].Init()
	cmds = append(cmds, cmd)

	cmd = a.status.Init()
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

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

		return a, tea.Batch(cmds...)
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, keys.Quit):
			return a, tea.Quit
		}
	}

	s, _ := a.status.Update(msg)
	a.status = s.(status.Status)
	a.pages[a.currentPage], cmd = a.pages[a.currentPage].Update(msg)
	cmds = append(cmds, cmd)

	return a, tea.Batch(cmds...)
}

func (a application) View() string {
	components := []string{
		a.pages[a.currentPage].View(),
	}
	components = slices.Append(components, a.status.View())

	appView := lipgloss.JoinVertical(lipgloss.Top, components...)

	return appView
}
