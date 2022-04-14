package cmd

import (
	"orbit_nft/cmd/token"

	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Token related commands",
	Long:  `Token related commands.`,
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	tokenCmd.AddCommand(token.AddSellingAddressCmd)
	tokenCmd.AddCommand(token.AddNoTaxAddressCmd)
	tokenCmd.AddCommand(token.RemoveNoTaxAddressCmd)
}
