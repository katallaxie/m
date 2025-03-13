package config

import (
	"os"
	"os/user"
	"sync"
)

// Flags contains the command line flags.
type Flags struct {
	// Verbose toggles the verbosity.
	Verbose bool
}

// NewFlags returns a new flags.
func NewFlags() *Flags {
	return &Flags{}
}

// Config contains the configuration.
type Config struct {
	// Flags ...
	Flags *Flags

	sync.RWMutex `json:"-" yaml:"-"`
}

// New returns a new config.
func New() *Config {
	return &Config{
		Flags: &Flags{},
	}
}

// Default returns the default configuration.
func Default() *Config {
	return &Config{}
}

// HomeDir returns the home directory.
func (c *Config) HomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, err
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}
