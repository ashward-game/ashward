package nft

import (
	"errors"
	"orbit_nft/contract"
	"orbit_nft/contract/abi/nft"
	serviceNft "orbit_nft/contract/service/nft"
	"orbit_nft/nft/metadata"
	"orbit_nft/util"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const ErrNoDataFile = "nft: no data file for minting"

type NftMint struct {
	client     *contract.Client
	addressNFT string
	metadata   *metadata.Metadata
}

func NewMintor(cli *contract.Client, metadata *metadata.Metadata) (*NftMint, error) {
	addressNFT, err := util.GetContractAddress(cli.AddressFile(), nft.Name)
	if err != nil {
		return nil, err
	}
	return &NftMint{
		client:     cli,
		addressNFT: addressNFT,
		metadata:   metadata,
	}, nil
}

// Mint mints new tokens of type `baseDir` using metadata stored in
// `baseDir` directory, with the following structure:
//
// - baseDir/images: image of tokens
//
// - baseDir/data.csv: metadata stored in csv format
//
// - baseDir/template.stub: template for metadata in JSON using OpenSea standard
func (m *NftMint) Mint(baseDir string) error {
	csvFiles, err := util.GetFilesRecursively(baseDir, "data.csv")
	if err != nil {
		return err
	}

	if len(csvFiles) == 0 {
		return errors.New(ErrNoDataFile)
	}

	client := m.client.Client()
	addressContract := common.HexToAddress(m.addressNFT)
	nft, err := serviceNft.NewNFToken(addressContract, client)
	if err != nil {
		return err
	}

	for _, csvFile := range csvFiles {
		rarityPath := filepath.Dir(csvFile)

		contentCsv, err := util.ReadFileCsv(csvFile)
		if err != nil {
			return err
		}
		for index, record := range contentCsv {
			if index == 0 {
				// skip first line(header)
				continue
			}
			metadataCid, err := m.metadata.GenerateMetadata(rarityPath, record)
			if err != nil {
				return err
			}

			_, err = m.client.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
				return nft.Mint(opts, metadataCid)
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
