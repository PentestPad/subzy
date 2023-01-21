package cmd

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
)

//go:embed version.txt
var version string

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print subzy version",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("subzy version: %s",
			version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
