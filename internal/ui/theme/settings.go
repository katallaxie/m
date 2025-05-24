package theme

import "sync"

const defaultThemeName = "dracula"

// Settings defines the interface for theme settings in the application.
type Settings struct {
	themes  map[string]Theme
	current string
	sync.RWMutex
}

var settings = &Settings{
	themes:  make(map[string]Theme),
	current: defaultThemeName,
}

// Register adds a new theme to the settings.
func RegisterTheme(name string, theme Theme) {
	settings.Lock()
	defer settings.Unlock()

	settings.themes[name] = theme
}

// Set sets the current theme by name.
func Set(name string) {
	settings.Lock()
	defer settings.Unlock()

	if _, exists := settings.themes[name]; exists {
		settings.current = name
	}
}

// Current returns the currently set theme.
func Current() Theme {
	settings.RLock()
	defer settings.RUnlock()

	theme, exists := settings.themes[settings.current]
	if !exists {
		return settings.themes[defaultThemeName] // Fallback to default theme
	}

	return theme
}

// Available returns a list of all registered theme names.
func Available() []string {
	settings.RLock()
	defer settings.RUnlock()

	names := make([]string, 0, len(settings.themes))
	for name := range settings.themes {
		names = append(names, name)
	}

	return names
}
