package chat

import (
	"strings"

	"github.com/katallaxie/m/internal/models"
	"github.com/katallaxie/m/internal/ui/context"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	senderStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5"))
)

type Model struct {
	vp       viewport.Model
	ctx      *context.ProgramContext
	renderer *glamour.TermRenderer
}

func NewModel(ctx *context.ProgramContext) Model {
	m := Model{}
	m.ctx = ctx

	vp := viewport.New(50, 5)
	m.vp = vp

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(0), // we do hard-wrapping ourselves
	)
	m.renderer = renderer

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var sb strings.Builder
	curr := m.ctx.Chats.Current()

	for _, msg := range curr.Messages {
		switch msg := msg.(type) {
		case *models.UserMessage:
			sb.WriteString(senderStyle.Render("You: "))

			content := msg.Content()
			content = wordwrap.String(content, 50)
			content, _ = m.renderer.Render(content)

			sb.WriteString(content)
		case *models.AIMessage:
			sb.WriteString(senderStyle.Render("AI: "))
			content := msg.Content()
			content = wordwrap.String(content, 50)
			content, _ = m.renderer.Render(content)
			sb.WriteString(content)
		}
	}

	m.vp.SetContent(sb.String())
	m.vp.GotoBottom()

	return m, nil
}

func (m Model) View() string {
	return m.vp.View()
}

func (m *Model) SetWidth(width int) {
	m.vp.Width = width
}

func (m *Model) SetHeight(height int) {
	m.vp.Height = height
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
}
