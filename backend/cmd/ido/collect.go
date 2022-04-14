package ido

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"orbit_nft/contract"
	idoabi "orbit_nft/contract/abi/ido"
	idolog "orbit_nft/contract/event/ido"
	"orbit_nft/util"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

// CollectCmd represents the collect command
var CollectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Collect IDO information",
	Long:  `Collect IDO information.`,
	Run: func(cmd *cobra.Command, args []string) {
		chainId, err := cmd.Flags().GetString("chainId")
		if err != nil {
			log.Fatal(err)
		}
		addressFile, err := cmd.Flags().GetString("address")
		if err != nil {
			log.Fatal(err)
		}
		creationBlock, err := cmd.Flags().GetInt("creationBlock")
		if err != nil {
			log.Fatal(err)
		}
		outFile, err := cmd.Flags().GetString("out")
		if err != nil {
			log.Fatal(err)
		}

		if err = collectInfo(chainId, addressFile, creationBlock, outFile); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	CollectCmd.Flags().String("chainId", "97", "Chain Id, default is testnet with chain id 97")
	CollectCmd.Flags().StringP("address", "a", "../common/address.json", "Contracts' address file")
	CollectCmd.Flags().Int("creationBlock", 16155054, "Contract's creation block number")
	CollectCmd.Flags().String("out", "../assets/vesting/ido/whitelist.csv", "Output file location")
}

func collectInfo(chainId, addressFile string, creationBlock int, outFile string) error {
	buyers := make(map[string]*big.Int)

	fmt.Print("Collecting IDO data starting from block ")
	fmt.Println(creationBlock)

	cli, err := contract.NewBscClient(addressFile, chainId)
	if err != nil {
		return err
	}
	idoAddress, err := util.GetContractAddress(addressFile, idoabi.Name)
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(idoAddress)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(creationBlock)),
		ToBlock:   nil, // latest
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := cli.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(idoabi.ABI)))
	if err != nil {
		return err
	}

	for _, vLog := range logs {
		if vLog.Topics[0].Hex() == crypto.Keccak256Hash([]byte(idolog.LogBuySig)).Hex() {
			var evt idolog.LogBuy
			if len(vLog.Data) > 0 {
				err := contractAbi.UnpackIntoInterface(&evt, idolog.LogBuyName, vLog.Data)
				if err != nil {
					return err
				}
			}

			evt.Buyer = common.HexToAddress(vLog.Topics[1].Hex())
			buyers[evt.Buyer.String()] = evt.Amount
		}
	}

	fmt.Print("There are ")
	fmt.Print(len(buyers))
	fmt.Println(" buyers")

	return writeToCsv(outFile, buyers)
}

func writeToCsv(file string, content map[string]*big.Int) error {
	outFile, err := os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer outFile.Close()
	w := csv.NewWriter(outFile)
	defer w.Flush()

	var data [][]string
	for address, amount := range content {
		row := []string{address, amount.String()}
		data = append(data, row)
	}
	return w.WriteAll(data)
}
