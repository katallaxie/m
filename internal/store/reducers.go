package store

import (
	"github.com/katallaxie/pkg/redux"
)

// AddNotebookReducer ...
func AddNotebookReducer(prev State, action redux.Action) State {
	if action.Type() != AddNotebook {
		return prev
	}

	payload := action.Payload().(AddNotebookPayload)
	prev.Notebooks[payload.Notebook.ID] = payload.Notebook

	return prev
}

// AddMessageReducer ...
func AddMessageReducer(prev State, action redux.Action) State {
	if action.Type() != AddMessage {
		return prev
	}

	payload := action.Payload().(AddMessagePayload)
	notebook := prev.Notebooks[payload.NotebookID]
	notebook.AddMessages(payload.Message)

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
