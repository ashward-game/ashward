package ido

import (
	"errors"
	"fmt"
	"log"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

// AddSubscribersCmd represents the addSubscribers command
var AddSubscribersCmd = &cobra.Command{
	Use:   "add-subscribers",
	Short: "Add subscribers from whitelist to IDO contract",
	Long:  `Add subscribers from whitelist to IDO contract.`,
	Run: func(cmd *cobra.Command, args []string) {
		chainId, err := cmd.Flags().GetString("chainId")
		if err != nil {
			log.Fatal(err)
		}
		csvFile, err := cmd.Flags().GetString("input")
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

		if err = addSubscribers(chainId, csvFile, secretFile, addressFile); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	AddSubscribersCmd.Flags().String("chainId", "97", "Chain Id, default is testnet with chain id 97")
	AddSubscribersCmd.Flags().StringP("secrets", "s", "../contracts/secrets.json", "Secrets file for contract's owner account")
	AddSubscribersCmd.Flags().StringP("input", "i", "../assets/ido/whitelist.csv", "Subscribers' address file")
	AddSubscribersCmd.Flags().StringP("address", "a", "../common/address.json", "Contracts' address file")
}

func addSubscribers(chainId, csvFile, secretFile, addressFile string) error {
	var data []common.Address

	// read file csv
	contentCsv, err := util.ReadFileCsv(csvFile)
	if err != nil {
		return err
	}

	for _, cell := range contentCsv {
		if len(cell[0]) > 0 {
			address := common.HexToAddress(cell[0])
			data = append(data, address)
		}
	}

	if len(data) <= 0 {
		return errors.New("no subscribers' address provided")
	}

	fmt.Print("Adding ")
	fmt.Print(len(data))
	fmt.Print(" addresses to IDO whitelist in ")
	times := len(data)/100 + 1
	fmt.Print(times)
	fmt.Println(" batches")

	for i := 0; i < times; i++ {
		fmt.Print("Batch ")
		fmt.Print(i)

		start := i * 100
		end := (i + 1) * 100
		if i == times-1 {
			end = len(data)
		}

		fmt.Print(": adding from ")
		fmt.Print(start + 1)
		fmt.Print(" to ")
		fmt.Println(end)

		batch := data[start:end]
		if i < times-1 && len(batch) != 100 {
			panic("wrong length of batch")
		}
		// if i == times-1 && len(batch) != 88 {
		// 	panic("wrong length of last batch")
		// }

		scIdo, cli, err := connectIDO(chainId, secretFile, addressFile)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return scIdo.AddSubscribers(opts, batch)
		})
		if err != nil {
			return err
		}
	}

	return nil
}
