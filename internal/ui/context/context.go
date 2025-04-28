package context

import (
	"context"
	"time"

	"github.com/katallaxie/m/internal/config"

	tea "github.com/charmbracelet/bubbletea"
)

// NewProgramContext creates a new ProgramContext with default values.
func NewProgramContext() *ProgramContext {
	return &ProgramContext{}
}

// WithContext creates a new ProgramContext with the given context.
func WithContext(ctx context.Context) *ProgramContext {
	return &ProgramContext{
		ctx: ctx,
	}
}

// ProgramContext holds the context for the program, including the screen size, configuration, and error handling.
type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int
	Config            *config.Config
	ConfigPath        string
	Version           string
	View              config.ViewType
	Error             error
	program           *tea.Program
	ctx               context.Context
}

// SetContext sets the context for the program.
func (c *ProgramContext) SetContext(ctx context.Context) {
	c.ctx = ctx
}

// Context returns the context of the program.
func (c *ProgramContext) Context() context.Context {
	if c.ctx == nil {
		c.ctx = context.Background()
	}

	return c.ctx
}

// Send s a message to the program's message channel.
func (c *ProgramContext) Send(msg tea.Msg) {
	c.program.Send(msg)
}

// SetProgram sets the program for the context.
func (c *ProgramContext) SetProgram(p *tea.Program) {
	c.program = p
}

type State = int

const (
	TaskStart State = iota
	TaskFinished
	TaskError
)

// Task represents a task with an ID, start and finished text, state, error, start time, and finished time.
type Task struct {
	Id           string
	StartText    string
	FinishedText string
	State        State
	Error        error
	StartTime    time.Time
	FinishedTime *time.Time
}
