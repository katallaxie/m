package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/katallaxie/m/internal/config"
	"github.com/katallaxie/m/pkg/chats"
	"github.com/katallaxie/m/pkg/models/ollama"

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
	RootCmd.PersistentFlags().StringVarP(&cfg.Flags.Model, "model", "m", cfg.Flags.Model, "model to use (default: smollm)")

	RootCmd.SilenceErrors = true
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
	api, err := ollama.New(ollama.WithBaseURL("http://localhost:7869"))
	if err != nil {
		return err
	}

	var sb strings.Builder
	for _, arg := range args {
		sb.WriteString(arg)
		sb.WriteString(" ")
	}

	msgs := []chats.Message{
		{
			Role:    "user",
			Content: sb.String(),
		},
	}

	chat := chats.NewChat(cfg.Flags.Model, msgs, nil, nil, "", "")

	_, err = api.Generate(ctx, chat)
	if err != nil {
		return err
	}

	return nil
}
