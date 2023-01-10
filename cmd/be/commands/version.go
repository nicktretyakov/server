package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"be/pkg/version"
)

// versionCmd represents the version command.
//
//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Build time: %s\n", version.BuildTime) //nolint:forbidigo
			fmt.Printf("Git commit: %s\n", version.GitCommit) //nolint:forbidigo
			fmt.Printf("Version: %s\n", version.Version)      //nolint:forbidigo
		},
	})
}
