package model

type MakeRand struct {
	Model
	Hash   string `gorm:"uniqueIndex;type:varchar(255)" json:"hash"`
	Random string `gorm:"type:varchar(255)" json:"random"`
}
