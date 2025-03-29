package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version/build info",
	Long:  "Print version/build information",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runVersion(cmd.Context())
	},
}

func runVersion(_ context.Context) error {
	const fmat = "%-20s %s\n"

	log.Printf(fmat, "Version", version)

	return nil
}
