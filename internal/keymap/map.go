package keymap

import (
	"github.com/gdamore/tcell/v2"

	"github.com/katallaxie/m/internal/cmd"
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
	HomeGroup = "home"
	HelpGroup = "help"
)

var Keymaps = KeymapSystem{
	Groups: map[string]Map{
		HomeGroup: {
			Bind{Key: Key{Code: tcell.KeyF1, Mod: tcell.ModNone}, Cmd: cmd.HelpPopup, Description: "Help"},
			Bind{Key: Key{Code: tcell.KeyF2, Mod: tcell.ModNone}, Cmd: cmd.NewChat, Description: "New"},
			Bind{Key: Key{Code: tcell.KeyCtrlQ, Mod: tcell.ModCtrl}, Cmd: cmd.Quit, Description: "Quit"},
			Bind{Key: Key{Code: tcell.KeyCtrlP, Mod: tcell.ModCtrl}, Cmd: cmd.FocusPrompt, Description: "Focus prompt"},
			Bind{Key: Key{Code: tcell.KeyCtrlC, Mod: tcell.ModCtrl}, Cmd: cmd.FocusChat, Description: "Focus chat"},
		},
		HelpGroup: {
			Bind{Key: Key{Code: tcell.KeyEsc}, Cmd: cmd.Close, Description: "Close"},
		},
		ChatGroup: {
			Bind{Key: Key{Char: 'L'}, Cmd: cmd.MoveRight, Description: "Focus table"},
			Bind{Key: Key{Char: 'H'}, Cmd: cmd.MoveLeft, Description: "Focus tree"},
			Bind{Key: Key{Code: tcell.KeyCtrlE}, Cmd: cmd.SwitchToEditorView, Description: "Open SQL editor"},
			Bind{Key: Key{Code: tcell.KeyCtrlS}, Cmd: cmd.Save, Description: "Execute pending changes"},
			// Bind{Key: Key{Char: 'q'}, Cmd: cmd.Quit, Description: "Quit"},
			Bind{Key: Key{Code: tcell.KeyBackspace2}, Cmd: cmd.SwitchToConnectionsView, Description: "Switch to connections list"},
			Bind{Key: Key{Char: '?'}, Cmd: cmd.HelpPopup, Description: "Help"},
			Bind{Key: Key{Code: tcell.KeyCtrlP}, Cmd: cmd.SearchGlobal, Description: "Global search"},
		},
	},
}

// Map is a collection of keybinds
type Map []Bind

// Resolve translates a tcell.EventKey to a
// command based on the bindings in the map.
//
// If no binding could be found. commands.Noop is returned.
func (m Map) Resolve(event *tcell.EventKey) cmd.Command {
	for _, bind := range m {
		if event.Key() == tcell.KeyRune && event.Modifiers() == bind.Key.Mod {
			if bind.Key.Char == event.Rune() {
				return bind.Cmd
			}
		} else if event.Key() == bind.Key.Code && event.Modifiers() == bind.Key.Mod {
			return bind.Cmd
		}
	}

	return cmd.Noop
}
