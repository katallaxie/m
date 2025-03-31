package ui

import (
	"github.com/katallaxie/pkg/fsmx"
)

// Application ...
type Application interface {
	QueueUpdateDraw(f func())
	GetState() fsmx.Storable
}
