package cmd

import (
	"context"
	"fmt"

	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/ui"
	pctx "github.com/katallaxie/m/internal/ui/context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var cfg = config.Default()

const (
	defaultLogFile = "m.log"
	versionFmt     = "%s (%s %s)"
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

	// RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
}

var RootCmd = &cobra.Command{
	Use:   "m",
	Short: "m",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context(), args...)
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(ctx context.Context, args ...string) error {
	// see https://github.com/charmbracelet/lipgloss/issues/73
	lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())

	f, err := tea.LogToFile(defaultLogFile, "")
	if err != nil {
		return err
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)

	err = cfg.LoadSpec()
	if err != nil {
		return err
	}

	cfg.Lock()
	defer cfg.Unlock()

	err = cfg.Spec.Validate()
	if err != nil {
		return err
	}

	c := pctx.WithContext(ctx)

	p := tea.NewProgram(
		ui.New(c),
		// enable mouse motion will make text not able to select
		// tea.WithMouseCellMotion(),
		tea.WithAltScreen(),
		tea.WithReportFocus(),
	)

	// first the program needs to be initialized
	c.SetProgram(p)

	_, err = p.Run()
	if err != nil {
		log.Fatal("failed to run program", err)
	}

	return nil
}
