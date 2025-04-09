package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/katallaxie/pkg/filex"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var validate = validator.New()

const (
	// DefaultDirectory is the default directory for the configuration file.
	DefaultDirectory = "m"
	// DefaultPath is the default path for the configuration file.
	DefaultPath = ".m"
	// DefaultFilename is the default filename for the configuration file.
	DefaultFilename = ".m.yml"
)

// Api is the configuration for the API.
type Api struct {
	// Model is the model to use.
	Model string `yaml:"model" validate:"required"`
	// Provider is the provider for the model.
	Provider string `yaml:"provider" validate:"required"`
	// URL is the URL for the API.
	URL string `yaml:"url"`
}

// Spec is the configuration specification for `m`.
type Spec struct {
	// Version is the version of the configuration file.
	Version int `yaml:"version" validate:"required,eq=1"`
	// Api is the API configuration.
	Api Api `yaml:"api" validate:"required"`

	sync.Mutex `yaml:"-"`
}

// UnmarshalYAML unmarshals the configuration file.
func (s *Spec) UnmarshalYAML(data []byte) error {
	spec := struct {
		Version int `yaml:"version" validate:"required,eq=1"`
		Api     struct {
			Model    string `yaml:"model" validate:"required"`
			Provider string `yaml:"provider" validate:"required"`
			URL      string `yaml:"url"`
		} `yaml:"api" validate:"required"`
	}{}

	if err := yaml.Unmarshal(data, &spec); err != nil {
		return errors.WithStack(err)
	}

	s.Version = spec.Version
	s.Api = Api{spec.Api.Model, spec.Api.Provider, spec.Api.URL}

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
