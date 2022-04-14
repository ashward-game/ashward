package vesting

import (
	"bytes"
	"math/big"
	"orbit_nft/contract/abi/token"
	"orbit_nft/contract/abi/vestingido"
	"orbit_nft/contract/rpc"
	"orbit_nft/testutil"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const secretFile = "../../../contracts/secrets.json"
const csvWhitelist = "../../../assets/mocks/vestingido/whitelist.csv"

func deployContract(t *testing.T) (string, func()) {
	address := make(map[string]string)
	deployer := testutil.NewDeployer(t, os.Getenv("CHAIN_ID"), secretFile)
	addToken := deployer.DeployContract(t, token.Name, "Token", "TK")
	addVesting := deployer.DeployContract(t, vestingido.Name, common.HexToAddress(addToken))
	address[token.Name] = addToken
	address[vestingido.Name] = addVesting

	return testutil.WriteAddressToFile(t, address)
}

func TestAddSubscriber(t *testing.T) {
	err := godotenv.Load("../../.test-env")
	assert.NoError(t, err)
	rpc.Initialize("../../config/rpc.json")

	addressFile, teardown := deployContract(t)
	defer teardown()

	chainId := os.Getenv("CHAIN_ID")

	_getWhitelistFile = func(pool string) string {
		return csvWhitelist
	}

	err = addSubscribers(chainId, "ido", secretFile, addressFile)
	assert.NoError(t, err)
}

func TestReadBigInt(t *testing.T) {
	// not use SetBytes if input is number string
	numStr := "100000000000000000000000000000"

	bi, ok := new(big.Int).SetString(numStr, 10)
	assert.True(t, ok)
	assert.Equal(t, bi.String(), numStr)
	bytes.EqualFold(bi.Bytes(), []byte(numStr))
}
