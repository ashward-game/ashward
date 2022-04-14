package cmd

import (
	"orbit_nft/cmd/vesting"

	"github.com/spf13/cobra"
)

var vestingCmd = &cobra.Command{
	Use:   "vesting",
	Short: "Vesting related commands",
	Long:  `Vesting related commands.`,
}

func init() {
	rootCmd.AddCommand(vestingCmd)
	vestingCmd.AddCommand(vesting.AddSubscribersCmd)
}
