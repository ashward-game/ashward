package openbox

import (
	"crypto/rand"
	"errors"
	"math/big"
	"orbit_nft/util"
	"path/filepath"
)

const ErrEmptyFile = "openbox: data.csv is empty"

func RandomNft(rarityPath string) ([]string, error) {
	csvFile := filepath.Join(rarityPath, "data.csv")
	contentCsv, err := util.ReadFileCsv(csvFile)
	if err != nil {
		return nil, err
	}
	data := contentCsv[1:]

	max := len(data)
	if max <= 0 {
		return nil, errors.New(ErrEmptyFile)
	}
	ran, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return nil, err
	}
	return data[int(ran.Int64())], nil
}
