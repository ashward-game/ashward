package model

type Hero struct {
	Model
	Owner    string `gorm:"index;type:varchar(255)" json:"owner"`
	Metadata string `gorm:"type:varchar(255);not null" json:"metadata"`
}
