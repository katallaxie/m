package store

import (
	"github.com/google/uuid"
	"github.com/katallaxie/m/internal/models"
	"github.com/katallaxie/pkg/redux"
)

// AddChatMsg ...
type AddChatMsg struct {
	Chat *models.Chat
}

// NewAddChat ...
func NewAddChat(chat *models.Chat) func() redux.Update {
	return func() redux.Update {
		return &AddChatMsg{
			Chat: chat,
		}
	}
}

// AddCompletionMsg ...
type AddCompletionMsg struct {
	NotebookID uuid.UUID
	Content    string
}
