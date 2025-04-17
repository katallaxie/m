package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/katallaxie/m/internal/ui/components/footer"
	"github.com/katallaxie/m/internal/ui/context"
	"github.com/katallaxie/m/internal/ui/keys"
)

var _ tea.Model = (*Model)(nil)

// Model ...
type Model struct {
	width      int
	height     int
	historyIdx int
	answering  bool
	err        error
	footer     footer.Model
	spinner    spinner.Model
	keys       *keys.KeyMap
	ctx        *context.ProgramContext
}

// New creates a new model.
func New() Model {
	m := Model{}

	m.spinner = spinner.Model{Spinner: spinner.Dot}
	m.footer = footer.NewModel(m.ctx)
	m.keys = keys.Keys

	m.ctx = &context.ProgramContext{}

	return m
}

// Init returns the initial command for the model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update updates the model with the given message.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View returns the view of the model.
func (m Model) View() string {
	return ""
}
