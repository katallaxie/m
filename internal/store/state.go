package store

import (
	"github.com/katallaxie/m/internal/models"
)

const (
	Initial = iota
	Loading
	Error
	Success
)

// State ...
type State struct {
	Status  int
	Error   error
	History *models.History
}

// NewState returns a new state.
func NewState() State {
	return State{
		Status:  Initial,
		History: models.NewHistory(),
	}
}
