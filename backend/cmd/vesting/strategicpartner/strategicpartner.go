package strategicpartner

import (
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/vestingstrategicpartner"
	"orbit_nft/contract/service/vestingstrategicpartner"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func Connect(addressFile string, cli *contract.Client) (*vestingstrategicpartner.VestingStrategicPartner, error) {
	strategicpartnerAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, err
	}
	sc, err := vestingstrategicpartner.NewVestingStrategicPartner(common.HexToAddress(strategicpartnerAddress), cli.Client())
	if err != nil {
		return nil, err
	}

	return sc, nil
}
