package util_test

import (
	"orbit_nft/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLink(t *testing.T) {
	base := "https://ipfs.io/ipfs/"
	path1 := "some_ipfs_hash"
	path2 := "image.png"

	expected := base + path1 + "/" + path2

	actual, err := util.ToLink(base, path1, path2)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
