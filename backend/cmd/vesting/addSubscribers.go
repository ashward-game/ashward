package vesting

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"orbit_nft/cmd/vesting/advisory"
	"orbit_nft/cmd/vesting/ido"
	"orbit_nft/cmd/vesting/liquidity"
	"orbit_nft/cmd/vesting/marketing"
	"orbit_nft/cmd/vesting/play2earn"
	"orbit_nft/cmd/vesting/private"
	"orbit_nft/cmd/vesting/reserve"
	"orbit_nft/cmd/vesting/staking"
	"orbit_nft/cmd/vesting/strategicpartner"
	"orbit_nft/cmd/vesting/team"
	"orbit_nft/constant"
	"orbit_nft/contract"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

// AddSubscribersCmd represents the addSubscribers command
var AddSubscribersCmd = &cobra.Command{
	Use:   "add-subscribers",
	Short: "Add subscribers from whitelist to vesting contract",
	Long:  `Add subscribers from whitelist to vetsing contract.`,
	Run: func(cmd *cobra.Command, args []string) {
		chainId, err := cmd.Flags().GetString("chainId")
		if err != nil {
			log.Fatal(err)
		}
		pool, err := cmd.Flags().GetString("pool")
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

		if err = addSubscribers(chainId, pool, secretFile, addressFile); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	AddSubscribersCmd.Flags().String("chainId", "97", "Chain Id, default is testnet with chain id 97")
	AddSubscribersCmd.Flags().StringP("secrets", "s", "../contracts/secrets.json", "Secrets file for contract's owner account")
	AddSubscribersCmd.Flags().String("pool", "ido", "Vesting pool's name")
	AddSubscribersCmd.Flags().StringP("address", "a", "../common/address.json", "Contracts' address file")

	AddSubscribersCmd.MarkFlagRequired("pool")
}

var _getWhitelistFile = func(pool string) string {
	return "../assets/vesting/" + pool + "/whitelist.csv"
}

func addSubscribers(chainId, pool, secretFile, addressFile string) error {
	fmt.Println("Adding subscribers for pool " + pool)
	var addresses []common.Address
	var amounts []*big.Int

	// read file csv
	contentCsv, err := util.ReadFileCsv(_getWhitelistFile(pool))
	if err != nil {
		return err
	}

	for _, cell := range contentCsv {
		if len(cell[0]) > 0 {
			address := common.HexToAddress(cell[0])
			addresses = append(addresses, address)
		}
		if len(cell[1]) > 0 {
			amount, ok := new(big.Int).SetString(cell[1], 10)
			if !ok {
				return errors.New("amount is invalid")
			}
			amounts = append(amounts, amount)
		}
	}

	if len(addresses) <= 0 {
		return errors.New("no subscribers' address provided")
	}
	if len(addresses) != len(amounts) {
		return errors.New("beneficiaries and amounts' length should be equal")
	}

	if len(addresses) <= 100 {
		return singleCall(addressFile, secretFile, chainId, pool, addresses, amounts)
	}

	fmt.Print("Adding ")
	fmt.Print(len(addresses))
	fmt.Print(" addresses ")
	times := len(addresses)/100 + 1
	fmt.Print(times)
	fmt.Println(" batches")

	for i := 0; i < times; i++ {
		fmt.Print("Batch ")
		fmt.Print(i)

		start := i * 100
		end := (i + 1) * 100
		if i == times-1 {
			end = len(addresses)
		}

		fmt.Print(": adding from ")
		fmt.Print(start + 1)
		fmt.Print(" to ")
		fmt.Println(end)

		batchAddr := addresses[start:end]
		batchAmount := amounts[start:end]
		if i < times-1 && len(batchAddr) != 100 {
			panic("wrong length of batch")
		}
		err = singleCall(addressFile, secretFile, chainId, pool, batchAddr, batchAmount)
		if err != nil {
			return err
		}
	}
	return nil
}

func singleCall(addressFile, secretFile, chainId, pool string, addresses []common.Address, amounts []*big.Int) error {
	mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
	if err != nil {
		return err
	}
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, chainId, secrets)
	if err != nil {
		return err
	}

	switch pool {
	case "advisory":
		sc, err := advisory.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "ido":
		sc, err := ido.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "liquidity":
		sc, err := liquidity.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "marketing":
		sc, err := marketing.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "play2earn":
		sc, err := play2earn.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "private":
		sc, err := private.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "reserve":
		sc, err := reserve.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "staking":
		sc, err := staking.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "strategicpartner":
		sc, err := strategicpartner.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err

	case "team":
		sc, err := team.Connect(addressFile, cli)
		if err != nil {
			return err
		}

		_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return sc.AddBeneficiaries(opts, addresses, amounts)
		})
		return err
	default:
		return errors.New("unsupported pool")
	}
}
