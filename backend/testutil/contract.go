package testutil

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"orbit_nft/constant"
	"orbit_nft/contract/abi/marketplace"
	"orbit_nft/contract/abi/nft"
	"orbit_nft/contract/abi/openboxgenesis"
	"orbit_nft/contract/abi/stakingrewards"
	"orbit_nft/contract/abi/token"
	"orbit_nft/contract/abi/vestingido"
	"orbit_nft/contract/rpc"
	contractdata "orbit_nft/testutil/contract"
	"orbit_nft/util"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/stretchr/testify/assert"
)

type deployer struct {
	fromAddress common.Address
	auth        *bind.TransactOpts
	backend     bind.ContractBackend
}

func NewDeployer(t *testing.T, chainId string, secretFile string) *deployer {
	client, err := rpc.Dial(rpc.BSC, chainId)
	assert.NoError(t, err)

	chainID, err := client.ChainID(context.Background())
	assert.NoError(t, err)

	mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
	assert.NoError(t, err)
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	assert.NoError(t, err)
	derivationPath := hdwallet.MustParseDerivationPath(constant.EthDerivationPath)
	account, err := wallet.Derive(derivationPath, true)
	assert.NoError(t, err)

	privateKey, err := wallet.PrivateKey(account)
	assert.NoError(t, err)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	assert.NoError(t, err)
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	return &deployer{auth: auth, backend: client, fromAddress: fromAddress}

}

func (d *deployer) DeployContract(t *testing.T, name string, params ...interface{}) string {
	// update nonce
	nonce, err := d.backend.PendingNonceAt(context.Background(), d.fromAddress)
	assert.NoError(t, err)
	d.auth.Nonce = big.NewInt(int64(nonce))

	var contractABI string
	var contractBytecode string
	switch name {
	case token.Name:
		contractABI = contractdata.TokenMetaData.ABI
		contractBytecode = contractdata.TokenMetaData.Bin
	case marketplace.Name:
		contractABI = contractdata.MarketplaceMetaData.ABI
		contractBytecode = contractdata.MarketplaceMetaData.Bin
	case stakingrewards.Name:
		contractABI = contractdata.StakingRewardsMetaData.ABI
		contractBytecode = contractdata.StakingRewardsMetaData.Bin
	case nft.Name:
		contractABI = contractdata.NFTMetaData.ABI
		contractBytecode = contractdata.NFTMetaData.Bin
	case openboxgenesis.Name:
		contractABI = contractdata.OpenboxGenesisMetaData.ABI
		contractBytecode = contractdata.OpenboxGenesisMetaData.Bin
	case vestingido.Name:
		contractABI = contractdata.VestingIDOMetaData.ABI
		contractBytecode = contractdata.VestingIDOMetaData.Bin
	default:
		t.Fatal("contract is not supported")
	}

	parsed, err := abi.JSON(strings.NewReader(contractABI))
	assert.NoError(t, err)

	address, _, _, err := bind.DeployContract(d.auth, parsed, common.FromHex(contractBytecode), d.backend, params...)
	assert.NoError(t, err)

	return address.Hex()
}

func WriteAddressToFile(t *testing.T, address map[string]string) (string, func()) {
	dir, err := ioutil.TempDir("", "orbit-test")
	assert.NoError(t, err)
	data, err := json.Marshal(address)
	assert.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(dir, "address.json"), data, 0644)
	assert.NoError(t, err)

	return filepath.Join(dir, "address.json"), func() {
		os.RemoveAll(dir)
	}
}
