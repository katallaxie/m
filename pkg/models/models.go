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
type Messages[V Value] chan messages.Message[V]

// Model ...
type Model[V any] interface {
	// Generate ...
	Generate(ctx context.Context, input []messages.Message[V], opts ...Opt) (messages.Message[V], error)
}
