package state

import (
	"github.com/katallaxie/pkg/fsmx"
)

// Actions ...
const (
	AddMessage fsmx.ActionType = iota
)

// AddMessagePayload ...
type AddMessagePayload struct {
	// ID ...
	ID string
	// Message ...
	Message string
}

// NewAddMessage returns a new action.
func NewAddMessage(message string) fsmx.Action {
	return fsmx.NewAction(AddMessage, AddMessagePayload{Message: message})
}

// AddMessageReducer ...
func AddMessageReducer(prev fsmx.State, action fsmx.Action) fsmx.State {
	if action.Type() != AddMessage {
		return prev
	}

	state := prev.(*State)
	payload := action.Payload().(AddMessagePayload)
	state.Messages = append(state.Messages, payload.Message)

	return state
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
func NewState() *State {
	return &State{
		Status:   Initial,
		Messages: make([]string, 0),
	}
}
