package cmd

import (
	"context"
	"log"
	"os"

	"github.com/katallaxie/m/pkg/spec"
	"github.com/katallaxie/pkg/filex"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new config",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit(cmd.Context())
	},
}

func runInit(_ context.Context) error {
	log.Printf("initializing config (%s)", cfg.Flags.File)

	if err := spec.Write(spec.Default(), cfg.Flags.File, cfg.Flags.Force); err != nil {
		return err
	}

	log.Printf("creating config folder (%s)", cfg.Flags.Path)

	err := filex.MkdirAll(cfg.Flags.Path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
