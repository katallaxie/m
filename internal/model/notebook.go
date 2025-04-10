package model

import (
	"github.com/google/uuid"
	"github.com/katallaxie/pkg/slices"
)

const defaultName = "Untitled"

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
