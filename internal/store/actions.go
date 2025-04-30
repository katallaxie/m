package store

import (
	"github.com/google/uuid"
	"github.com/katallaxie/m/internal/models"
)

// AddNotebookMsg ...
type AddNotebookMsg struct {
	Notebook models.Notebook
}

// AddChatMsg ...
type AddChatMsg struct {
	NotebookID uuid.UUID
	models.Message
}

// AddCompletionMsg ...
type AddCompletionMsg struct {
	NotebookID uuid.UUID
	Content    string
}
