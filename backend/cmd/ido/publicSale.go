package ido

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

// PublicSaleCmd represents the publicSale command
var PublicSaleCmd = &cobra.Command{
	Use:   "publicSale",
	Short: "Open IDO public sale",
	Long:  `Open IDO public sale.`,
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

		if err = publicSale(chainId, secretFile, addressFile); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	PublicSaleCmd.Flags().String("chainId", "97", "Chain Id, default is testnet with chain id 97")
	PublicSaleCmd.Flags().StringP("secrets", "s", "../contracts/secrets.json", "Secrets file for contract's owner account")
	PublicSaleCmd.Flags().StringP("address", "a", "../common/address.json", "Contracts' address file")
}

func publicSale(chainId, secretFile, addressFile string) error {
	ido, cli, err := connectIDO(chainId, secretFile, addressFile)
	if err != nil {
		return err
	}

	_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return ido.PublicSale(opts)
	})
	return err
}
