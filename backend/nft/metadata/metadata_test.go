package metadata

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"orbit_nft/storage/ipfs"
	"orbit_nft/storage/localstorage"
	"orbit_nft/util"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const envPath = "../../.test-env"
const assetsBaseDir = "../../../assets"

var expectedToken = map[string]string{"name": "nft1", "image": "image"}

func TestGenerateMetadataLocalstorage(t *testing.T) {
	err := godotenv.Load(envPath)
	assert.NoError(t, err)

	workingDir := filepath.Join(assetsBaseDir, "nft")
	baseURI := os.Getenv("LOCAL_STORAGE_URI")
	rariryPathMockNft := filepath.Join(assetsBaseDir, "mocks", NftCharacter, "normal")

	localStorage, err := localstorage.New(workingDir)
	assert.NoError(t, err)

	contentCsv, err := util.ReadFileCsv(filepath.Join(rariryPathMockNft, "data.csv"))
	assert.NoError(t, err)
	assert.Equal(t, 2, len(contentCsv))

	meta := NewMetadata(baseURI, localStorage)
	// skip first line: first line is header
	metadataCid, err := meta.GenerateMetadata(rariryPathMockNft, contentCsv[1])
	assert.NoError(t, err)
	assert.Equal(t, ".json", filepath.Ext(metadataCid))
}

func TestGenerateMetadataIPFS(t *testing.T) {
	err := godotenv.Load(envPath)
	assert.NoError(t, err)

	ipfsShell, err := ipfs.NewShell(os.Getenv("IPFS_SHELL_URI"))
	assert.NoError(t, err)
	baseURI := os.Getenv("IPFS_URI")
	rariryPathMockNft := filepath.Join(assetsBaseDir, "mocks", NftCharacter, "normal")

	contentCsv, err := util.ReadFileCsv(filepath.Join(rariryPathMockNft, "data.csv"))
	assert.NoError(t, err)
	assert.Equal(t, 2, len(contentCsv))

	meta := NewMetadata(baseURI, ipfsShell)
	// skip first line: first line is header
	metadataCid, err := meta.GenerateMetadata(rariryPathMockNft, contentCsv[1])
	assert.NoError(t, err)

	tokenURI, err := util.ToLink(baseURI, metadataCid)
	assert.NoError(t, err)

	assertContentMetadata(t, tokenURI)
	assertContentImage(t, tokenURI)
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
