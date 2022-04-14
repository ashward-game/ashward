package get

import "orbit_nft/api/util/pagination"

type InputList struct {
	Type         string `form:"type" json:"type" binding:"required,oneof=character pet"`
	Class        string `form:"class" json:"class" binding:"omitempty,oneof=ranged melee mage"`
	Rarity       string `form:"rarity" json:"rarity" binding:"omitempty,oneof=normal rare epic legendary mythical"`
	Search       string `form:"search" json:"search" binding:"omitempty,alphaUnicodeNumericSpaceHyphen"`
	OrderByPrice string `form:"order_by_price" json:"order_by_price" binding:"omitempty,oneof=asc desc"`
	pagination.InputPagination
}

type InputShowNft struct {
	Id *int `uri:"id" binding:"required,min=0"`
}

type InputTradingHistory struct {
	Address string `uri:"address" binding:"required,alphanum"`
	pagination.InputPagination
}
