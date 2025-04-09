package ui

import (
	"github.com/katallaxie/pkg/redux"
)

// Application ...
type Application[S redux.State] interface {
	QueueUpdateDraw(f func())
	GetState() S
	GetStore() redux.Store[S]
	Stop()
	Draw()
}
