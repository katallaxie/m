package prompt

import (
	"github.com/katallaxie/m/internal/models"
	pctx "github.com/katallaxie/m/internal/ui/context"
	"github.com/katallaxie/m/internal/ui/keys"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx *pctx.ProgramContext
	ta  textarea.Model
}

func NewModel(ctx *pctx.ProgramContext) Model {
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
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, keys.Keys.Submit) {
			v := m.ta.Value()
			cmd = func() tea.Msg {
				return pctx.PromptMsg{
					Messages: []models.Message{
						models.NewSystemMessage().SetContent("You are a helpful assistant. You start every answers with 'Sure!'"),
						models.NewUserMessage().SetContent(v),
					},
				}
			}
			cmds = append(cmds, cmd)

			m.ta.Reset()
			m.ta.Placeholder = "Send a message..."
		}
	}

	m.ta, cmd = m.ta.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View returns the view of the model.
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

// SetWidth sets the width of the model.
func (m Model) SetWidth(width int) {
	m.ta.SetWidth(width)
}

// SetHeight sets the height of the model.
func (m Model) SetHeight(height int) {
	m.ta.SetHeight(height)
}

// Value returns the value of the model.
func (m Model) Value() string {
	return m.ta.Value()
}

// Reset resets the model.
func (m Model) Reset() {
	m.ta.Reset()
}

// UpdateProgramContext updates the program context of the model.
func (m *Model) UpdateProgramContext(ctx *pctx.ProgramContext) {
	m.ctx = ctx
}
