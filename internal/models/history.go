package models

import (
	"github.com/google/uuid"
	"github.com/katallaxie/pkg/slices"
)

const defaultName = "Untitled"

// Chat is a struct that represents a chat in the application.
type Chat struct {
	// ID is the unique identifier for the chat.
	ID uuid.UUID `json:"id"`
	// Name is the name of the chat.
	Name string `json:"name"`
	// Desc is the description of the chat.
	Desc string `json:"desc"`
}

// NewChat creates a new chat with the given name and description.
func NewChat() *Chat {
	return &Chat{
		ID:   uuid.New(),
		Name: defaultName,
	}
}

// History is a struct that represents the history of a chat.
type History struct {
	// Chats is a slice of chats in the history.
	Chats []*Chat `json:"chats"`
}

// NewHistory creates a new history with the given chats.
func NewHistory() *History {
	return &History{
		Chats: []*Chat{},
	}
}

// Append adds a chat to the history.
func (h *History) Append(chat ...*Chat) {
	h.Chats = slices.Append(h.Chats, chat...)
}

// Next returns the next chat in the history.
func (h *History) Next() *Chat {
	chat := NewChat()
	h.Append(chat)

	return chat
}

// Notebook is a struct that represents a notebook in the application.
type Notebook struct {
	// ID is the unique identifier for the notebook.
	ID uuid.UUID `json:"id"`
	// Name is the name of the notebook.
	Name string `json:"name"`
	// Desc is the description of the notebook.
	Desc string `json:"desc"`
	// Messages is a slice of messages in the notebook.
	Messages []Message `json:"messages"`
}

// NewNotebook creates a new notebook with the given name and description.
func NewNotebook() Notebook {
	return Notebook{
		ID:       uuid.New(),
		Name:     defaultName,
		Messages: []Message{},
	}
}

// AddMessage adds a message to the notebook.
func (n *Notebook) AddMessages(messages ...Message) {
	n.Messages = slices.Append(n.Messages, messages...)
}

// UpdateMessages updates the messages in the notebook.
func (n *Notebook) UpdateMessages(messages ...Message) {
	for i := range messages {
		for j := range n.Messages {
			if n.Messages[j].ID() == messages[i].ID() {
				n.Messages[j] = messages[i]
			}
		}
	}
}
