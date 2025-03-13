package cmd

import (
	"context"
	"fmt"

	"github.com/katallaxie/m/internal/config"

	"github.com/spf13/cobra"
)

var cfg = config.Default()

const (
	versionFmt = "%s (%s %s)"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
}

var RootCmd = &cobra.Command{
	Use:   "m",
	Short: "m",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(ctx context.Context) error {
	return nil
}
