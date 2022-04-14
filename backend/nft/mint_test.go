package nft

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"orbit_nft/constant"
	"orbit_nft/contract"
	"orbit_nft/contract/abi/nft"
	"orbit_nft/contract/rpc"
	serviceNft "orbit_nft/contract/service/nft"
	"orbit_nft/nft/metadata"
	"orbit_nft/storage"
	"orbit_nft/storage/ipfs"
	"orbit_nft/storage/localstorage"
	"orbit_nft/testutil"
	"orbit_nft/util"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var expectedToken = map[string]string{"name": "nft1", "image": "image"}

const secretFile = "../../contracts/secrets.json"

func deployContracts(t *testing.T, baseURI string) (string, func()) {
	address := make(map[string]string)
	deployer := testutil.NewDeployer(t, os.Getenv("CHAIN_ID"), secretFile)
	add := deployer.DeployContract(t, "NFT", "NFT", "NFT", baseURI)
	address["NFT"] = add

	return testutil.WriteAddressToFile(t, address)
}

func nftService(t *testing.T, cli *contract.Client) *serviceNft.NFToken {
	addressNFT, err := util.GetContractAddress(cli.AddressFile(), nft.Name)
	assert.NoError(t, err)

	addressContract := common.HexToAddress(addressNFT)
	nftContract, err := serviceNft.NewNFToken(addressContract, cli.Client())
	assert.NoError(t, err)
	return nftContract
}

func setupMintor(t *testing.T, storage storage.Storage, baseUri, addressFile string) (*NftMint, error) {
	chainId := os.Getenv("CHAIN_ID")

	logger, _ := zap.NewDevelopment()

	mnemonic, err := util.GetSecrets(secretFile, "mnemonic")
	assert.NoError(t, err)

	clientSecrets := contract.NewClientSecret(mnemonic, constant.EthDerivationPath)
	cli, err := contract.NewBscAuthenticatedClient(addressFile, chainId, clientSecrets, contract.WithLogger(logger))
	assert.NoError(t, err)

	metadata := metadata.NewMetadata(baseUri, storage)
	return NewMintor(cli, metadata)
}

func TestMintIPFS(t *testing.T) {
	err := godotenv.Load("../.test-env")
	assert.NoError(t, err)
	rpc.Initialize("../config/rpc.json")

	addressFile, teardown := deployContracts(t, os.Getenv("IPFS_URI"))
	defer teardown()

	ipfsShell, err := ipfs.NewShell(os.Getenv("IPFS_SHELL_URI"))
	assert.NoError(t, err)

	mintor, err := setupMintor(t, ipfsShell, os.Getenv("IPFS_URI"), addressFile)
	assert.NoError(t, err)

	// mint
	err = mintor.Mint("../../assets/mocks/character")
	assert.NoError(t, err)

	// get tokenURI
	nftContract := nftService(t, mintor.client)
	tokenURI, err := nftContract.TokenURI(big.NewInt(0))
	assert.NoError(t, err)

	assertContentMetadata(t, tokenURI)
	assertContentImage(t, tokenURI)
}

func TestMintLocalStorage(t *testing.T) {
	err := godotenv.Load("../.test-env")
	assert.NoError(t, err)
	rpc.Initialize("../config/rpc.json")

	workingDir := "../../assets/nft"
	localStorage, err := localstorage.New(workingDir)
	assert.NoError(t, err)

	addressFile, teardown := deployContracts(t, os.Getenv("LOCAL_STORAGE_URI"))
	defer teardown()
	mintor, err := setupMintor(t, localStorage, os.Getenv("LOCAL_STORAGE_URI"), addressFile)
	assert.NoError(t, err)

	// create nft character
	err = mintor.Mint("../../assets/mocks/character")
	assert.NoError(t, err)

	// get tokenURI
	nftContract := nftService(t, mintor.client)
	tokenURI, err := nftContract.TokenURI(big.NewInt(0))
	assert.NoError(t, err)
	assert.Contains(t, tokenURI, os.Getenv("LOCAL_STORAGE_URI"))
}

func assertContentMetadata(t *testing.T, tokenURI string) {
	res, err := http.Get(strings.Replace(tokenURI, "localhost", "127.0.0.1", 1))
	assert.NoError(t, err, err)
	defer res.Body.Close()

	var token map[string]interface{}
	content, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err, err)
	json.Unmarshal(content, &token)
	assert.Equal(t, expectedToken["name"], token["name"])
}

func assertContentImage(t *testing.T, tokenURI string) {
	res, err := http.Get(strings.Replace(tokenURI, "localhost", "127.0.0.1", 1))
	assert.NoError(t, err)
	defer res.Body.Close()

	var token map[string]interface{}
	content, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	json.Unmarshal(content, &token)

	res, err = http.Get(strings.Replace(token["image"].(string), "localhost", "127.0.0.1", 1))
	assert.NoError(t, err)
	defer res.Body.Close()

	contentText, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, expectedToken["image"], string(contentText))
}
