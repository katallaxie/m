package models

import (
	"github.com/google/uuid"
)

// Chat represents a chat session.
type Chat struct {
	ID       uuid.UUID
	Messages []Message
	manager  *ChatManager
}

// ChatManager is a struct that manages chat sessions.
type ChatManager struct {
	chats   map[uuid.UUID]*Chat
	current *Chat
}

// NewChatManager creates a new chat manager.
func NewChatManager() *ChatManager {
	return &ChatManager{
		chats: map[uuid.UUID]*Chat{},
	}
}

// AddMessages adds messages to the current chat session.
func (c *Chat) AddMessages(messages ...Message) {
	c.Messages = append(c.Messages, messages...)
}

// Current returns the current chat session.
func (cm *ChatManager) Current() *Chat {
	return cm.current
}

// Next creates the next chat session.
func (cm *ChatManager) Next() *Chat {
	c := &Chat{
		ID:       uuid.New(),
		Messages: []Message{},
		manager:  cm,
	}

	cm.current = c
	cm.chats[c.ID] = c

	return c
}
