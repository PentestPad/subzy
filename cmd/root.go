package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "subzy",
	Short: "Subdomain takeover tool",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	rootCmd.Execute()
}
