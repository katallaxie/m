package chats

import (
	"encoding/json"
	"time"
)

// Chat ...
type Chat struct {
	// Model is the model name, as in [GenerateRequest].
	Model string `json:"model"`
	// Messages is the list of messages to send to the model.
	Messages []Message `json:"messages"`
	// Format is the format to return the response in (e.g. "json").
	Format json.RawMessage `json:"format,omitempty"`
	// Options lists model-specific options.
	Options map[string]interface{} `json:"options"`
	// Suffix is the suffix to append to the response.
	Suffix string `json:"suffix"`
	// Prefix is the prefix to prepend to the response.
	Prefix string `json:"prefix"`
}

// NewChat ...
func NewChat(model string, messages []Message, format json.RawMessage, options map[string]interface{}, suffix, prefix string) Chat {
	return Chat{
		Model:    model,
		Messages: messages,
		Format:   format,
		Options:  options,
		Suffix:   suffix,
		Prefix:   prefix,
	}
}

// Message ...
type Message struct {
	// Role is the role of the message sender.
	Role string `json:"role"`
	// Content is the message content.
	Content string `json:"content"`
	// Timestamp is the time the message was sent.
	Timestamp time.Time `json:"timestamp"`
}
