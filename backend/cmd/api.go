package cmd

import (
	"orbit_nft/cmd/api"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run API related commands",
	Long:  `Run API related commands.`,
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.AddCommand(api.InitCmd)
	apiCmd.AddCommand(api.ServeCmd)
	apiCmd.AddCommand(api.StaticCmd)
}
