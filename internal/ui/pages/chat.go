package pages

import (
	"github.com/katallaxie/m/internal/app"
	"github.com/katallaxie/m/internal/ui/layout"
	"github.com/katallaxie/m/internal/ui/prompt"
	"github.com/katallaxie/m/internal/ui/sidebar"

	tea "github.com/charmbracelet/bubbletea"
)

var Chat ID = "chat"

type chat struct {
	app    *app.App
	layout layout.SplitPaneLayout
}

func (c *chat) Init() tea.Cmd {
	cmds := []tea.Cmd{
		c.layout.Init(),
	}

	return tea.Batch(cmds...)
}

func (c *chat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	cmd := c.setSidebar()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmd := c.layout.SetSize(msg.Width, msg.Height)
		cmds = append(cmds, cmd)
	}

	u, cmd := c.layout.Update(msg)
	cmds = append(cmds, cmd)
	c.layout = u.(layout.SplitPaneLayout)

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

func (c *chat) setSidebar() tea.Cmd {
	sidebarContainer := layout.NewContainer(
		sidebar.New(),
		layout.WithPadding(1, 1, 1, 1),
	)

	return tea.Batch(c.layout.SetRightPanel(sidebarContainer), sidebarContainer.Init())
}

func NewChat(app *app.App) tea.Model {
	c := new(chat)
	c.app = app

	prompt := layout.NewContainer(
		prompt.NewPrompt(app),
		layout.WithBorder(true, false, false, false),
	)

	c.layout = layout.NewSplitPane(
		layout.WithBottomPanel(prompt),
	)

	return c
}
