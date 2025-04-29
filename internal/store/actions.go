package store

import (
	"github.com/google/uuid"
	"github.com/katallaxie/m/internal/models"
	"github.com/katallaxie/pkg/redux"
)

// AddNotebookMsg ...
type AddNotebookMsg struct {
	Notebook models.Notebook
}

// AddChatMsg ...
type AddChatMsg struct {
	NotebookID uuid.UUID
	models.Message
}

// AddCompletionMsg ...
type AddCompletionMsg struct {
	NotebookID uuid.UUID
	Content    string
}

// Actions ...
const (
	AddMessage redux.ActionType = iota
	UpdateMessage
	DeleteMessage
	AddNotebook
	SetStatus
)

// AddNotebookPayload ...
type AddNotebookPayload struct {
	Notebook models.Notebook
}

// AddMessagePayload ...
type AddMessagePayload struct {
	NotebookID uuid.UUID
	Message    models.Message
}

// UpdateMessagePayload ...
type UpdateMessagePayload struct {
	NotebookID uuid.UUID
	Message    models.Message
}

// // NewUpdateMessage returns a new action.
// func NewUpdateMessage(notebookId uuid.UUID, message model.Message) redux.Action {
// 	return redux.NewAction(UpdateMessage, UpdateMessagePayload{
// 		NotebookID: notebookId,
// 		Message:    message,
// 	})
// }

// // NewSetStatus returns a new action.
// func NewSetStatus(status int) redux.Action {
// 	return redux.NewAction(SetStatus, status)
// }

// // NewAddMessage returns a new action.
// func NewAddMessage(notebookId uuid.UUID, message model.Message) redux.Action {
// 	return redux.NewAction(AddMessage, AddMessagePayload{
// 		Message: message,
// 	})
// }
