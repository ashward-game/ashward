package cmd

import (
	"orbit_nft/cmd/contract"

	"github.com/spf13/cobra"
)

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "Run contract related commands",
	Long:  `Run contract related commands.`,
}

func init() {
	rootCmd.AddCommand(contractCmd)
	contractCmd.AddCommand(contract.SubscribeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
