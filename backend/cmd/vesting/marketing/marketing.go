package marketing

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingmarketing"
	"orbit_nft/contract/service/vestingmarketing"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingmarketing.VestingMarketing, error) {
	marketingAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingmarketing.NewVestingMarketing(common.HexToAddress(marketingAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
