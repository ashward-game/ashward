package ido

import (
	"math/big"
	"orbit_nft/util"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestWriteToCsv(t *testing.T) {
	outFile := "./whitelist.csv"
	defer os.Remove(outFile)

	address := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	amount := util.ToWei(6666.67, 18)
	content := make(map[string]*big.Int)
	content[address.String()] = amount

	writeToCsv(outFile, content)
	actual, err := util.ReadFileCsv(outFile)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, address.String(), actual[0][0])
	assert.Equal(t, content[address.String()].String(), actual[0][1])

	address2 := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429f")
	amount2 := util.ToWei(6666.67, 18)
	content[address2.String()] = amount2

	assert.Equal(t, 2, len(content))
	writeToCsv(outFile, content)
	actual, err = util.ReadFileCsv(outFile)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(actual))
}
