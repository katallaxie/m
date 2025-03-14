package messages

import (
	"time"
)

// Message is a message.
type Message interface {
	// Content ...
	Content() string
	// Timestamp ...
	Timestamp() time.Time
}

// MessageImpl ...
type MessageImpl struct {
	content   string
	message   string
	timestamp time.Time
}

// Value ...
func (m *MessageImpl) Content() string {
	return m.message
}

// Message ...
func (m *MessageImpl) Message() string {
	return m.message
}

// Timestamp ...
func (m *MessageImpl) Timestamp() time.Time {
	return m.timestamp
}

// SetTimestamp ...
func (m *MessageImpl) SetTimestamp(t time.Time) {
	m.timestamp = t
}

// SetContent ...
func (m *MessageImpl) SetContent(c string) {
	m.content = c
}

// NewMessage ...
func NewMessage(msg string) Message {
	return &MessageImpl{
		message:   msg,
		timestamp: time.Now(),
	}
}
