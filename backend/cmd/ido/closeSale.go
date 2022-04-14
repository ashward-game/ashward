package ido

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

// CloseSaleCmd represents the closeSale command
var CloseSaleCmd = &cobra.Command{
	Use:   "closeSale",
	Short: "Close IDO sale",
	Long:  `Close IDO sale.`,
	Run: func(cmd *cobra.Command, args []string) {
		chainId, err := cmd.Flags().GetString("chainId")
		if err != nil {
			log.Fatal(err)
		}
		secretFile, err := cmd.Flags().GetString("secrets")
		if err != nil {
			log.Fatal(err)
		}
		addressFile, err := cmd.Flags().GetString("address")
		if err != nil {
			log.Fatal(err)
		}

		if err = stopContract(chainId, secretFile, addressFile); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	CloseSaleCmd.Flags().String("chainId", "97", "Chain Id, default is testnet with chain id 97")
	CloseSaleCmd.Flags().StringP("secrets", "s", "../contracts/secrets.json", "Secrets file for contract's owner account")
	CloseSaleCmd.Flags().StringP("address", "a", "../common/address.json", "Contracts' address file")
}

func stopContract(chainId, secretFile, addressFile string) error {
	ido, cli, err := connectIDO(chainId, secretFile, addressFile)
	if err != nil {
		return err
	}

	_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return ido.Stop(opts)
	})
	return err
}
