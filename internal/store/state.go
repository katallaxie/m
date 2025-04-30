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
	History         *models.History
	Notebooks       map[uuid.UUID]models.Notebook
}

// NewState returns a new state.
func NewState() State {
	return State{
		Status:    Initial,
		History:   models.NewHistory(),
		Notebooks: make(map[uuid.UUID]models.Notebook),
	}
}
