package advisory

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingadvisory"
	"orbit_nft/contract/service/vestingadvisory"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingadvisory.VestingAdvisory, error) {
	advisoryAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingadvisory.NewVestingAdvisory(common.HexToAddress(advisoryAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
