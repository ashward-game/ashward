package repository

import (
	"orbit_nft/db"
	"orbit_nft/db/model"
	"orbit_nft/db/util"

	"gorm.io/gorm"
)

type MarketplaceRepository struct {
	db *db.Database
}

func NewMarketplaceRepository(db *db.Database) *MarketplaceRepository {
	return &MarketplaceRepository{
		db: db,
	}
}

func (e *MarketplaceRepository) Transaction(fc func(tx *gorm.DB) error) error {
	return e.db.DB.Transaction(fc)
}

func (e *MarketplaceRepository) WithTx(tx *gorm.DB) *MarketplaceRepository {
	return NewMarketplaceRepository(&db.Database{DB: tx})
}

func (e *MarketplaceRepository) Create(marketplace *model.Marketplace) error {
	return e.db.DB.Model(&model.Marketplace{}).Create(marketplace).Error
}

func (e *MarketplaceRepository) Update(marketplace *model.Marketplace) error {
	result := e.db.DB.Where(&model.Marketplace{
		Model: model.Model{
			ID: marketplace.ID,
		},
	}).Updates(marketplace)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return db.ErrNoRowsAffected
	}
	return nil
}

func (e *MarketplaceRepository) FindOnSaleByTokenId(nftId uint) (*model.Marketplace, error) {
	var marketplace model.Marketplace
	if err := e.db.DB.Where(&model.Marketplace{NFTId: nftId, Status: model.OnSale}).First(&marketplace).Error; err != nil {
		return nil, err
	}
	return &marketplace, nil
}

func (e *MarketplaceRepository) FindTradingHistoryOf(pg *util.Pagination, address string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := e.db.DB.Model(&model.Marketplace{}).
		Scopes(util.Paginate(pg, "`marketplaces`.*, `nfts`.token_id, `nfts`.id")).
		Joins("join `nfts` on `nfts`.id = `marketplaces`.nft_id").
		Where(&model.Marketplace{Buyer: address}).
		Or(&model.Marketplace{Seller: address}).
		Order("`marketplaces`.created_at desc").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
