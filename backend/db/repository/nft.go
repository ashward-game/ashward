package repository

import (
	"fmt"
	"math/big"
	"orbit_nft/db"
	"orbit_nft/db/model"
	"orbit_nft/db/util"
	"strconv"

	"gorm.io/gorm"
)

type gormWrapper struct {
	gorm *gorm.DB
}

func (query *gormWrapper) WithType(tokenType string) *gormWrapper {
	if len(tokenType) == 0 {
		return query
	}
	query.gorm = query.gorm.Where(&model.NFT{
		Type: tokenType,
	})
	return query
}

func (query *gormWrapper) WithClass(class string) *gormWrapper {
	query.gorm = query.gorm.Where(&model.NFT{
		Class: class,
	})
	return query
}

func (query *gormWrapper) WithRarity(rarity string) *gormWrapper {
	if len(rarity) == 0 {
		return query
	}
	query.gorm = query.gorm.Where(&model.NFT{
		Rarity: rarity,
	})
	return query
}

func (query *gormWrapper) WithTradingHistory(args ...interface{}) *gormWrapper {
	if len(args) == 0 {
		args = append(args, func(db *gorm.DB) *gorm.DB {
			return db.Order("`marketplaces`.created_at desc").Where("`marketplaces`.status = ?", model.Sold).Limit(model.TradingHistoryLimit)
		})
	}
	query.gorm = query.gorm.Preload("Marketplaces", args...)
	return query
}

func (query *gormWrapper) ByNameOrTokenId(search string) *gormWrapper {
	if len(search) == 0 {
		return query
	}
	if tokenId, err := strconv.Atoi(search); err != nil {
		query.gorm = query.gorm.Where("`nfts`.name like ?", "%"+search+"%")
	} else {
		query.gorm = query.gorm.Where(&model.NFT{
			TokenId: uint(tokenId),
		})
	}
	return query
}

func (query *gormWrapper) SortByPrice(sort string) *gormWrapper {
	if sort == "asc" || sort == "desc" {
		query.gorm = query.gorm.Order(fmt.Sprintf("`nfts`.price %s", sort))
	}
	return query
}

func (query *gormWrapper) Find(dest interface{}, conds ...interface{}) error {
	return query.gorm.Find(dest, conds...).Error
}

func (query *gormWrapper) First(dest interface{}, conds ...interface{}) error {
	return query.gorm.First(dest, conds...).Error
}

func (query *gormWrapper) FirstToken() (*model.NFT, error) {
	var nft model.NFT
	if err := query.gorm.First(&nft).Error; err != nil {
		return nil, err
	}
	return &nft, nil
}

type NftRepository struct {
	db *db.Database
}

func NewNftRepository(db *db.Database) *NftRepository {
	return &NftRepository{
		db: db,
	}
}

func (e *NftRepository) WithTx(tx *gorm.DB) *NftRepository {
	return NewNftRepository(&db.Database{DB: tx})
}

func (e *NftRepository) CreateForSaleToken(nft *model.NFT, price *big.Int) error {
	nft.Sell(price)
	return e.db.DB.Model(&model.NFT{}).Create(nft).Error
}

func (e *NftRepository) CreateNotForSaleToken(nft *model.NFT) error {
	nft.NotForSale()
	return e.db.DB.Model(&model.NFT{}).Create(nft).Error
}

func (e *NftRepository) Update(nft *model.NFT) error {
	result := e.db.DB.
		Where(&model.NFT{TokenId: nft.TokenId}).
		Updates(nft)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return db.ErrNoRowsAffected
	}
	return nil
}

func (e *NftRepository) FindByTokenID(tokenId uint) *gormWrapper {
	query := &gormWrapper{gorm: e.db.DB}
	query.gorm = e.db.DB.Where(&model.NFT{TokenId: tokenId})
	return query
}

func (e *NftRepository) FindTokens(pg *util.Pagination) *gormWrapper {
	query := &gormWrapper{gorm: e.db.DB}
	query.gorm = e.db.DB.Model(&model.NFT{}).Scopes(util.Paginate(pg, "`nfts`.*"))
	return query
}

func (query *gormWrapper) OnSale() *gormWrapper {
	query.gorm = query.gorm.Where("`nfts`.price IS NOT NULL")
	return query
}

func (query *gormWrapper) OwnedBy(owner string) *gormWrapper {
	query.gorm = query.gorm.Where(&model.NFT{Owner: owner})
	return query
}
