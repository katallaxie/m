package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/katallaxie/m/internal/app"
	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/internal/db"
	"github.com/katallaxie/m/internal/logs"
	"github.com/katallaxie/m/internal/models"
	"github.com/katallaxie/m/internal/ui"
	"github.com/katallaxie/pkg/dbx"
	"github.com/katallaxie/prompts"
	"github.com/katallaxie/prompts/ollama"
	"github.com/katallaxie/prompts/perplexity"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	conn, err := gorm.Open(sqlite.Open("./m.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	store, err := dbx.NewDatabase(conn, db.NewReadTx(), db.NewWriteTx())
	if err != nil {
		return err
	}

	err = store.Migrate(ctx,
		&models.Session{},
	)
	if err != nil {
		return err
	}

	var client prompts.Chat

	switch cfg.Spec.Provider.API {
	case "perplexity":
		client = perplexity.New(perplexity.WithApiKey(cfg.Spec.Provider.Key))
	default:
		client = ollama.New(ollama.WithBaseURL(cfg.Spec.Provider.URL))
	}

	app, err := app.New(ctx, store, client, cfg)
	if err != nil {
		return err
	}

	defer app.Dispose()

	zone.NewGlobal()
	program := tea.NewProgram(
		ui.New(app),
		tea.WithAltScreen(),
	)

	_, err = program.Run()
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
