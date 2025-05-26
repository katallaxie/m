package models

import "github.com/google/uuid"

// Session is a struct that represents a session in the application.
type Session struct {
	// ID is the primary key
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()" params:"id"`
	// Name is the name of the session.
	Name string `json:"name" gorm:"type:varchar(255);not null;default:'Untitled'" params:"name"`
}
