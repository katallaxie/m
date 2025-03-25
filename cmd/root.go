package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/katallaxie/m/internal/config"

	"github.com/katallaxie/pkg/slices"
	"github.com/katallaxie/prompts"
	"github.com/katallaxie/prompts/ollama"
	"github.com/katallaxie/streams"
	"github.com/katallaxie/streams/sinks"
	"github.com/katallaxie/streams/sources"
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

func mapCompletionMessages(msg prompts.Completion) string {
	f := slices.First(msg.Choices...)
	return f.Message.GetContent()
}

func init() {
	RootCmd.AddCommand(InitCmd)

	RootCmd.PersistentFlags().StringVarP(&cfg.Flags.File, "file", "f", cfg.Flags.File, "configuration file")
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
	err := cfg.LoadSpec()
	if err != nil {
		return err
	}

	cfg.Lock()
	defer cfg.Unlock()

	err = cfg.Spec.Validate()
	if err != nil {
		return err
	}

	api, err := ollama.New(ollama.WithBaseURL("http://localhost:7869"))
	if err != nil {
		return err
	}

	var sb strings.Builder
	for _, arg := range args {
		sb.WriteString(arg)
		sb.WriteString(" ")
	}

	prompt := prompts.Prompt{
		Model: prompts.Model(cfg.Spec.Model),
		Messages: []prompts.Message{
			&prompts.SystemMessage{
				Content: "You are a helpful, but funny AI assistant. You are here to help me with my daily tasks. You add emojies to your answers to make them more fun. You give short answers.",
			},
			&prompts.UserMessage{
				Content: sb.String(),
			},
		},
	}

	res, err := api.Complete(ctx, &prompt)
	if err != nil {
		panic(err)
	}

	source := sources.NewChanSource(res)
	sink := sinks.NewStdout()

	source.Pipe(streams.NewPassThrough()).Pipe(streams.NewMap(mapCompletionMessages)).To(sink)

	return nil
}
