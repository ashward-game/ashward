package ido

import (
	"orbit_nft/constant"
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/ido"
	"orbit_nft/contract/service/ido"
	"orbit_nft/util"

	"github.com/ethereum/go-ethereum/common"
)

func connectIDO(chainId, secretFile, addressFile string) (*ido.Ido, *contract.Client, error) {
	mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
	if err != nil {
		return nil, nil, err
	}
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, chainId, secrets)
	if err != nil {
		return nil, nil, err
	}
	idoAddress, err := util.GetContractAddress(addressFile, abi.Name)
	if err != nil {
		return nil, nil, err
	}
	scIdo, err := ido.NewIdo(common.HexToAddress(idoAddress), cli.Client())
	if err != nil {
		return nil, nil, err
	}

	return scIdo, cli, nil
}
