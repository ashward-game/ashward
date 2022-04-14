package model

type Metadata struct {
	Model
	Name  string `gorm:"uniqueIndex;type:char(255)"`
	Value string `gorm:"type:text"`
}

const MetadataCurrentBlock = "CurrentBlock"
