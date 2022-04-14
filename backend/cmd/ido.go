package cmd

import (
	"orbit_nft/cmd/ido"

	"github.com/spf13/cobra"
)

var idoCmd = &cobra.Command{
	Use:   "ido",
	Short: "IDO related commands",
	Long:  `IDO related commands.`,
}

func init() {
	rootCmd.AddCommand(idoCmd)
	idoCmd.AddCommand(ido.CollectCmd)
	idoCmd.AddCommand(ido.PublicSaleCmd)
	idoCmd.AddCommand(ido.CloseSaleCmd)
	idoCmd.AddCommand(ido.AddSubscribersCmd)
}
