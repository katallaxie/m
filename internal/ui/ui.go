package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/ui/components/footer"
	"github.com/katallaxie/m/internal/ui/context"
	"github.com/katallaxie/m/internal/ui/keys"
)

var _ tea.Model = (*Model)(nil)

type initMsg struct {
	Config config.Config
}

const (
	FooterHeight      = 1
	TabsBorderHeight  = 1
	TabsContentHeight = 2
	TabsHeight        = TabsBorderHeight + TabsContentHeight
)

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
	return tea.Batch(m.initScreen, tea.EnterAltScreen)
}

// Update updates the model with the given message.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)
	case initMsg:
		log.Debug("initMsg", "config", msg.Config)
		m.syncMainContentWidth()
	}

	m.syncProgramContext()

	return m, nil
}

// View returns the view of the model.
func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.footer.View(),
	)
}

func (m *Model) initScreen() tea.Msg {
	return initMsg{}
}

func (m Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.footer.SetWidth(msg.Width)
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	m.ctx.MainContentHeight = msg.Height - TabsHeight - FooterHeight

	m.syncMainContentWidth()
}

func (m *Model) syncProgramContext() {}

func (m *Model) syncMainContentWidth() {
	sideBarOffset := 0
	m.ctx.MainContentWidth = m.ctx.ScreenWidth - sideBarOffset
}
