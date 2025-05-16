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

// AddChatMessageMsg ...
type AddChatMessageMsg struct {
	ChatID  uuid.UUID
	Message *models.Message
}

// NewAddChatMessage ...
func NewAddChatMessage(chatID uuid.UUID, message *models.Message) func() redux.Update {
	return func() redux.Update {
		return &AddChatMessageMsg{
			ChatID:  chatID,
			Message: message,
		}
	}
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
