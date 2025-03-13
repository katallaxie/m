package messages

import "time"

// Message is a message.
type Message[V any] interface {
	// Value ...
	Value() V
	// Timestamp ...
	Timestamp() time.Time
}
