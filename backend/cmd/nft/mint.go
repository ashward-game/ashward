package nft

import (
	"fmt"
	"orbit_nft/cmd/common"
	"orbit_nft/constant"
	"orbit_nft/contract"
	"orbit_nft/nft"
	"orbit_nft/nft/metadata"
	"orbit_nft/storage/ipfs"
	"orbit_nft/util"
	"os"

	"github.com/spf13/cobra"
)

// MintCmd represents the mint command
var MintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Mint new random NFTs",
	Long:  `Mint new random NFTs in batch with metadata stored in /assets/*.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		num, err := cmd.Flags().GetInt32("number")
		if err != nil {
			panic(err)
		}
		fmt.Println(num)
		mint(path, int(num))
	},
}

const addressFile = "../common/address.json"
const secretFile = "../contracts/secrets.json"

func init() {
	MintCmd.Flags().StringP("path", "p", "../assets/character/", "Path to the metadata directory")
	MintCmd.Flags().Int32P("number", "n", 0, "Number of NFTs shall be minted")
	MintCmd.MarkFlagRequired("path")
}

func mint(path string, num int) {
	// TODO: fix me: take num as param and mint that amount of new tokens
	mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
	if err != nil {
		panic(err)
	}
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	chainId := os.Getenv("CHAIN_ID")
	logger := common.NewLogger()
	defer logger.Sync()
	cli, err := contract.NewBscAuthenticatedClient(
		addressFile, chainId,
		secrets,
		contract.WithLogger(logger),
	)
	if err != nil {
		panic(err)
	}

	ipfs, err := ipfs.NewShell(os.Getenv("IPFS_SHELL_URI"))
	if err != nil {
		panic(err)
	}
	baseURI := os.Getenv("IPFS_URI")
	storage := ipfs

	metadt := metadata.NewMetadata(baseURI, storage)
	mintor, err := nft.NewMintor(cli, metadt)
	if err != nil {
		panic(err)
	}

	err = mintor.Mint(path)
	if err != nil {
		panic(err)
	}
}
