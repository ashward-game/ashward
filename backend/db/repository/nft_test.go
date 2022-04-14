package repository

import (
	"math/big"
	"orbit_nft/db/model"
	"orbit_nft/db/util"
	"orbit_nft/testutil"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newMockNFTRepo() (*NftRepository, func()) {
	db, teardown := testutil.NewMockDB()

	db.AutoMigrate()
	return NewNftRepository(db), teardown
}

func TestCreateNFTNotForSale(t *testing.T) {
	db, teardown := newMockNFTRepo()
	defer teardown()

	expect := &model.NFT{
		TokenId: 1,
		Name:    "orbit_nft",
		Owner:   "0x00",
	}

	err := db.CreateNotForSaleToken(expect)
	assert.NoError(t, err)

	actual, err := db.FindByTokenID(1).FirstToken()
	assert.NoError(t, err)
	assert.False(t, actual.IsForSale())

	// explicitly verify by raw sql query
	var record model.NFT
	err = db.db.DB.Select("token_id").Where("price IS NULL").First(&record).Error
	assert.NoError(t, err)
	assert.Equal(t, expect.TokenId, record.TokenId)
}

func TestCreateNFTForSale(t *testing.T) {
	db, teardown := newMockNFTRepo()
	defer teardown()

	expect := &model.NFT{
		TokenId: 1,
		Name:    "orbit_nft",
		Owner:   "0x00",
	}

	err := db.CreateForSaleToken(expect, big.NewInt(100))
	assert.NoError(t, err)

	actual, err := db.FindByTokenID(1).FirstToken()
	assert.NoError(t, err)
	assert.True(t, actual.IsForSale())
	assert.EqualValues(t, expect.SellingPrice(), actual.SellingPrice())
}

func generateTokenType(t *testing.T, i int) string {
	var typeNft string
	if i%2 == 0 {
		typeNft = "character"
	} else {
		typeNft = "wagon"
	}
	return typeNft
}

func generateTokenClass(t *testing.T, i int) string {
	var class string
	if i%25 == 0 {
		class = "Mage"
	} else {
		class = "Ranged"
	}
	return class
}

func generateTokenRarity(t *testing.T, i int) string {
	var rarity string
	if i%2 == 0 {
		rarity = "normal"
	}
	return rarity
}

func TestNftCreation(t *testing.T) {
	db, teardown := newMockNFTRepo()
	defer teardown()

	// insert records for testing
	for i := 0; i < 100; i++ {
		nft := &model.NFT{
			TokenId: uint(i),
			Type:    generateTokenType(t, i),
			Name:    "name" + strconv.Itoa(i),
			Class:   generateTokenClass(t, i),
			Rarity:  generateTokenRarity(t, i),
		}
		err := db.CreateForSaleToken(nft, big.NewInt(1))
		assert.NoError(t, err)
	}
	for i := 101; i < 150; i++ {
		nft := &model.NFT{
			TokenId: uint(i),
			Type:    generateTokenType(t, i),
			Name:    "name" + strconv.Itoa(i),
			Class:   generateTokenClass(t, i),
			Rarity:  generateTokenRarity(t, i),
		}
		err := db.CreateNotForSaleToken(nft)
		assert.NoError(t, err)
	}

	// query
	testTotalNft(t, db)
	testWithType(t, db)
	testWithClass(t, db)
	testWithRarity(t, db)
	testAll(t, db)
}

func testTotalNft(t *testing.T, repo *NftRepository) {
	pg := &util.Pagination{}
	var result []model.NFT
	err := repo.FindTokens(pg).OnSale().Find(&result)
	assert.NoError(t, err)
	assert.Equal(t, 100, len(result))
}

func testWithType(t *testing.T, repo *NftRepository) {
	pg := &util.Pagination{}
	result := make([]model.NFT, 0)
	err := repo.FindTokens(pg).OnSale().WithType("character").Find(&result)
	assert.NoError(t, err)
	assert.Equal(t, 50, len(result))
}

func testWithClass(t *testing.T, repo *NftRepository) {
	pg := &util.Pagination{}
	result := make([]model.NFT, 0)
	err := repo.FindTokens(pg).OnSale().WithClass("Mage").Find(&result)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(result))
}

func testWithRarity(t *testing.T, repo *NftRepository) {
	pg := &util.Pagination{}
	result := make([]model.NFT, 0)
	err := repo.FindTokens(pg).OnSale().WithRarity("normal").Find(&result)
	assert.NoError(t, err)
	assert.Equal(t, 50, len(result))
}

func testAll(t *testing.T, repo *NftRepository) {
	pg := &util.Pagination{}
	result := make([]model.NFT, 0)
	err := repo.FindTokens(pg).OnSale().
		WithType("character").
		WithClass("Mage").
		WithRarity("normal").
		SortByPrice("desc").
		Find(&result)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
}
