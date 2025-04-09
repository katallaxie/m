package keymap

import (
	"fmt"

	"github.com/katallaxie/m/internal/cmd"
)

// Struct that holds a key and a command
type Bind struct {
	// Key is the key that is bound to the command
	Key Key
	// Cmd is the command that is bound to the key
	Cmd cmd.Command
	// Description is a description of the command
	Description string
}

// String returns a string representation of the binding
func (b Bind) String() string {
	return fmt.Sprintf("%s = %s", b.Key.String(), b.Cmd.String())
}
