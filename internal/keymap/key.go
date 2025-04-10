package keymap

import "github.com/gdamore/tcell/v2"

// Key is a structure that represents a key that can be bound to an command.
type Key struct {
	// Code is the tcell key code
	Code tcell.Key
	// Char is the rune that is associated with the key
	Char rune
	// Mod is the modifier that is associated with the key
	Mod tcell.ModMask
}

// String returns a string representation of the key.
func (k Key) String() string {
	if k.Char != 0 {
		return string(k.Char)
	}

	if desc, ok := tcell.KeyNames[k.Code]; ok {
		return "<" + desc + ">"
	}
	return ""
}
