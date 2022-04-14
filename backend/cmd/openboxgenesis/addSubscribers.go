package openboxgenesis

import (
	"errors"
	"log"
	"orbit_nft/constant"
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/openboxgenesis"
	"orbit_nft/contract/service/openboxgenesis"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

var AddSubscribersCmd = &cobra.Command{
	Use:   "add-subscribers",
	Short: "OpenboxGenesis related commands",
	Long:  `OpenboxGenesis related commands.`,
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
	AddSubscribersCmd.Flags().StringP("input", "i", "../assets/openbox-genesis/whitelist.csv", "Subscribers' address file")
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

	ob, cli, err := connectOpenbox(chainId, secretFile, addressFile)
	if err != nil {
		return err
	}

	_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return ob.AddSubscribers(opts, data)
	})
	return err
}

func connectOpenbox(chainId, secretFile, addressFile string) (*openboxgenesis.Openboxgenesis, *contract.Client, error) {
	mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
	if err != nil {
		return nil, nil, err
	}
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, chainId, secrets)
	if err != nil {
		return nil, nil, err
	}
	obAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, nil, err
	}
	ob, err := openboxgenesis.NewOpenboxgenesis(common.HexToAddress(obAddress), cli.Client())
	if err != nil {
		return nil, nil, err
	}

	return ob, cli, nil
}
