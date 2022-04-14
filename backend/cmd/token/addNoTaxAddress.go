package token

import (
	"log"
	"orbit_nft/constant"
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/token"
	"orbit_nft/contract/service/token"
	"orbit_nft/util"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

// AddNoTaxAddressCmd represents the addNoTaxAddress command
var AddNoTaxAddressCmd = &cobra.Command{
	Use:   "addNoTaxAddress",
	Short: "Add no tax address",
	Long:  `Add no tax address.`,
	Run: func(cmd *cobra.Command, args []string) {
		chainId, err := cmd.Flags().GetString("chainId")
		if err != nil {
			log.Fatal(err)
		}
		secretFile, err := cmd.Flags().GetString("secrets")
		if err != nil {
			log.Fatal(err)
		}
		addressFile, err := cmd.Flags().GetString("addressFile")
		if err != nil {
			log.Fatal(err)
		}
		address, err := cmd.Flags().GetString("address")
		if err != nil {
			log.Fatal(err)
		}

		re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
		if !re.MatchString(address) {
			log.Fatal("invalid addresss")
		}
		addr := common.HexToAddress(address)

		if err = addNoTaxAddress(addressFile, secretFile, chainId, addr); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	AddNoTaxAddressCmd.Flags().String("chainId", "97", "Chain Id, default is testnet with chain id 97")
	AddNoTaxAddressCmd.Flags().String("secrets", "../contracts/secrets.json", "Secrets file for contract's owner account")
	AddNoTaxAddressCmd.Flags().String("addressFile", "../common/address.json", "Contracts' address file")
	AddNoTaxAddressCmd.Flags().String("address", "0xHAHA", "Address which will be adding")

	AddNoTaxAddressCmd.MarkFlagRequired("address")
}

func addNoTaxAddress(addressFile, secretFile, chainId string, address common.Address) error {
	mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
	if err != nil {
		return err
	}
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, chainId, secrets)
	if err != nil {
		return err
	}
	tokenAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return err
	}

	sc, err := token.NewToken(common.HexToAddress(tokenAddress), cli.Client())
	if err != nil {
		return err
	}

	_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return sc.AddNoTaxAddress(opts, address)
	})

	return err
}
