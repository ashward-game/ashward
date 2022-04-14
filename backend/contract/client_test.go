package contract_test

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"orbit_nft/constant"
	"orbit_nft/contract"
	"orbit_nft/contract/abi/nft"
	_ "orbit_nft/contract/event/nft"
	serviceNft "orbit_nft/contract/service/nft"
	"orbit_nft/testutil"
	"orbit_nft/util"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// custom test for event parsers, including:
// - topics
// - data
// test data is given in /orbit_nft/contract/mocks/EmitEvent.sol
var abiSample = `[
    {
      "inputs": [],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "bool",
          "name": "value",
          "type": "bool"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "data",
          "type": "bytes"
        }
      ],
      "name": "TestData",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "hash",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "id",
          "type": "uint256"
        }
      ],
      "name": "TestIndexed",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "sampleHash",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "test",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`
var eventIndexedName = "TestIndexed"
var eventIndexedSample = `{"address":"0xe1e4a51a41238083aa5f3cefd2c426b6ba0f593c","topics":["0x5bfd90c95f768d457b450c44808140795270deeee943f8051d8ffb3a8620765a","0x0000000000000000000000000000000000000000000000010203040506070809","0x7f5a84a02eaadf1082893324d9a1ffe8ec75b0cdc4b4b949b49682ec13e209e3","0x00000000000000000000000000000000000000000000000000000000499602d2"],"data":"0x","blockNumber":"0x12","transactionHash":"0xcecb7bb452f065e87e14f2ea2aae2df2a00e3b1f89ceb7e14558d21bd6671a74","transactionIndex":"0x0","blockHash":"0x1dfc243d81eed718e1b745c639f8ce41814e3cf943c89f6af17a7c17d9362d06","logIndex":"0x0","removed":false}
`
var eventDataName = "TestData"
var eventDataSample = `{"address":"0xe1e4a51a41238083aa5f3cefd2c426b6ba0f593c","topics":["0x41ed0517b75efd28b616010c0830f5d6ac1929024ded9d4979b26c84501e20cf"],"data":"0x000000000000000000000000000000000000000000000001020304050607080900000000000000000000000000000000000000000000000000000000499602d20000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000207f5a84a02eaadf1082893324d9a1ffe8ec75b0cdc4b4b949b49682ec13e209e3","blockNumber":"0x12","transactionHash":"0xcecb7bb452f065e87e14f2ea2aae2df2a00e3b1f89ceb7e14558d21bd6671a74","transactionIndex":"0x0","blockHash":"0x1dfc243d81eed718e1b745c639f8ce41814e3cf943c89f6af17a7c17d9362d06","logIndex":"0x1","removed":false}
`

var eventUint8Name = "TestUint8"
var eventUint8Sample = `{"address":"0x4749c387982DbE433D1423254C85bBB02f415d43","topics":["0x42f1b4a3bcdb8a0f614ceeed9730ff7bdfc722eb24ff9ad30201fbb700106c71","0x000000000000000000000000000000000000000000000000000000000000000d"],"data":"0x","blockNumber":"0x12","transactionHash":"0xcecb7bb452f065e87e14f2ea2aae2df2a00e3b1f89ceb7e14558d21bd6671a74","transactionIndex":"0x0","blockHash":"0x1dfc243d81eed718e1b745c639f8ce41814e3cf943c89f6af17a7c17d9362d06","logIndex":"0x1","removed":false}`

type eventIndexed struct {
	User common.Address
	Hash common.Hash
	Id   *big.Int
}

type eventData struct {
	User   common.Address
	Amount *big.Int
	Value  bool
	Data   []byte
}

type eventUint8 struct {
	Value int
}

func TestEventIndexedParser(t *testing.T) {
	_, err := abi.JSON(strings.NewReader(string(abiSample)))
	assert.NoError(t, err)

	var vLog types.Log
	err = json.Unmarshal([]byte(eventIndexedSample), &vLog)
	assert.NoError(t, err)
	assert.Equal(t, len(vLog.Data), 0)

	var evt eventIndexed

	evt.User = common.HexToAddress(vLog.Topics[1].Hex())
	evt.Hash = common.HexToHash(vLog.Topics[2].Hex())
	evt.Id = util.HexToBigInt(vLog.Topics[3].Hex())

	address, _ := hex.DecodeString("0000000000000000000000010203040506070809")
	assert.Equal(t, evt.User, common.BytesToAddress(address))

	hash := crypto.Keccak256Hash([]byte("HASH"))
	assert.Equal(t, evt.Hash, hash)

	id := new(big.Int)
	id.SetString("1234567890", 10)
	assert.Equal(t, evt.Id, id)
}

func TestEventDataParser(t *testing.T) {
	contractAbi, err := abi.JSON(strings.NewReader(string(abiSample)))
	assert.NoError(t, err)

	var vLog types.Log
	err = json.Unmarshal([]byte(eventDataSample), &vLog)
	assert.NoError(t, err)

	var evt eventData
	err = contractAbi.UnpackIntoInterface(&evt, eventDataName, vLog.Data)
	assert.NoError(t, err)

	user, _ := hex.DecodeString("0000000000000000000000010203040506070809")
	assert.Equal(t, evt.User, common.BytesToAddress(user))

	amount := new(big.Int)
	amount.SetString("1234567890", 10)
	assert.Equal(t, evt.Amount, amount)

	value := true
	assert.Equal(t, evt.Value, value)

	data := crypto.Keccak256Hash([]byte("HASH"))
	assert.Equal(t, evt.Data, data[:])
}

func TestEventUint8Parser(t *testing.T) {
	_, err := abi.JSON(strings.NewReader(string(abiSample)))
	assert.NoError(t, err)

	var vLog types.Log
	err = json.Unmarshal([]byte(eventUint8Sample), &vLog)
	assert.NoError(t, err)

	var evt eventUint8

	evt.Value = int(util.HexToBigInt(vLog.Topics[1].Hex()).Int64())
	assert.Equal(t, evt.Value, 13)
}

func _TestSubscribeNFTEvent(t *testing.T) {
	// setting up
	err := godotenv.Load("../.test-env")
	assert.NoError(t, err)

	const addressFile = "../../common/address.json"
	const secretFile = "../../contracts/secrets.json"
	chainId := os.Getenv("CHAIN_ID")
	sqlDB, teardown := testutil.NewMockDB()
	defer teardown()

	logger, _ := zap.NewDevelopment()

	mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
	assert.NoError(t, err)

	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(
		addressFile, chainId, secrets,
		contract.WithLogger(logger),
		contract.WithDB(sqlDB),
	)
	assert.NoError(t, err)

	// TODO: how to run it automatically?
	// mockClientSubscribe(t, cli, nft.Name)

	addressNFT, err := util.GetContractAddress(cli.AddressFile(), nft.Name)
	assert.NoError(t, err)
	nft, err := serviceNft.NewNFToken(common.HexToAddress(addressNFT), cli.Client())
	assert.NoError(t, err)
	_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return nft.Mint(opts, "/test/path/to/metadata")
	})
	assert.NoError(t, err)
}
