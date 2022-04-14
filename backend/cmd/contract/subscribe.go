package contract

import (
	"context"
	"orbit_nft/cmd/common"
	"orbit_nft/contract"
	orbitContext "orbit_nft/contract/context"
	"orbit_nft/logger"
	"orbit_nft/util"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

const addressFile = "../common/address.json"

var (
	specifiedSubscribeContracts []string
)

// SubscribeCmd represents the subscribe command
var SubscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to all Orbit's deployed contracts and update the database in real-time",
	Long: `Subscribe to all Orbit's deployed contracts.
	Each log record will be parsed and used to update our database, that tracks all transactions related to Orbit Metaverse.`,
	Run: func(cmd *cobra.Command, args []string) {
		subscribe()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	SubscribeCmd.Flags().StringSliceVarP(&specifiedSubscribeContracts, "contract", "c", []string{}, "List specified contract name, if empty, will subscribe all register contract")
}

func subscribe() {
	chainId := os.Getenv("CHAIN_ID")
	openboxGenesisRPCAdrr := os.Getenv("OPENBOX_SERVICE_ADDRESS")
	sqlDB := common.MySQL()

	mLogger := common.NewLogger()
	defer mLogger.Sync()
	defer logger.Shutdown()

	cli, err := contract.NewBscClient(
		addressFile, chainId,
		contract.WithLogger(mLogger),
		contract.WithDB(sqlDB),
	)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx = orbitContext.WithOpenboxRPCAddress(ctx, openboxGenesisRPCAdrr)
	cli.SubscribeAll(ctx, specifiedSubscribeContracts...)

	util.WaitOSSignalHandler(func() {
		cancel()
	}, syscall.SIGINT, syscall.SIGTERM)
}
