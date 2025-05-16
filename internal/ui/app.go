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
	// QueueUpdate adds a function to the queue to be executed in the main thread.
	QueueUpdate(f func())
	// Init initializes the application.
	Init() error
	// GetState returns the state of the application.
	GetState() S
	// GetStore returns the store of the application.
	GetStore() redux.Store[S]
	// Stop stops the application.
	Stop()
	// Draw draws the application.
	Draw()
	// Pages returns the pages of the application.
	Pages() *tview.Pages
}
