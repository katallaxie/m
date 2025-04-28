package ui

import (
	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/ui/components/footer"
	"github.com/katallaxie/m/internal/ui/components/prompt"
	pctx "github.com/katallaxie/m/internal/ui/context"
	"github.com/katallaxie/m/internal/ui/keys"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
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
	answering  bool
	ctx        *pctx.ProgramContext
	err        error
	footer     footer.Model
	height     int
	historyIdx int
	keys       *keys.KeyMap
	prompt     prompt.Model
	renderer   *glamour.TermRenderer
	vp         viewport.Model
	width      int
}

// New creates a new model.
func New(ctx *pctx.ProgramContext) Model {
	m := Model{}
	m.ctx = ctx

	m.footer = footer.NewModel(m.ctx)
	m.keys = keys.Keys

	p := prompt.NewModel(m.ctx)
	m.prompt = p

	vp := viewport.New(50, 5)
	m.vp = vp

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(0), // we do hard-wrapping ourselves
	)
	m.renderer = renderer
	m = m.SetInputMode(keys.InputModelMultiLine)

	return m
}

// Init returns the initial command for the model.
func (m Model) Init() tea.Cmd {
	return tea.Batch(m.initScreen, tea.EnterAltScreen)
}

// Update updates the model with the given message.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.prompt, cmd = m.prompt.Update(msg)
	cmds = append(cmds, cmd)

	m.vp, cmd = m.vp.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.footer, cmd = m.footer.Update(msg)
		cmds = append(cmds, cmd)
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}

		if key.Matches(msg, m.keys.Submit) {
			m.footer, cmd = m.footer.Update(msg)
			cmds = append(cmds, cmd)
		}
	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)
	case initMsg:
		m.syncMainContentWidth()
	}

	// synchronize the program context before updating the view
	m.syncProgramContext()

	return m, tea.Batch(cmds...)
}

// View returns the view of the model.
func (m Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.vp.View(),
		m.prompt.View(),
		m.footer.View(),
	)
}

func (m Model) SetInputMode(mode keys.InputMode) Model {
	return m
}

func (m *Model) initScreen() tea.Msg {
	return initMsg{}
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.footer.SetWidth(msg.Height)
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	m.ctx.MainContentHeight = msg.Height - TabsHeight - FooterHeight
	m.vp.Width = msg.Width
	m.vp.Height = msg.Height - m.prompt.Height() - lipgloss.Height(m.footer.View())
	m.prompt.SetWidth(msg.Width)

	m.syncMainContentWidth()
}

func (m *Model) syncProgramContext() {
	m.footer.UpdateProgramContext(m.ctx)
	m.prompt.UpdateProgramContext(m.ctx)
}

func (m *Model) syncMainContentWidth() {
	sideBarOffset := 0
	m.ctx.MainContentWidth = m.ctx.ScreenWidth - sideBarOffset
}
