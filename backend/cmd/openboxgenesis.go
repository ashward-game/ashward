package cmd

import (
	"orbit_nft/cmd/openboxgenesis"

	"github.com/spf13/cobra"
)

// openboxgenesisCmd represents the sign command
var openboxgenesisCmd = &cobra.Command{
	Use:   "openboxgenesis",
	Short: "OpenboxGenesis add whitelist to contract",
	Long:  `OpenboxGenesis add whitelist to contract.`,
}

func init() {
	rootCmd.AddCommand(openboxgenesisCmd)
	openboxgenesisCmd.AddCommand(openboxgenesis.AddSubscribersCmd)
}
