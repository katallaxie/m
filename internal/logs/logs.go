package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

// LogToFile ...
func LogToFile(path string, prefix string) (*os.File, error) {
	return LogToFileWith(path, prefix, log.Default())
}

// LogOptionsSetter ...
type LogOptionsSetter interface {
	SetOutput(io.Writer)
	SetPrefix(string)
}

// LogToFileWith ...
func LogToFileWith(path string, prefix string, log LogOptionsSetter) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600) //nolint:mnd
	if err != nil {
		return nil, fmt.Errorf("error opening file for logging: %w", err)
	}
	log.SetOutput(f)

	// Add a space after the prefix if a prefix is being specified and it
	// doesn't already have a trailing space.
	if len(prefix) > 0 {
		finalChar := prefix[len(prefix)-1]
		if !unicode.IsSpace(rune(finalChar)) {
			prefix += " "
		}
	}
	log.SetPrefix(prefix)

	return f, nil
}
