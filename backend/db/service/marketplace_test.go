package service

import (
	"errors"
	"math/big"
	"orbit_nft/db/repository"
	"orbit_nft/testutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMarketplaceService(t *testing.T) {
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

	err := nftService.CreateNotForSaleToken(tokenId, address1, metadataURI, rawStringMetadata)
	assert.NoError(t, err)

	//test OpenOffer
	err = service.OpenOffer(uint(10000), address1, big.NewInt(15))
	assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	err = service.OpenOffer(tokenId, address2, big.NewInt(15))
	assert.EqualError(t, err, errors.New(ErrOwnerAndSellerNotMatch).Error())
	err = service.OpenOffer(tokenId, address1, big.NewInt(15))
	assert.NoError(t, err)
	err = service.OpenOffer(tokenId, address1, big.NewInt(15))
	assert.EqualError(t, err, errors.New(ErrNftIsOnMarketplace).Error())

	//test CancalOffer
	err = service.CancelOffer(uint(100), address1)
	assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	err = service.CancelOffer(tokenId, address2)
	assert.EqualError(t, err, errors.New(ErrOwnerAndSellerNotMatch).Error())
	err = service.CancelOffer(tokenId, address1)
	assert.NoError(t, err)
	err = service.CancelOffer(tokenId, address1)
	assert.EqualError(t, err, errors.New(ErrNftNotOnMarketplace).Error())

	//test Purchase
	err = service.Purchase(uint(100), address2)
	assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	err = service.Purchase(tokenId, address2)
	assert.EqualError(t, err, errors.New(ErrNftNotOnMarketplace).Error())
	err = service.OpenOffer(tokenId, address1, big.NewInt(12))
	assert.NoError(t, err)
	err = service.Purchase(tokenId, address2)
	assert.NoError(t, err)
}
