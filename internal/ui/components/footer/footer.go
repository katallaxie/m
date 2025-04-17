package footer

import (
	"github.com/katallaxie/m/internal/ui/context"

	bbHelp "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
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
	return m, nil
}

func (m Model) View() string {
	return ""
}

func (m *Model) SetWidth(width int) {
	m.help.Width = width
}
