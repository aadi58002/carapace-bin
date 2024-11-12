package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-install",
	Short: "package manager utility",
	Long:  "https://man.voidlinux.org/xbps-install.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("automatic","A", false, "Enables automatic installation mode, i.e. package will be treated as orphan if no package is depending on it directly.")
	rootCmd.Flags().StringP("config","C","/etc/xbps.d", "Specifies a path to the XBPS configuration directory.")
	rootCmd.Flags().BoolP("download-only","D", false, "Only download packages to the cache, do not do any other installation steps.")
	rootCmd.Flags().BoolP("debug","d", false, "Enables extra debugging shown to stderr.")
	rootCmd.Flags().BoolP("force","f", false, "Force downgrade installation or reinstallation")
	rootCmd.Flags().BoolP("help", "h", false, "Show the help message.")
	rootCmd.Flags().BoolP("ignore-conf-repos", "i", false, "Ignore repositories defined in configuration files. Only repositories specified in the command line via --repository will be used.")
	rootCmd.Flags().BoolP("ignore-file-conflicts", "I", false, "Ignore detected file conflicts in a transaction.")
	rootCmd.Flags().BoolP("version", "V", false, "Display version")


	// Completion handling
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
	})

	// Positional completion
	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
