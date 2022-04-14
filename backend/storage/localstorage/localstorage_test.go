package localstorage

import (
	"orbit_nft/storage"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAddFile(t *testing.T) {
	// setting up
	err := godotenv.Load("../../.test-env")
	assert.NoError(t, err, err)

	baseDir := "../../../assets/nft"
	localStorage, err := New(baseDir)
	assert.NoError(t, err, err)

	var storageIpfs storage.Storage = localStorage

	fileHash, err := storageIpfs.AddFile("../../../assets/mocks/character/normal/images/image.txt")
	assert.NoError(t, err, err)

	content, err := os.ReadFile(filepath.Join(baseDir, fileHash))
	assert.NoError(t, err, err)

	assert.Equal(t, string(content), "image")
}
