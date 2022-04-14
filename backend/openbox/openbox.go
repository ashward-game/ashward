package openbox

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"orbit_nft/contract"
	abi "orbit_nft/contract/abi/openboxgenesis"
	"orbit_nft/contract/service/openboxgenesis"
	"orbit_nft/math"
	"orbit_nft/nft/metadata"
	"orbit_nft/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const BoxEmpty = "none"
const ErrBoxNotSupport = "openbox: box grade is not supported"

type box struct {
	Level   int
	Content map[string]uint
}

type BoxOpener struct {
	client *contract.Client
	metadt *metadata.Metadata
	boxes  map[int]*math.WeightedDistribution
}

func NewBoxOpener(cli *contract.Client, metadt *metadata.Metadata, configFile string) (*BoxOpener, error) {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var records map[string]box
	if err := json.Unmarshal(jsonData, &records); err != nil {
		return nil, err
	}

	boxes := make(map[int]*math.WeightedDistribution)
	for _, box := range records {
		var weightPoints []math.WeightedPoint
		for rarity, ratio := range box.Content {
			weightPoints = append(weightPoints, math.NewWeightedPoint(rarity, ratio))
		}
		boxes[box.Level] = math.NewWeightedDistribution(weightPoints...)
	}

	return &BoxOpener{
		client: cli,
		metadt: metadt,
		boxes:  boxes,
	}, nil
}

func (b *BoxOpener) OpenBox(tokenTypePath string, grade int, hash [32]byte, sRandom, cRandom []byte) error {
	rarityRandom, err := b.randomByGrade(grade, sRandom, cRandom)
	if err != nil {
		return err
	}
	rarity := rarityRandom.(string)
	metadata := BoxEmpty

	isEmpty := strings.EqualFold(rarity, BoxEmpty)

	if !isEmpty {
		rarityPath := filepath.Join(tokenTypePath, rarity)
		recordNft, err := RandomNft(rarityPath)
		if err != nil {
			return err
		}
		metadata, err = b.metadt.GenerateMetadata(rarityPath, recordNft)
		if err != nil {
			return err
		}
	}

	addressOB, err := util.GetContractAddress(b.client.AddressFile(), abi.Name)
	if err != nil {
		return err
	}

	ob, err := openboxgenesis.NewOpenboxgenesis(common.HexToAddress(addressOB), b.client.Client())
	if err != nil {
		return err
	}

	_, err = b.client.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return ob.OpenBox(opts, hash, isEmpty, metadata)
	})

	return err
}

func (b *BoxOpener) randomByGrade(grade int, sRandom, cRandom []byte) (interface{}, error) {
	rand, err := util.XorBytes(sRandom, cRandom)
	if err != nil {
		return nil, err
	}
	source := bytes.NewReader(rand)
	if _, ok := b.boxes[grade]; ok {
		return b.boxes[grade].Sample(source)
	}
	return nil, errors.New(ErrBoxNotSupport)
}
