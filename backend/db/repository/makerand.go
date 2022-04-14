package repository

import (
	"orbit_nft/db"
	"orbit_nft/db/model"
)

type MakeRandRepository struct {
	db *db.Database
}

func NewMakeRandRepository(db *db.Database) *MakeRandRepository {
	return &MakeRandRepository{
		db: db,
	}
}

func (r *MakeRandRepository) Get(hash string) (*model.MakeRand, error) {
	var result model.MakeRand
	if err := r.db.DB.Model(&model.MakeRand{}).Where(&model.MakeRand{
		Hash: hash,
	}).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *MakeRandRepository) Create(hash, random string) error {
	return r.db.DB.Model(&model.MakeRand{}).Create(&model.MakeRand{
		Hash:   hash,
		Random: random,
	}).Error
}
