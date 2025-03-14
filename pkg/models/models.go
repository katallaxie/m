package models

import (
	"context"

	"github.com/katallaxie/m/pkg/chats"
)

// Chatter ...
type Chatter interface {
	// Generate ...
	Generate(ctx context.Context, chat chats.Chat, opts ...Opt) (chats.Chat, error)
}
