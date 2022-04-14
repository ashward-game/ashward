package private

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingprivate"
	"orbit_nft/contract/service/vestingprivate"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingprivate.VestingPrivate, error) {
	privateAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingprivate.NewVestingPrivate(common.HexToAddress(privateAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
