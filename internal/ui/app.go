package ui

import (
	"context"

	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/pkg/redux"
)

// Application ...
type Application[S redux.State] interface {
	// Context returns the context of the application.
	Context() context.Context
	// Config returns the configuration of the application.
	Config() *config.Config
	QueueUpdateDraw(f func())
	GetState() S
	GetStore() redux.Store[S]
	Stop()
	Draw()
}
