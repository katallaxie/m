package keymap

import (
	"fmt"

	"github.com/katallaxie/m/internal/cmd"
)

// Struct that holds a key and a command
type Bind struct {
	Key         Key
	Cmd         cmd.Command
	Description string
}

func (b Bind) String() string {
	return fmt.Sprintf("%s = %s", b.Key.String(), b.Cmd.String())
}
