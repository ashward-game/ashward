package play2earn

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingplay2earn"
	"orbit_nft/contract/service/vestingplay2earn"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingplay2earn.VestingPlay2Earn, error) {
	play2earnAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingplay2earn.NewVestingPlay2Earn(common.HexToAddress(play2earnAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
