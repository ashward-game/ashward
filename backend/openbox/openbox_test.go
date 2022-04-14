package openbox

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"orbit_nft/constant"
	"orbit_nft/contract"
	"orbit_nft/contract/abi/nft"
	"orbit_nft/contract/abi/openboxgenesis"
	"orbit_nft/contract/rpc"
	scnft "orbit_nft/contract/service/nft"
	scob "orbit_nft/contract/service/openboxgenesis"
	cryp "orbit_nft/crypto"
	"orbit_nft/nft/metadata"
	"orbit_nft/storage"
	"orbit_nft/storage/localstorage"
	"orbit_nft/testutil"
	"orbit_nft/util"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/stretchr/testify/assert"
)

const deployerSecret = "../../contracts/secrets.json"
const backendSecret = "../secrets/wallet.json"
const assetsBaseDir = "../../assets"
const boxesConfigFile = "../config/openboxgenesis.json"

const nonSupportedGrade = 3
const supportedGrade = 1
const boxPrice = 1

func deployContracts(t *testing.T, accBackend common.Address) (string, func()) {
	var pub [32]byte
	copy(pub[:], accBackend.Bytes())

	address := make(map[string]string)
	deployer := testutil.NewDeployer(t, os.Getenv("CHAIN_ID"), deployerSecret)
	addNft := deployer.DeployContract(t, nft.Name, "NFT", "NFT", os.Getenv("LOCAL_STORAGE_URI"))
	address[nft.Name] = addNft
	addOB := deployer.DeployContract(t, openboxgenesis.Name, pub, common.HexToAddress(addNft), uint32(1), big.NewInt(1), uint32(1), big.NewInt(1), uint32(1), big.NewInt(1))
	address[openboxgenesis.Name] = addOB
	return testutil.WriteAddressToFile(t, address)
}

func setupBackend(t *testing.T) (*ecdsa.PrivateKey, storage.Storage) {
	var storage storage.Storage

	err := godotenv.Load("../.test-env")
	assert.NoError(t, err)
	rpc.Initialize("../config/rpc.json")

	// account backend
	mnemonic, err := util.GetSecrets(backendSecret, "mnemonic")
	assert.NoError(t, err)
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	assert.NoError(t, err)
	derivationPath := hdwallet.MustParseDerivationPath(constant.EthDerivationPath)
	account, err := wallet.Derive(derivationPath, true)
	assert.NoError(t, err)
	pk, err := wallet.PrivateKey(account)
	assert.NoError(t, err)

	// storage
	storage, err = localstorage.New(filepath.Join(assetsBaseDir, "nft"))
	assert.NoError(t, err)
	return pk, storage
}

func setupSmartContract(t *testing.T, addressFile string, accBackend, buyer common.Address) *scob.Openboxgenesis {
	// setup client & caller (owner)
	mnemonic, err := util.GetSecrets(deployerSecret, "mnemonic")
	assert.NoError(t, err)
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, os.Getenv("CHAIN_ID"), secrets)
	assert.NoError(t, err)

	// setup opener
	addrOB, err := util.GetContractAddress(addressFile, openboxgenesis.Name)
	assert.NoError(t, err)
	obsc, err := scob.NewOpenboxgenesis(common.HexToAddress(addrOB), cli.Client())
	assert.NoError(t, err)
	_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return obsc.SetupOpener(opts, accBackend)
	})
	assert.NoError(t, err)

	// setup minter
	addrNft, err := util.GetContractAddress(addressFile, nft.Name)
	assert.NoError(t, err)
	nftsc, err := scnft.NewNFToken(common.HexToAddress(addrNft), cli.Client())
	assert.NoError(t, err)

	_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return nftsc.SetupMinter(opts, common.HexToAddress(addrOB))
	})
	assert.NoError(t, err)

	// add subscriber
	_, err = cli.Transact(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return obsc.AddSubscribers(opts, []common.Address{buyer})
	})
	assert.NoError(t, err)

	// transfer ETH to backend
	var data []byte
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	assert.NoError(t, err)
	derivationPath := hdwallet.MustParseDerivationPath(constant.EthDerivationPath)
	account, err := wallet.Derive(derivationPath, true)
	assert.NoError(t, err)
	pk, err := wallet.PrivateKey(account)
	assert.NoError(t, err)
	client, err := rpc.Dial(rpc.BSC, os.Getenv("CHAIN_ID"))
	assert.NoError(t, err)
	nonce, err := client.PendingNonceAt(context.Background(), account.Address)
	assert.NoError(t, err)
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)
	tx := types.NewTransaction(nonce, accBackend, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	assert.NoError(t, err)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
	assert.NoError(t, err)
	err = client.SendTransaction(context.Background(), signedTx)
	assert.NoError(t, err)

	return obsc
}

