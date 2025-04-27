package footer

import (
	"fmt"
	"strings"

	"github.com/katallaxie/m/internal/ui/context"
	"github.com/katallaxie/m/internal/ui/keys"

	bbHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	footerStyle = lipgloss.NewStyle().
		Height(1).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("8")).
		Faint(true)
)

type Model struct {
	ctx             *context.ProgramContext
	spinner         spinner.Model
	leftSection     *string
	rightSection    *string
	help            bbHelp.Model
	ShowAll         bool
	ShowConfirmQuit bool
}

func NewModel(ctx *context.ProgramContext) Model {
	m := Model{}
	m.ctx = ctx
	m.spinner = spinner.Model{Spinner: spinner.Dot}

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Submit):
			cmd = func() tea.Msg {
				return m.spinner.Tick()
			}
		case key.Matches(msg, keys.Keys.Quit):
			if m.ShowConfirmQuit {
				return m, tea.Quit
			}
			m.ShowConfirmQuit = true
		case m.ShowConfirmQuit && !key.Matches(msg, keys.Keys.Quit):
			m.ShowConfirmQuit = false
		case key.Matches(msg, keys.Keys.Help):
			m.ShowAll = !m.ShowAll
		}
	}

	return m, cmd
}

func (m Model) View() string {
	if m.ShowConfirmQuit {
		return lipgloss.NewStyle().Render("Really quit? (Press q/esc again to quit)")
	}

	var columns []string
	columns = append(columns, m.spinner.View())
	columns = append(columns, fmt.Sprintf("%s ctrl+h", "? "))

	footer := strings.Join(columns, " ")
	footer = footerStyle.Render(footer)

	return footer
}

func (m *Model) SetWidth(width int) {
	m.help.Width = width
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
