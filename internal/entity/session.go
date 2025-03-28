package entity

import (
	"github.com/google/uuid"
)

type Session struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// NewSession returns a new session.
func NewSession() Session {
	return Session{
		ID: uuid.New(),
	}
}
