package spec_test

import (
	"testing"

	"github.com/katallaxie/m/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	t.Parallel()

	defaults := spec.Default()
	assert.Equal(t, 1, defaults.Version)
	assert.Equal(t, "dracula", defaults.Theme)
}
