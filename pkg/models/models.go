package models

import (
	"context"

	"github.com/katallaxie/m/pkg/messages"
)

// Value is a message value.
type Value interface {
	int | ~string | []byte
}

// Messages is a channel of messages.
type Messages chan messages.Message

// Model ...
type Model interface {
	// Generate ...
	Generate(ctx context.Context, input []messages.Message, opts ...Opt) (messages.Message, error)
}
