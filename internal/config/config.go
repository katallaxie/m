package config

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"

	"github.com/charmbracelet/bubbles/key"
	"github.com/katallaxie/m/pkg/spec"
)

// Flags contains the command line flags.
type Flags struct {
	// File is the configuration file.
	File string
	// Path ...
	Path string
	// Verbose toggles the verbosity.
	Verbose bool
	// Force is forcing creating the file
	Force bool
	// Model is the model.
	Model string
}

type ViewType string

const ()

type Keybinding struct {
	Key     string `yaml:"key"`
	Command string `yaml:"command"`
	Builtin string `yaml:"builtin"`
	Name    string `yaml:"name,omitempty"`
}

func (kb Keybinding) NewBinding(previous *key.Binding) key.Binding {
	helpDesc := ""
	if previous != nil {
		helpDesc = previous.Help().Desc
	}

	if kb.Name != "" {
		helpDesc = kb.Name
	}

	return key.NewBinding(
		key.WithKeys(kb.Key),
		key.WithHelp(kb.Key, helpDesc),
	)
}

type Keybindings struct {
	Universal []Keybinding `yaml:"universal"`
}

// NewFlags returns a new flags.
func NewFlags() *Flags {
	return &Flags{}
}

// Config contains the configuration.
type Config struct {
	// Stdin ...
	Stdin *os.File
	// Stdout ...
	Stdout *os.File
	// Stderr ...
	Stderr *os.File
	// Spec is the configuration specification.
	Spec *spec.Spec
	// Flags ...
	Flags *Flags

	sync.RWMutex `json:"-" yaml:"-"`
}

// New returns a new config.
func New() *Config {
	return &Config{}
}

// Default returns the default configuration.
func Default() *Config {
	return &Config{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Spec:   spec.Default(),
		Flags: &Flags{
			File:    "~/.m.yml",
			Verbose: false,
			Model:   "smollm",
		},
	}
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

// LoadSpec is a helper to load the spec from the config file.
func (c *Config) LoadSpec() error {
	f, err := os.ReadFile(filepath.Clean(c.Flags.File))
	if err != nil {
		return err
	}

	return c.Spec.UnmarshalYAML(f)
}

// TruncateCommand truncates the command to 30 characters.
func TruncateCommand(cmd string) string {
	cmd = strings.ReplaceAll(cmd, "\n", "")

	if len(cmd) > 30 {
		return cmd[:30] + "..."
	}

	return cmd
}
