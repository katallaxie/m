package state

import (
	"github.com/katallaxie/pkg/fsmx"
)

// Actions ...
const (
	AddMessage fsmx.Action = iota
)

// AddMessagePayload ...
type AddMessagePayload struct {
	// ID ...
	ID string
	// Message ...
	Message string
}

// NewAddMessage returns a new action.
func NewAddMessage(message string) fsmx.Actionable {
	return fsmx.NewAction(AddMessage, AddMessagePayload{Message: message})
}

// AddMessageReducer ...
func AddMessageReducer(prev fsmx.State, action fsmx.Actionable) fsmx.State {
	if action.GetType() != AddMessage {
		return prev
	}

	state := prev.(*State)
	payload := action.GetPayload().(AddMessagePayload)
	state.Messages = append(state.Messages, payload.Message)

	return state
}

// State ...
type State struct {
	Messages []string
}

// NewState returns a new state.
func NewState() *State {
	return &State{
		Messages: make([]string, 0),
	}
}
