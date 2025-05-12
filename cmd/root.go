package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/katallaxie/m/internal/app"
	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/logs"

	pctx "github.com/katallaxie/m/internal/context"

	"github.com/spf13/cobra"
)

var cfg = config.Default()

const appName = "🤖 M"

const (
	versionFmt = "%s (%s %s)"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(VersionCmd)

	RootCmd.PersistentFlags().StringVarP(&cfg.Flags.File, "config", "c", cfg.Flags.File, "configuration file")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")
	RootCmd.PersistentFlags().StringVarP(&cfg.Flags.Model, "model", "m", cfg.Flags.Model, "model to use (default: smollm)")

	RootCmd.SilenceUsage = true
}

var RootCmd = &cobra.Command{
	Use:   "m",
	Short: "m",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context(), cmd, args...)
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(ctx context.Context, _ *cobra.Command, _ ...string) error {
	err := cfg.LoadSpec()
	if err != nil {
		return err
	}

	err = cfg.Spec.Validate()
	if err != nil {
		return err
	}

	_, err = logs.LogToFile("debug.log", "")
	if err != nil {
		return err
	}

	log.Print("debug log file created")

	cfg.Version = version
	cfg.AppName = appName

	pctx := pctx.New(ctx)

	err = app.New(pctx, cfg).Run()
	if err != nil {
		return err
	}

	return nil
}

// isTTY returns whether the passed reader is a TTY or not.
// func isTTY(cmd *cobra.Command) bool {
// 	file, ok := cmd.InOrStdin().(*os.File)
// 	if !ok {
// 		return false
// 	}

// 	return isatty.IsTerminal(file.Fd())
// }
