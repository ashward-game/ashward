package model

import (
	"math/big"

	"github.com/shopspring/decimal"
)

type Marketplace struct {
	Model
	NFTId  uint            `gorm:"index" json:"nft_id"`
	Seller string          `gorm:"index;type:varchar(255)" json:"seller"`
	Buyer  string          `gorm:"index;type:varchar(255)" json:"buyer"`
	Price  decimal.Decimal `gorm:"index;type:decimal(30,0)" json:"price"`
	Status string          `gorm:"index;type:varchar(255)" json:"status"`
}

const OnSale = "on_sale"
const Cancelled = "cancelled"
const Sold = "sold"
const TradingHistoryLimit = 10

func NewMarketplace(nftId uint, owner string, status string, price *big.Int) *Marketplace {
	marketplace := &Marketplace{}
	marketplace.NFTId = nftId
	marketplace.Seller = owner
	marketplace.Status = status
	marketplace.SetPrice(price)
	return marketplace
}

func (marketplace *Marketplace) SetPrice(price *big.Int) {
	marketplace.Price = decimal.NewFromBigInt(price, 0)
}
