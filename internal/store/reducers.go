package store

import "github.com/katallaxie/pkg/redux"

// AddMessageReducer ...
func AddMessageReducer(prev State, action redux.Action) State {
	if action.Type() != AddMessage {
		return prev
	}

	payload := action.Payload().(AddMessagePayload)
	prev.Messages = append(prev.Messages, payload.Message)

	return prev
}

// SetStatusReducer ...
func SetStatusReducer(prev State, action redux.Action) State {
	if action.Type() != SetStatus {
		return prev
	}

	status := action.Payload().(int)
	prev.Status = status

	return prev
}
