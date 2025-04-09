package store

import "github.com/katallaxie/pkg/redux"

// Actions ...
const (
	AddMessage redux.ActionType = iota
	SetStatus
)

// AddMessagePayload ...
type AddMessagePayload struct {
	// ID ...
	ID string
	// Message ...
	Message string
}

// NewSetStatus returns a new action.
func NewSetStatus(status int) redux.Action {
	return redux.NewAction(SetStatus, status)
}

// NewAddMessage returns a new action.
func NewAddMessage(message string) redux.Action {
	return redux.NewAction(AddMessage, AddMessagePayload{
		Message: message,
	})
}
