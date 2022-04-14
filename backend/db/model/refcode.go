package model

import (
	"gorm.io/datatypes"
)

type Refcode struct {
	Model
	Code      string         `gorm:"unique;index;type:varchar(255)" json:"code"`
	Owner     string         `gorm:"unique;index;type:varchar(255)" json:"owner"`
	Used      uint           `gorm:"type:bigint;default:0;not null" json:"used"`
	NumHero   uint           `gorm:"type:bigint;default:0;not null" json:"num_hero"`
	Addresses datatypes.JSON `gorm:"type:json" json:"addresses"`
}
