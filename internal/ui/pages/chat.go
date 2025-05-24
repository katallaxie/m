package pages

import (
	"github.com/katallaxie/m/internal/app"
	"github.com/katallaxie/m/internal/ui/layout"

	tea "github.com/charmbracelet/bubbletea"
)

var Chat ID = "chat"

type chat struct {
	app    *app.App
	layout layout.SplitPaneLayout
}

func (c *chat) Init() tea.Cmd {
	cmds := []tea.Cmd{}

	return tea.Batch(cmds...)
}

func (c *chat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	return c, tea.Batch(cmds...)
}

func (c *chat) SetSize(width, height int) tea.Cmd {
	return c.layout.SetSize(width, height)
}

func (c *chat) GetSize() (int, int) {
	return c.layout.GetSize()
}

func (c *chat) View() string {
	view := c.layout.View()

	return view
}

func NewChat(app *app.App) tea.Model {
	return &chat{
		app: app,
	}
}
