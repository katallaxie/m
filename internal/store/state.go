package store

import (
	"github.com/google/uuid"
	"github.com/katallaxie/m/internal/model"
)

const (
	Initial = iota
	Loading
	Error
	Success
)

// State ...
type State struct {
	Status    int
	Error     error
	Messages  []string
	Notebooks map[uuid.UUID]model.Notebook
}

// NewState returns a new state.
func NewState() State {
	return State{
		Status:    Initial,
		Messages:  make([]string, 0),
		Notebooks: make(map[uuid.UUID]model.Notebook),
	}
}
