package team

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingteam"
	"orbit_nft/contract/service/vestingteam"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingteam.VestingTeam, error) {
	teamAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingteam.NewVestingTeam(common.HexToAddress(teamAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
