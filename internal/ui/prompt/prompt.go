package prompt

import (
	"github.com/katallaxie/m/internal/app"
	"github.com/katallaxie/m/internal/ui/layout"
	"github.com/katallaxie/m/internal/ui/styles"
	"github.com/katallaxie/m/internal/ui/theme"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type prompt struct {
	width  int
	height int
	app    *app.App
	// session     session.Session
	textarea textarea.Model
	// attachments []message.Attachment
	deleteMode bool
}

type PromptKeyMaps struct {
	Send key.Binding
}

type bluredPromptKeyMaps struct {
	Send  key.Binding
	Focus key.Binding
}

var promptMaps = PromptKeyMaps{
	Send: key.NewBinding(
		key.WithKeys("enter", "ctrl+s"),
		key.WithHelp("enter", "send message"),
	),
}

func (p *prompt) Init() tea.Cmd {
	return textarea.Blink
}

func (p *prompt) send() tea.Cmd {
	return tea.Batch()
}

func (p *prompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if p.textarea.Focused() && key.Matches(msg, promptMaps.Send) {
			value := p.textarea.Value()
			if len(value) > 0 && value[len(value)-1] == '\\' {
				// If the last character is a backslash, remove it and add a newline
				p.textarea.SetValue(value[:len(value)-1] + "\n")
				return p, nil
			} else {
				return p, p.send()
			}
		}

	}

	p.textarea, cmd = p.textarea.Update(msg)

	return p, cmd
}

func (p *prompt) View() string {
	t := theme.Current()

	// Style the prompt with theme colors
	style := lipgloss.NewStyle().
		Padding(0, 0, 0, 1).
		Bold(true).
		Foreground(t.Primary())

	p.textarea.SetHeight(p.height - 1)

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, style.Render(">"),
			p.textarea.View()),
	)
}

func (p *prompt) SetSize(width, height int) tea.Cmd {
	p.width = width
	p.height = height
	p.textarea.SetWidth(width - 3) // account for the prompt and padding right
	p.textarea.SetHeight(height)
	p.textarea.SetWidth(width)

	return nil
}

func (p *prompt) GetSize() (int, int) {
	return p.textarea.Width(), p.textarea.Height()
}

func (p *prompt) BindingKeys() []key.Binding {
	bindings := []key.Binding{}

	bindings = append(bindings, layout.KeyMapToSlice(promptMaps)...)

	return bindings
}

func CreateTextArea(existing *textarea.Model) textarea.Model {
	t := theme.Current()
	bgColor := t.Background()
	textColor := t.Text()
	textMutedColor := t.TextMuted()

	ta := textarea.New()
	ta.BlurredStyle.Base = styles.BaseStyle().Background(bgColor).Foreground(textColor)
	ta.BlurredStyle.CursorLine = styles.BaseStyle().Background(bgColor)
	ta.BlurredStyle.Placeholder = styles.BaseStyle().Background(bgColor).Foreground(textMutedColor)
	ta.BlurredStyle.Text = styles.BaseStyle().Background(bgColor).Foreground(textColor)
	ta.FocusedStyle.Base = styles.BaseStyle().Background(bgColor).Foreground(textColor)
	ta.FocusedStyle.CursorLine = styles.BaseStyle().Background(bgColor)
	ta.FocusedStyle.Placeholder = styles.BaseStyle().Background(bgColor).Foreground(textMutedColor)
	ta.FocusedStyle.Text = styles.BaseStyle().Background(bgColor).Foreground(textColor)

	ta.Prompt = " "
	ta.ShowLineNumbers = false
	ta.CharLimit = -1

	if existing != nil {
		ta.SetValue(existing.Value())
		ta.SetWidth(existing.Width())
		ta.SetHeight(existing.Height())
	}

	ta.Focus()

	return ta
}

func NewPrompt(app *app.App) tea.Model {
	ta := CreateTextArea(nil)
	return &prompt{
		app:      app,
		textarea: ta,
	}
}
