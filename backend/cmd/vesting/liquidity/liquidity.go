package liquidity

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingliquidity"
	"orbit_nft/contract/service/vestingliquidity"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingliquidity.VestingLiquidity, error) {
	liquidityAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingliquidity.NewVestingLiquidity(common.HexToAddress(liquidityAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
