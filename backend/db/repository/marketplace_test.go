package repository

import (
	"math/big"
	"orbit_nft/db"
	"orbit_nft/db/model"
	"orbit_nft/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newMockMarketplaceRepo() (*MarketplaceRepository, *NftRepository, func()) {
	db, teardown := testutil.NewMockDB()

	db.AutoMigrate()
	return NewMarketplaceRepository(db), NewNftRepository(db), teardown
}

func TestTransactionSuccess(t *testing.T) {
	mkRepo, nftRepo, teardown := newMockMarketplaceRepo()
	defer teardown()

	price := big.NewInt(1)
	marketplace := model.NewMarketplace(1, "", model.OnSale, price)
	nft := model.NewNFToken(1, "", "", "")

	nftRepo.CreateNotForSaleToken(nft)

	err := mkRepo.Transaction(func(tx *gorm.DB) error {
		if err := mkRepo.WithTx(tx).Create(marketplace); err != nil {
			return err
		}

		nft.Sell(price)
		if err := nftRepo.WithTx(tx).Update(nft); err != nil {
			return err
		}
		return nil
	})
	assert.NoError(t, err)
}

func TestTransactionFailLast(t *testing.T) {
	mkRepo, nftRepo, teardown := newMockMarketplaceRepo()
	defer teardown()

	price := big.NewInt(1)
	marketplace := model.NewMarketplace(1, "", model.OnSale, price)
	nft := model.NewNFToken(1, "", "", "")

	err := mkRepo.Transaction(func(tx *gorm.DB) error {
		err := mkRepo.WithTx(tx).Create(marketplace)
		assert.NoError(t, err)

		_, err = nftRepo.WithTx(tx).FindByTokenID(1).FirstToken()
		assert.Error(t, err)

		// should fail here
		nft.Sell(price)
		err = nftRepo.WithTx(tx).Update(nft)
		assert.EqualError(t, err, db.ErrNoRowsAffected.Error())
		return err
	})
	assert.EqualError(t, err, db.ErrNoRowsAffected.Error())
	var count int64
	mkRepo.db.DB.Model(&marketplace).Count(&count)
	assert.EqualValues(t, 0, count)
}
