package store

import (
	"github.com/katallaxie/pkg/redux"
)

// AddNotebookReducer ...
func AddNotebookReducer(prev State, msg redux.Msg) State {
	switch m := msg.(type) {
	case AddNotebookMsg:
		prev.Notebooks[m.Notebook.ID] = m.Notebook
		prev.CurrentNotebook = m.Notebook.ID
	}

	return prev
}

// ChatMessageReducer ...
func ChatMessageReducer(curr State, msg redux.Msg) State {
	switch m := msg.(type) {
	case AddChatMsg:
		notebook := curr.Notebooks[m.NotebookID]
		notebook.AddMessages(m.Message)
		curr.Notebooks[m.NotebookID] = notebook
	}

	return curr
}

// // UpdateMessageReducer ...
// func UpdateMessageReducer(prev State, action redux.Action) State {
// 	if action.Type() != UpdateMessage {
// 		return prev
// 	}

// 	payload := action.Payload().(UpdateMessagePayload)
// 	notebook := prev.Notebooks[payload.NotebookID]
// 	notebook.UpdateMessages(payload.Message)

// 	return prev
// }

// // AddMessageReducer ...
// func AddMessageReducer(prev State, action redux.Action) State {
// 	if action.Type() != AddMessage {
// 		return prev
// 	}

// 	payload := action.Payload().(AddMessagePayload)
// 	notebook := prev.Notebooks[payload.NotebookID]
// 	notebook.AddMessages(payload.Message)

// 	return prev
// }

// // SetStatusReducer ...
// func SetStatusReducer(prev State, action redux.Action) State {
// 	if action.Type() != SetStatus {
// 		return prev
// 	}

// 	status := action.Payload().(int)
// 	prev.Status = status

// 	return prev
// }
