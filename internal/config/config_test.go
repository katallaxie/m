package config_test

import (
	"testing"

	"github.com/katallaxie/m/internal/config"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	cfg := config.New()
	require.NotNil(t, cfg)
}
