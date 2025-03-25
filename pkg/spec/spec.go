package spec

import (
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert/yaml"
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

// Spec is the configuration specification for `m`.
type Spec struct {
	// Version is the version of the configuration file.
	Version int `yaml:"version" validate:"required,eq=1"`
	// Model is the model to use.
	Model string `yaml:"model" validate:"required"`

	sync.Mutex `yaml:"-"`
}

// UnmarshalYAML unmarshals the configuration file.
func (s *Spec) UnmarshalYAML(data []byte) error {
	spec := struct {
		Version int    `yaml:"version" validate:"required,eq=1"`
		Model   string `yaml:"model" validate:"required"`
	}{}

	if err := yaml.Unmarshal(data, &spec); err != nil {
		return errors.WithStack(err)
	}

	s.Version = spec.Version
	s.Model = spec.Model

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
