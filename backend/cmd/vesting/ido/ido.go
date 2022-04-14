package ido

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingido"
	"orbit_nft/contract/service/vestingido"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingido.VestingIDO, error) {
	idoAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingido.NewVestingIDO(common.HexToAddress(idoAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
