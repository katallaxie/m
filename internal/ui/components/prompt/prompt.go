package prompt

import (
	"github.com/katallaxie/m/internal/ui/context"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx *context.ProgramContext
	ta  textarea.Model
}

func NewModel(ctx *context.ProgramContext) Model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = -1
	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return Model{
		ctx: ctx,
		ta:  ta,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.ta, cmd = m.ta.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.ta.View()
}

// Height returns the height of the model.
func (m Model) Height() int {
	return m.ta.Height()
}

// Width returns the width of the model.
func (m Model) Width() int {
	return m.ta.Width()
}

func (m *Model) SetWidth(width int) {
	m.ta.SetWidth(width)
}

func (m *Model) Reset() {
	m.ta.Reset()
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
