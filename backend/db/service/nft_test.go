package service

import (
	"errors"
	"math/big"
	"orbit_nft/db/model"
	"orbit_nft/db/repository"
	"orbit_nft/testutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var metadataURI = "metadaURI"
var rawStringMetadata = `{}`

func TestCreateNotForSaleThenTransfer(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()

	repo := repository.NewNftRepository(db)
	service := NewNftService(repo)

	tokenId := uint(1)
	tokenNotFound := uint(9)
	address1 := common.HexToAddress("0x0beef").Hex()
	address2 := common.HexToAddress("0x0aaaa").Hex()

	err := service.CreateNotForSaleToken(tokenId, address1, metadataURI, rawStringMetadata)
	assert.NoError(t, err)
	err = service.TransferNFT(tokenId, address1, address2)
	assert.NoError(t, err)
	err = service.TransferNFT(tokenId, address1, address2)
	assert.EqualError(t, err, errors.New(ErrNftTransferFromNotFound).Error())
	err = service.TransferNFT(tokenNotFound, address1, address2)
	assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())

	err = service.TransferNFT(tokenId, "", address2)
	assert.NoError(t, err)
}

func TestQueryNFTWithTradingHistory(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()
	// _ = teardown

	nftRepo := repository.NewNftRepository(db)
	marketRepo := repository.NewMarketplaceRepository(db)
	nftService := NewNftService(nftRepo)
	service := NewMarketplaceService(marketRepo, nftRepo)

	tokenId := uint(1)
	address1 := common.HexToAddress("0x0beef").Hex()
	address2 := common.HexToAddress("0x0aaaa").Hex()

	_, err := nftService.GetTokenById(uint(1000))
	assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())

	err = nftService.CreateNotForSaleToken(tokenId, address1, metadataURI, rawStringMetadata)
	assert.NoError(t, err)

	err = service.OpenOffer(tokenId, address1, big.NewInt(15))
	assert.NoError(t, err)
	err = service.CancelOffer(tokenId, address1)
	assert.NoError(t, err)
	err = service.OpenOffer(tokenId, address1, big.NewInt(12))
	assert.NoError(t, err)
	err = service.Purchase(tokenId, address2)
	assert.NoError(t, err)

	nft, err := nftService.GetTokenById(tokenId)
	assert.NoError(t, err)
	assert.EqualValues(t, nft.Owner, address2)
	assert.EqualValues(t, nft.Marketplaces[0].Buyer, address2)

	err = service.OpenOffer(tokenId, address2, big.NewInt(13))
	assert.NoError(t, err)
	err = service.Purchase(tokenId, address1)
	assert.NoError(t, err)

	nft, err = nftService.GetTokenById(tokenId)
	assert.NoError(t, err)

	assert.EqualValues(t, len(nft.Marketplaces), 2)
	assert.EqualValues(t, nft.Marketplaces[0].Status, model.Sold)
	assert.EqualValues(t, nft.Marketplaces[0].Buyer, address1)
	assert.EqualValues(t, nft.Owner, address1)
}
