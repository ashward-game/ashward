package model

import (
	"encoding/json"
	"math/big"

	"github.com/shopspring/decimal"
)

type NFT struct {
	Model
	TokenId     uint                `gorm:"uniqueIndex" json:"token_id"`
	Name        string              `gorm:"index;type:varchar(255)" json:"name"`
	Type        string              `gorm:"index;type:varchar(255)" json:"type"`
	Owner       string              `gorm:"index;type:varchar(255)" json:"owner"`
	Image       string              `gorm:"type:text" json:"image"`
	MetadataURI string              `gorm:"type:text" json:"metadata_uri"`
	Price       decimal.NullDecimal `gorm:"index;type:decimal(30,0)" json:"price"`
	Rarity      string              `gorm:"index;type:varchar(255)" json:"rarity"`
	Class       string              `gorm:"index;type:varchar(255)" json:"class"`

	Marketplaces []Marketplace `gorm:"foreignKey:NFTId" json:"marketplaces"`
}

func NewNFToken(tokenId uint, owner, metadataURI string, rawStringMetadata string) *NFT {
	nft := &NFT{}
	var mapData map[string]interface{}
	json.Unmarshal([]byte(rawStringMetadata), &mapData)

	var class string
	var rarity string
	var name string
	var tokenType string
	var image string

	if v, ok := mapData["name"].(string); ok {
		name = v
	}
	if v, ok := mapData["type"].(string); ok {
		tokenType = v
	}
	if v, ok := mapData["image"].(string); ok {
		image = v
	}

	if properties, ok := mapData["properties"].(map[string]interface{}); ok {
		if mapClass, ok := properties["class"].(map[string]interface{}); ok {
			if strClass, ok := mapClass["value"].(string); ok {
				class = strClass
			}
		}
		if mapRatity, ok := properties["grade"].(map[string]interface{}); ok {
			if strRarity, ok := mapRatity["value"].(string); ok {
				rarity = strRarity
			}
		}
	}

	nft.Class = class
	nft.Rarity = rarity
	nft.Name = name
	nft.Type = tokenType
	nft.Owner = owner
	nft.TokenId = tokenId
	nft.MetadataURI = metadataURI
	nft.Image = image
	return nft
}

func (token *NFT) NotForSale() {
	token.Price.Decimal = decimal.NewFromInt(-1)
	token.Price.Valid = false
}

func (token *NFT) Sell(price *big.Int) {
	token.Price.Decimal = decimal.NewFromBigInt(price, 0)
	token.Price.Valid = true
}

func (token *NFT) IsForSale() bool {
	return token.Price.Valid
}

func (token *NFT) SellingPrice() *big.Int {
	return token.Price.Decimal.BigInt()
}
