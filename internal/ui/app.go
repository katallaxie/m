package ui

import (
	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/context"

	"github.com/katallaxie/pkg/redux"
	"github.com/rivo/tview"
)

// Application ...
type Application[S redux.State] interface {
	// Context returns the context of the application.
	Context() *context.ProgramContext
	// Config returns the configuration of the application.
	Config() *config.Config
	// QuweueUpdateDraw adds a function to the queue to be executed in the main thread.
	QueueUpdateDraw(f func())
	// Init initializes the application.
	Init() error
	GetState() S
	GetStore() redux.Store[S]
	Stop()
	Draw()
	Pages() *tview.Pages
}
