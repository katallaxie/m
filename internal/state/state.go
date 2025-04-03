package state

import (
	"github.com/katallaxie/pkg/fsmx"
)

// Actions ...
const (
	AddMessage fsmx.ActionType = iota
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
func NewSetStatus(status int) fsmx.Action {
	return fsmx.NewAction(SetStatus, status)
}

// NewAddMessage returns a new action.
func NewAddMessage(message string) fsmx.Action {
	return fsmx.NewAction(AddMessage, AddMessagePayload{
		Message: message,
	})
}

// AddMessageReducer ...
func AddMessageReducer(prev State, action fsmx.Action) State {
	if action.Type() != AddMessage {
		return prev
	}

	payload := action.Payload().(AddMessagePayload)
	prev.Messages = append(prev.Messages, payload.Message)
	prev.Status = Success

	return prev
}

// SetStatusReducer ...
func SetStatusReducer(prev State, action fsmx.Action) State {
	if action.Type() != SetStatus {
		return prev
	}

	status := action.Payload().(int)
	prev.Status = status

	return prev
}

const (
	Initial = iota
	Loading
	Error
	Success
)

// State ...
type State struct {
	Status   int
	Messages []string
}

// NewState returns a new state.
func NewState() State {
	return State{
		Status:   Initial,
		Messages: make([]string, 0),
	}
}
