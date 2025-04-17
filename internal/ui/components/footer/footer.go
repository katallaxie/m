package footer

import (
	"github.com/katallaxie/m/internal/ui/context"
	"github.com/katallaxie/m/internal/ui/keys"

	bbHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx             *context.ProgramContext
	leftSection     *string
	rightSection    *string
	help            bbHelp.Model
	ShowAll         bool
	ShowConfirmQuit bool
}

func NewModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx: ctx,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
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

	return m, nil
}

func (m Model) View() string {
	if m.ShowConfirmQuit {
		return lipgloss.NewStyle().Render("Really quit? (Press q/esc again to quit)")
	}

	helpIndicator := lipgloss.NewStyle().
		Padding(0, 1).
		Render("? help")

	return helpIndicator
}

func (m *Model) SetWidth(width int) {
	m.help.Width = width
}
