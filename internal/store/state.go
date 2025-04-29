package store

import (
	"github.com/google/uuid"
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
	Status          int
	Error           error
	CurrentNotebook uuid.UUID
	Notebooks       map[uuid.UUID]models.Notebook
}

// NewState returns a new state.
func NewState() State {
	return State{
		Status:    Initial,
		Notebooks: make(map[uuid.UUID]models.Notebook),
	}
}
