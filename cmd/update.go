package cmd

import (
	"github.com/LukaSikic/subzy/runner"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update local fingerprints.json file",
	Aliases: []string{"u"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runner.CheckFingerprints()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
