package ui

import (
	"context"

	"github.com/katallaxie/pkg/redux"
)

// Application ...
type Application[S redux.State] interface {
	// Context returns the context of the application.
	Context() context.Context
	QueueUpdateDraw(f func())
	GetState() S
	GetStore() redux.Store[S]
	Stop()
	Draw()
}
