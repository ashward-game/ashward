package cmd

import (
	"orbit_nft/cmd/nft"

	"github.com/spf13/cobra"
)

// nftCmd represents the nft command
var nftCmd = &cobra.Command{
	Use:   "nft",
	Short: "Run NFT related commands",
	Long:  `Run NFT related commands.`,
}

func init() {
	rootCmd.AddCommand(nftCmd)
	nftCmd.AddCommand(nft.MintCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
