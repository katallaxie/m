package status

import (
	"time"

	"github.com/katallaxie/m/internal/ui/styles"
	"github.com/katallaxie/m/internal/ui/theme"
	"github.com/katallaxie/m/internal/ui/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type Status interface {
	tea.Model
}

type status struct {
	width int
}

func (s status) clearMessageCmd(ttl time.Duration) tea.Cmd {
	return tea.Tick(ttl, func(time.Time) tea.Msg {
		return utils.ClearStatusMsg{}
	})
}

func (s status) Init() tea.Cmd {
	return nil
}

func (s status) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width = msg.Width
		return s, nil
	}
	return s, nil
}

// getHelpWidget returns the help widget with current theme colors
func getHelpWidget() string {
	t := theme.Current()

	helpText := "ctrl+? help"

	return styles.Padded().
		Background(t.TextMuted()).
		Foreground(t.BackgroundDarker()).
		Bold(true).
		Render(helpText)
}

func (s status) View() string {
	// t := theme.Current()

	// Initialize the help widget
	status := getHelpWidget()

	return status
}

func (m status) model() string {
	t := theme.Current()

	return styles.Padded().
		Background(t.Secondary()).
		Foreground(t.Background()).
		Render("slm")
}

func NewStatus() Status {
	return &status{}
}
