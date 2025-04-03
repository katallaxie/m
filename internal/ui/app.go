package ui

import (
	"github.com/katallaxie/pkg/fsmx"
)

// Application ...
type Application[S fsmx.State] interface {
	QueueUpdateDraw(f func())
	GetState() S
	GetStore() fsmx.Store[S]
	Stop()
	Draw()
}
