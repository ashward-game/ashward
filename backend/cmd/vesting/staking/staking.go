package staking

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingstaking"
	"orbit_nft/contract/service/vestingstaking"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingstaking.VestingStaking, error) {
	stakingAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingstaking.NewVestingStaking(common.HexToAddress(stakingAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
