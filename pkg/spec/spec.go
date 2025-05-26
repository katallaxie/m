package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/katallaxie/pkg/filex"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var validate = validator.New()

const (
	// DefaultPath is the default path for the configuration file.
	DefaultDirectory = ".m"
	// DefaultFilename is the default filename for the configuration file.
	DefaultFilename = ".m.yml"
)

// Provider is the provider configuration for the API.
type Provider struct {
	// Model is the model to use.
	Model string `yaml:"model"`
	// API is the API provider.
	API string `yaml:"api"`
	// URL is the URL for the API.
	URL string `yaml:"url"`
	// Key is the key for the API.
	Key string `yaml:"key"`
}

// Spec is the configuration specification for `m`.
type Spec struct {
	// Version is the version of the configuration file.
	Version int `yaml:"version" validate:"required,eq=1"`
	// Provider is the provider configuration for the API.
	Provider *Provider `yaml:"provider" validate:"required"`
}

// UnmarshalYAML unmarshals the configuration file.
func (s *Spec) UnmarshalYAML(data []byte) error {
	spec := struct {
		Version  int `yaml:"version" validate:"required,eq=1"`
		Provider struct {
			Model string `yaml:"model"`
			API   string `yaml:"api"`
			URL   string `yaml:"url"`
			Key   string `yaml:"key"`
		} `yaml:"api" validate:"required"`
	}{}

	if err := yaml.Unmarshal(data, &spec); err != nil {
		return errors.WithStack(err)
	}

	s.Version = spec.Version
	s.Provider = &Provider{
		Model: spec.Provider.Model,
		API:   spec.Provider.API,
		URL:   spec.Provider.URL,
		Key:   spec.Provider.Key,
	}

	return nil
}

// Default returns the default configuration.
func Default() *Spec {
	return &Spec{
		Version: 1,
	}
}

// Validate validates the configuration.
func (s *Spec) Validate() error {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("yaml"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return validate.Struct(s)
}

// Write is the write function for the spec.
func Write(s *Spec, file string, force bool) error {
	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	ok, _ := filex.FileExists(filepath.Clean(file))
	if ok && !force {
		return fmt.Errorf("%s already exists, use --force to overwrite", file)
	}

	f, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
