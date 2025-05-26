package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Session is a struct that represents a session in the application.
type Session struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	// Name is the name of the session.
	Name string `json:"name" gorm:"type:varchar(255);not null;default:'Untitled'"`
	// CreatedAt ...
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt ...
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt ...
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
