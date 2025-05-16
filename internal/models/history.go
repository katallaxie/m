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
	// Active is a boolean that indicates if the chat is active.
	Active bool `json:"active"`
	// Messages is a slice of messages in the chat.
	Messages []Message `json:"messages"`
}

// NewChat creates a new chat with the given name and description.
func NewChat() *Chat {
	return &Chat{
		ID:     uuid.New(),
		Name:   defaultName,
		Active: true,
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

// Active returns the active chat in the history.
func (h *History) Active() *Chat {
	for _, chat := range h.Chats {
		if chat.Active {
			return chat
		}
	}

	return nil
}
