package sidebar

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Sidebar struct {
	width, height int
}

// New creates a new Sidebar instance.
func New() *Sidebar {
	s := &Sidebar{}

	return s
}

// Init initializes the sidebar.
func (s *Sidebar) Init() tea.Cmd {
	return nil
}

// Update updates the sidebar state.
func (s *Sidebar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
	}

	return s, nil
}

// View renders the sidebar.
func (s *Sidebar) View() string {
	// For now, we return an empty string as the sidebar view.
	// This can be expanded to include actual sidebar content.
	return ""
}

// SetSize sets the size of the sidebar.
func (s *Sidebar) SetSize(width, height int) tea.Cmd {
	s.width = width
	s.height = height

	return nil
}

// GetSize returns the current size of the sidebar.
func (s *Sidebar) GetSize() (int, int) {
	return s.width, s.height
}
