package app

import (
	"github.com/gdamore/tcell/v2"

	"github.com/katallaxie/m/internal/cmd"
	"github.com/katallaxie/m/internal/keymap"
)

// local alias added for clarity purpose
type (
	Bind = keymap.Bind
	Key  = keymap.Key
	Map  = keymap.Map
)

// KeymapSystem is the actual key mapping system.
// A map can have several groups. But it always has a "Global" one.
type KeymapSystem struct {
	Groups map[string]Map
	Global Map
}

func (c KeymapSystem) Group(name string) Map {
	// Lookup the group
	if group, ok := c.Groups[name]; ok {
		return group
	}

	// Did not find any maps. Return a empty one
	return Map{}
}

// Resolve translates a tcell.EventKey into a command based on the mappings in
// the global group
func (c KeymapSystem) Resolve(event *tcell.EventKey) cmd.Command {
	return c.Global.Resolve(event)
}

const (
	ChatGroup = "chat"
)

var Keymaps = KeymapSystem{
	Groups: map[string]Map{
		ChatGroup: {
			Bind{Key: Key{Char: 'L'}, Cmd: cmd.MoveRight, Description: "Focus table"},
			Bind{Key: Key{Char: 'H'}, Cmd: cmd.MoveLeft, Description: "Focus tree"},
			Bind{Key: Key{Code: tcell.KeyCtrlE}, Cmd: cmd.SwitchToEditorView, Description: "Open SQL editor"},
			Bind{Key: Key{Code: tcell.KeyCtrlS}, Cmd: cmd.Save, Description: "Execute pending changes"},
			Bind{Key: Key{Char: 'q'}, Cmd: cmd.Quit, Description: "Quit"},
			Bind{Key: Key{Code: tcell.KeyBackspace2}, Cmd: cmd.SwitchToConnectionsView, Description: "Switch to connections list"},
			Bind{Key: Key{Char: '?'}, Cmd: cmd.HelpPopup, Description: "Help"},
			Bind{Key: Key{Code: tcell.KeyCtrlP}, Cmd: cmd.SearchGlobal, Description: "Global search"},
		},
	},
}
