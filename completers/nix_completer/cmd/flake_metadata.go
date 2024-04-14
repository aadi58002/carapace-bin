package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_metadataCmd = &cobra.Command{
	Use:     "metadata [flags] [flake-url]",
	Short:   "show flake metadata",
	Run:     func(cmd *cobra.Command, args []string) {},
	Aliases: []string{"info"},
}

func init() {
	carapace.Gen(flake_metadataCmd).Standalone()

	flake_metadataCmd.Flags().Bool("json", false, "Produce output in JSON format")

	addEvaluationFlags(flake_metadataCmd)
	addFlakeFlags(flake_metadataCmd)
	addLoggingFlags(flake_metadataCmd)

	carapace.Gen(flake_metadataCmd).PositionalCompletion(nix.ActionFlakes())

	flakeCmd.AddCommand(flake_metadataCmd)
}
