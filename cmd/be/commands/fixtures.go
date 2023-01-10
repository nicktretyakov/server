//go:build dev
// +build dev

package commands

import (
	"github.com/spf13/cobra"

	"be/internal/datastore/testutil"
	"be/internal/lib"
)

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "fixtures",
		Short: "Loads fixtures",
		RunE:  loadFixtures,
	})
}

func loadFixtures(c *cobra.Command, _ []string) error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	return testutil.LoadFixtures(cfg.PostgresDsn, lib.String("internal/datastore/testutil/fixtures"))
}
