package reserve

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingreserve"
	"orbit_nft/contract/service/vestingreserve"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingreserve.VestingReserve, error) {
	reserveAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingreserve.NewVestingReserve(common.HexToAddress(reserveAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
