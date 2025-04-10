package store

import (
	"github.com/google/uuid"
	"github.com/katallaxie/m/internal/model"
	"github.com/katallaxie/pkg/redux"
)

// Actions ...
const (
	AddMessage redux.ActionType = iota
	AddNotebook
	SetStatus
)

// AddNotebookPayload ...
type AddNotebookPayload struct {
	Notebook model.Notebook
}

// NewAddNotebook returns a new action.
func NewAddNotebook(notebook model.Notebook) redux.Action {
	return redux.NewAction(AddNotebook, AddNotebookPayload{
		Notebook: notebook,
	})
}

// AddMessagePayload ...
type AddMessagePayload struct {
	NotebookID uuid.UUID
	Message    model.Message
}

// NewSetStatus returns a new action.
func NewSetStatus(status int) redux.Action {
	return redux.NewAction(SetStatus, status)
}

// NewAddMessage returns a new action.
func NewAddMessage(notebookId uuid.UUID, message model.Message) redux.Action {
	return redux.NewAction(AddMessage, AddMessagePayload{
		Message: message,
	})
}
