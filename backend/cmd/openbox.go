package cmd

import (
	"orbit_nft/cmd/openbox"

	"github.com/spf13/cobra"
)

// openboxCmd represents the openbox command
var openboxCmd = &cobra.Command{
	Use:   "openbox",
	Short: "Openbox services",
	Long:  `Openbox services.`,
}

func init() {
	rootCmd.AddCommand(openboxCmd)
	openboxCmd.AddCommand(openbox.ServiceCmd)
}