func setupBuyer(t *testing.T, pkBackend *ecdsa.PrivateKey) (common.Address, *bind.TransactOpts) {
	// accounts[1] in ganache, because has ETH
	pkBuyer, err := crypto.HexToECDSA("b2283dbf0d49a8d1f3dc2ca1f1ca220fe232d29145f400ab30e246366372501b")
	assert.NoError(t, err)
	buyer := crypto.PubkeyToAddress(pkBuyer.PublicKey)

	return buyer, bind.NewKeyedTransactor(pkBuyer)
}

func TestOpenbox(t *testing.T) {
	pkBackend, storage := setupBackend(t)
	accBackend := crypto.PubkeyToAddress(pkBackend.PublicKey)
	baseURIToken := os.Getenv("LOCAL_STORAGE_URI")

	buyer, auth := setupBuyer(t, pkBackend)

	addressFile, teardown := deployContracts(t, accBackend)
	defer teardown()

	obsc := setupSmartContract(t, addressFile, accBackend, buyer)

	//authenticate account backend
	mnemonic, err := util.GetSecrets(backendSecret, "mnemonic")
	assert.NoError(t, err)
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, os.Getenv("CHAIN_ID"), secrets)
	assert.NoError(t, err)
	metadt := metadata.NewMetadata(baseURIToken, storage)
	opener, err := NewBoxOpener(cli, metadt, boxesConfigFile)
	assert.NoError(t, err)

	var hash, cRandom [32]byte
	// backend commit
	sRandom, err := cryp.MakeRand()
	assert.NoError(t, err)
	hashB := cryp.HashMessage(sRandom)
	copy(hash[:], hashB)
	sign, err := cryp.EthSign(pkBackend, sRandom)
	assert.NoError(t, err)

	// buybox
	cRandomb, err := cryp.MakeRand()
	assert.NoError(t, err)
	copy(cRandom[:], cRandomb)
	auth.Value = big.NewInt(boxPrice)
	_, err = obsc.BuyBox(auth, supportedGrade, hash, sign, cRandom)
	assert.NoError(t, err)
	// reset
	auth.Value = nil

	err = opener.OpenBox(filepath.Join(assetsBaseDir, metadata.NftCharacter), supportedGrade, hash, sRandom, cRandomb)
	assert.NoError(t, err)

	err = opener.OpenBox(filepath.Join(assetsBaseDir, metadata.NftCharacter), nonSupportedGrade, hash, sRandom, cRandomb)
	assert.EqualError(t, err, ErrBoxNotSupport)
}

func TestOpenBoxConcurrent(t *testing.T) {
	// number of concurrent requests to openbox
	num := 100

	pkBackend, storage := setupBackend(t)
	accBackend := crypto.PubkeyToAddress(pkBackend.PublicKey)
	baseURIToken := os.Getenv("LOCAL_STORAGE_URI")

	buyer, auth := setupBuyer(t, pkBackend)

	addressFile, teardown := deployContracts(t, accBackend)
	defer teardown()

	obsc := setupSmartContract(t, addressFile, accBackend, buyer)

	//authenticate account backend
	mnemonic, err := util.GetSecrets(backendSecret, "mnemonic")
	assert.NoError(t, err)
	secrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, os.Getenv("CHAIN_ID"), secrets)
	assert.NoError(t, err)
	metadt := metadata.NewMetadata(baseURIToken, storage)
	opener, err := NewBoxOpener(cli, metadt, boxesConfigFile)
	assert.NoError(t, err)

	var hash, cRandom [32]byte
	// backend commit
	sRandom, err := cryp.MakeRand()
	assert.NoError(t, err)
	hashB := cryp.HashMessage(sRandom)
	copy(hash[:], hashB)
	sign, err := cryp.EthSign(pkBackend, sRandom)
	assert.NoError(t, err)

	// buybox
	cRandomb, err := cryp.MakeRand()
	assert.NoError(t, err)
	copy(cRandom[:], cRandomb)
	auth.Value = big.NewInt(boxPrice)
	_, err = obsc.BuyBox(auth, supportedGrade, hash, sign, cRandom)
	assert.NoError(t, err)
	// reset
	auth.Value = nil

	var wg sync.WaitGroup
	doOpenBox := func() {
		err = opener.OpenBox(filepath.Join(assetsBaseDir, metadata.NftCharacter), supportedGrade, hash, sRandom, cRandomb)
		assert.NoError(t, err)
		wg.Done()
	}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go doOpenBox()
	}
	wg.Wait()
}
